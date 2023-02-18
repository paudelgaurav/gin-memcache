package infrastructure

import (
	"github.com/bradfitz/gomemcache/memcache"
)

var Cache *memcache.Client

func SetUpMemCache() {
	Cache = memcache.New("localhost:11211")
}
