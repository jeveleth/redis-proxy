// TODO: Initialize a struct that the functions act on so they don't have to pass in the client each time
// TODO: restrict getter to HTTP GET Method

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var localCache = newCache(thisConfig.CacheCapacity)      // TODO: Make this not global
var myRedisClient = newRedisClient(thisConfig.RedisAddr) // TODO: Make this not global

var thisConfig = MustLoadConfig()
var cacheExpiryTime = thisConfig.CacheExpiryTime
var proxyPort = thisConfig.ProxyPort

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/getval/{key}", GetValueFromKeyHandler)
	// r.HandleFunc("/setter/{key:[0-9a-zA-Z]+}/{value}", setValueFromKeyHandler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", proxyPort), r))
}
