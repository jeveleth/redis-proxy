package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var thisConfig = MustLoadConfig()
var localCache = newCache(thisConfig.CacheCapacity)      // TODO: Make this not global
var myRedisClient = newRedisClient(thisConfig.RedisAddr) // TODO: Make this not global

// var maxConnections = thisConfig.MaxConnections

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/getval/{key}", GetValueFromKeyHandler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", thisConfig.ProxyPort), r))
}
