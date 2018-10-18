package datastore

import (
	"io"

	"github.com/r2d4/sh8s/pkg/datastore/redis"
)

type Datastore interface {
	Register()
	io.Closer

	Get(key string) (string, error)
	Set(key, value string) error

	List(id string) (map[string]string, error)
	SetList(id string, keys []string) error
}

// DefaultDatastore is where the files will be stored
var DefaultDatastore Datastore

func init() {
	DefaultDatastore = &redis.RedisStore{}
}
