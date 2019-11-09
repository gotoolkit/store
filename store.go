package store

type Store interface {
	Put(key string, value []byte) error

	Get(key string) ([]byte, error)

	Delete(key string) error

	Exists(key string) (bool, error)

	Close()
}
