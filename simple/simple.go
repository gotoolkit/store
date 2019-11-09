package simple

import "errors"

type Store struct {
	container map[string][]byte
}

func New() (*Store, error) {

	return &Store{
		container: make(map[string][]byte),
	}, nil
}
func (s *Store) Put(key string, value []byte) error {
	s.container[key] = value
	return nil
}

func (s *Store) Get(key string) ([]byte, error) {
	v, ok := s.container[key]
	if !ok {
		return nil, errors.New("NO KEY")
	}
	return v, nil
}

func (s *Store) Delete(key string) error {
	delete(s.container, key)
	return nil
}

func (s *Store) Exists(key string) (bool, error) {
	_, ok := s.container[key]
	return ok, nil
}

func (s *Store) Close() {

}
