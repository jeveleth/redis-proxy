package main

import (
	"log"

	"github.com/hashicorp/golang-lru"
)

// A GET request, directed at the proxy, returns the value of the specified key
// from the proxy’s local cache if the local cache contains a value for that key.
//  If the local cache does not contain a value for the specified key,
//  it fetches the value from the backing Redis instance, using the Redis GET command,
// and stores it in the local cache, associated with the specified key.

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
		log.Printf("Local cache. Key: %s, Val: %v.", key, res)
		return res, ""
	}
	log.Printf("Local cache can't find the Key: %v.", key)
	return nil, ""
}

func (c *cache) setCVal(key string, val string) {
	c.client.Add(key, val)
	// TODO: mutex me
}
