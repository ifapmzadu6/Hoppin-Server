package memcache

import "github.com/bradfitz/gomemcache/memcache"

func Open() memcache {
	mc := memcache.New("localhost:11211")
	mc.Set(&memcache.Item{Key: "foo", Value: []byte("my value")})
}
