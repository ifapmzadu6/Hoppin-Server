package memcache

import "github.com/bradfitz/gomemcache/memcache"

func Open() *memcache.Client {
	mc := memcache.New("localhost:11211")
	return mc
}

func Set(v string, k string, mc *memcache.Client) error {
	i := memcache.Item{Key: k, Value: []byte(v)}
	return mc.Set(&i)
}

func Get(k string, mc *memcache.Client) (string, error) {
	v, err := mc.Get(k)
	if err != nil {
		return "", err
	}
	return string(v.Value[:]), nil
}
