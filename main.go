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
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", proxyPort), r))
}
