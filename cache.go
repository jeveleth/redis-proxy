package main

import (
	"log"
	"time"

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

// cache is a wrapper around hashicorp's lru caching library
type cache struct {
	client *lru.Cache
}

// value to be set in local cache with key
type valWithTTL struct {
	Value     string
	TTL       time.Duration
	CreatedAt time.Time
}

// newCache instantiates a new LRU cache
func newCache(cacheCapacity int) *cache {
	l, err := lru.New(cacheCapacity)
	if err != nil {
		log.Printf("Error creating cache client. Error is: %v.", err)
	}
	return &cache{client: l}
}

// getCVal retrieves the key's value from the local cache
// Complexity: O(1)
func (c *cache) getCVal(key string) (interface{}, string) {

	res, ok := c.client.Get(key)
	if ok == true {
		// Peek retrieves the key without altering its LRU state
		peek, _ := c.client.Peek(key)
		result := peek.(valWithTTL)
		// If createdAt of value is older than TTL, remove key and return nil
		if time.Since(result.CreatedAt) > result.TTL {
			c.client.Remove(key)
			return nil, ""
		}
		return res, ""
	}
	return nil, ""
}

// setCVal adds a key/value to the cache, associating a TTL and createdAt time with the value
// Complexity: O(1)
func (c *cache) setCVal(key string, val string) {
	storeVal := valWithTTL{Value: val, TTL: thisConfig.CacheExpiryTime, CreatedAt: time.Now()}
	c.client.Add(key, storeVal)

}
