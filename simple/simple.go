package simple

import (
	"sync"

	"github.com/gotoolkit/store"
)

type Store struct {
	db sync.Map
}

func New() (*Store, error) {

	return &Store{}, nil
}
func (s *Store) Put(key string, value []byte) error {
	s.db.Store(key, value)
	return nil
}

func (s *Store) Get(key string) ([]byte, error) {
	v, ok := s.db.Load(key)
	if !ok {
		return nil, store.ErrKeyNotFound
	}
	return v.([]byte), nil
}

func (s *Store) Delete(key string) error {
	s.db.Delete(key)
	return nil
}

func (s *Store) Exists(key string) (bool, error) {
	_, ok := s.db.Load(key)
	return ok, nil
}

func (s *Store) List(dir string) ([][]byte, error) {
	return nil, store.ErrCallNotSupported
}

func (s *Store) DeleteTree(dir string) error {
	return store.ErrCallNotSupported
}

func (s *Store) Close() {
	return
}
