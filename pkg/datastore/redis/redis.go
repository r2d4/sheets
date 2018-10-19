package redis

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/xyproto/simpleredis"
)

// RedisStore implements the Datastore interface
type RedisStore struct {
	Address string

	conn *simpleredis.ConnectionPool
}

func (r *RedisStore) Get(key string) (string, error) {
	kv := simpleredis.NewKeyValue(r.conn, "0")
	return kv.Get(key)
}

func (r *RedisStore) Set(key, value string) error {
	kv := simpleredis.NewKeyValue(r.conn, "0")
	return kv.Set(key, value)
}

func (r *RedisStore) SetList(id string, keys []string) error {
	ls := simpleredis.NewList(r.conn, id)
	if err := ls.Clear(); err != nil {
		return errors.Wrap(err, "clearing list")
	}
	for _, k := range keys {
		if err := ls.Add(k); err != nil {
			return errors.Wrap(err, "adding to list")
		}
	}
	return nil
}

func (r *RedisStore) List(id string) (map[string]string, error) {
	ls := simpleredis.NewList(r.conn, id)
	fmt.Println(ls)
	keys, err := ls.All()
	if err != nil {
		return nil, errors.Wrap(err, "getting list keys")
	}
	fmt.Println(keys)
	kvs := simpleredis.NewKeyValue(r.conn, "0")
	m := map[string]string{}
	for _, k := range keys {
		v, err := kvs.Get(k)
		if err != nil {
			return nil, errors.Wrap(err, "getting key/value")
		}
		m[k] = v
	}
	return m, nil
}

func (r *RedisStore) Register() {
	r.conn = simpleredis.NewConnectionPoolHost(r.Address)
}

func (r *RedisStore) Close() error {
	r.conn.Close()
	return nil
}
