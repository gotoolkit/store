package bolt

import (
	"bytes"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/gotoolkit/store"
	bolt "go.etcd.io/bbolt"
)

const (
	dbFilePerm os.FileMode = 0644

	defaultTimeout = time.Duration(10) * time.Second
	defaultBucket  = "default"
)

type Store struct {
	db     *bolt.DB
	bucket []byte
	sync.Mutex
}

func New(path string, opts ...Option) (*Store, error) {
	var (
		db  *bolt.DB
		err error
	)
	options := options{
		Options: bolt.Options{Timeout: defaultTimeout},
		Bucket:  []byte(defaultBucket),
	}

	for _, o := range opts {
		o.apply(&options)
	}

	dir, _ := filepath.Split(path)
	if err = os.MkdirAll(dir, 0750); err != nil {
		return nil, err
	}

	db, err = bolt.Open(path, dbFilePerm, &options.Options)
	if err != nil {
		return nil, err
	}

	return &Store{
		db:     db,
		bucket: options.Bucket,
	}, nil
}
func (s *Store) Put(key string, value []byte) error {
	var (
		err error
	)
	s.Lock()
	defer s.Unlock()

	err = s.db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists(s.bucket)
		if err != nil {
			return err
		}

		err = bucket.Put([]byte(key), value)
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

func (s *Store) Get(key string) ([]byte, error) {
	var (
		val []byte
		err error
	)
	s.Lock()
	defer s.Unlock()
	err = s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(s.bucket))
		v := b.Get([]byte(key))
		val = make([]byte, len(v))
		copy(val, v)
		return nil
	})

	if len(val) == 0 {
		return nil, store.ErrKeyNotFound
	}

	if err != nil {
		return nil, err
	}

	return val, nil
}

func (s *Store) Delete(key string) error {
	var (
		err error
	)
	s.Lock()
	defer s.Unlock()
	err = s.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(s.bucket)
		if bucket == nil {
			return store.ErrKeyNotFound
		}
		err := bucket.Delete([]byte(key))
		return err
	})
	return err
}

func (s *Store) Exists(key string) (bool, error) {
	var (
		val []byte
		err error
	)
	s.Lock()
	defer s.Unlock()
	err = s.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(s.bucket)
		if bucket == nil {
			return store.ErrKeyNotFound
		}

		val = bucket.Get([]byte(key))

		return nil
	})

	return len(val) > 0, err
}

func (s *Store) List(keyPrefix string) ([][]byte, error) {
	var (
		err error
	)
	s.Lock()
	defer s.Unlock()

	kv := [][]byte{}

	err = s.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(s.bucket)
		if bucket == nil {
			return store.ErrKeyNotFound
		}

		cursor := bucket.Cursor()
		prefix := []byte(keyPrefix)
		for key, v := cursor.Seek(prefix); bytes.HasPrefix(key, prefix); key, v = cursor.Next() {
			kv = append(kv, v)
		}

		return nil
	})

	if len(kv) == 0 {
		return nil, store.ErrKeyNotFound
	}

	return kv, err
}

func (s *Store) DeleteTree(keyPrefix string) error {
	var (
		err error
	)
	s.Lock()
	defer s.Unlock()

	err = s.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(s.bucket)
		if bucket == nil {
			return store.ErrKeyNotFound
		}

		cursor := bucket.Cursor()
		prefix := []byte(keyPrefix)
		for key, _ := cursor.Seek(prefix); bytes.HasPrefix(key, prefix); key, _ = cursor.Next() {
			_ = bucket.Delete([]byte(key))
		}

		return nil
	})

	return err
}

func (s *Store) Close() {
	s.Lock()
	defer s.Unlock()
	s.db.Close()
	return
}
