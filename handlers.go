package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// HomeHandler serves as the index page/ status page
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the proxy!"))
}

// GetValueFromKeyHandler will GET value from local cache.
// If it fails, it fetches the value from the backing Redis instance,
// storing it in the local cache, associated with the specified key.
func GetValueFromKeyHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	val, _ := localCache.getCVal(vars["key"])
	// If local cache has the key, return value
	// Otherwise call Redis
	switch v := val.(type) {
	case string:
		w.Write([]byte(v))
	case nil:
		val, err := myRedisClient.getRVal(vars["key"])
		if err != nil {
			log.Printf("Error getting value from Redis. Error is: %v.", err)
			// Sad hack to check if Redis connection is down
			if strings.Contains(err.Error(), "connection refused") {
				http.Error(w, err.Error(), http.StatusServiceUnavailable)
			}
		} else {
			// If Redis returns a key, then set value to local cache.
			localCache.setCVal(vars["key"], val)
			w.Write([]byte(val))
		}
	}

}
