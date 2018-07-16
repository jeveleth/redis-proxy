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
		log.Printf("Error creating cache client %#v", err)
	}
	return &cache{client: l}
}

// getCVal checks the local cache for a key's value.
// If it fails, it fetches the value from the backing Redis instance,
// storing it in the local cache, associated with the specified key.
func (c *cache) getCVal(key string) interface{} {
	// log.Printf("DEBUG ==> checking cache for key %s.", key)
	res, _ := c.client.Get(key)
	// if ok != true {
	// 	// TODO: Error handle better
	// 	log.Printf("error getting value from local cache %#v", ok)
	// }
	if res != nil {
		log.Printf("Local cache => key: %s and val: %v.", key, res)
		return res
	}
	// TODO: check if error is because service is down or
	// if it's because key doesn't exist.
	// If latter, return empty string
	if len(key) > 0 {
	}

	val, err := myRedisClient.getRVal(key)
	if err != nil {
		log.Printf("error getting redis value in cache %#v", err)
	}
	c.setCVal(key, val)
	return val
}

func (c *cache) setCVal(key string, val string) {
	c.client.Add(key, val)
	//  mutex me
}
