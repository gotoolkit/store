package bolt

import (
	"time"

	bolt "go.etcd.io/bbolt"
)

type options struct {
	bolt.Options
	Bucket []byte
}

type Option interface {
	apply(*options)
}

type optionFunc func(*options)

func (f optionFunc) apply(o *options) {
	f(o)
}

func WithTimeout(timeout time.Duration) Option {
	return optionFunc(func(o *options) {
		o.Timeout = timeout
	})
}

func WithBucket(bucket []byte) Option {
	return optionFunc(func(o *options) {
		o.Bucket = bucket
	})
}
