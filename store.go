package store

import "errors"

var (
	ErrCallNotSupported = errors.New("The current call is not supported with this backend")

	ErrKeyNotFound = errors.New("Key not found in store")
)

type Store interface {
	Put(key string, value []byte) error

	Get(key string) ([]byte, error)

	Delete(key string) error

	Exists(key string) (bool, error)

	List(directory string) ([][]byte, error)

	DeleteTree(directory string) error

	Close()
}
