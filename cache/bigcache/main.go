package main

import (
	"log"
	"time"

	"github.com/allegro/bigcache"
)

func main() {
	cache, err := bigcache.NewBigCache(bigcache.DefaultConfig(10 * time.Minute))
	if err != nil {
		log.Println(err)
		return
	}
	entry, err := cache.Get("my-unique-key")
	if err != nil {
		log.Println(err)
		return
	}
	if entry == nil {
		entry = []byte("value")
		cache.Set("my-unique-key", entry)
		log.Println(string(entry))
	}
}
