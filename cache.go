package main

import (
	"log"

	"github.com/hashicorp/golang-lru"
)

// Global expiry
// Entries added to the proxy cache are expired after being in the cache for a time duration that is
// globally configured (per instance).
// After an entry is expired, a GET request will act as if the value associated with
// the key was never stored in the cache.

// LRU eviction
// Once the cache fills to capacity, the least recently used (i.e. read)
// key is evicted each time a new key needs to be added to the cache.

// Fixed key size
// The cache capacity is configured in terms of number of keys it retains.

type cache struct {
	client *lru.Cache
}

func newCache(cacheCapacity int) *cache {
	// TODO: Error handling
	l, err := lru.New(cacheCapacity)
	if err != nil {
		log.Printf("Error creating cache client. Error is: %v.", err)
	}
	return &cache{client: l}
}

// getCVal retrieves the key's value from the local cache.
func (c *cache) getCVal(key string) (interface{}, string) {
	res, ok := c.client.Get(key)
	if ok == true {
		return res, ""
	}
	return nil, ""
}

func (c *cache) setCVal(key string, val string) {
	c.client.Add(key, val)
	// TODO: mutex me
}
