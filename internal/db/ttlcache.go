package db

import (
	"log"

	"github.com/ReneKroon/ttlcache/v2"
)

func NewTTLCache() *ttlcache.Cache {
	t := ttlcache.NewCache()
	t.SetExpirationCallback(func(key string, value interface{}) {
		log.Println(key, value)
	})
	return t
}
