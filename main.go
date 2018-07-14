package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-redis/redis"
)

// TODO: Clients interface to the Redis proxy through HTTP, with the Redis “GET” command mapped to the HTTP “GET” method.
// -- Note that the proxy still uses the Redis protocol to communicate with the backend Redis server.

// TODO: How do I query the redis server to confirm that info was set? Not just in the cache.
// TODO: How do I build for testing?
// TODO: What am I trying to test?
// TODO: Initialize a struct that the functions act on so they don't have to pass in the client each time

var redisClient = newRedisClient()

func newRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Password: os.Getenv("REDIS_PW"),
		DB:       0,
	})
}

func setKeyVal(key string, value string) {
	err := redisClient.Set(key, value, 0).Err()
	if err != nil {
		log.Printf("error setting value %v for key %v. error is: %#v", value, key, err)
	}
}

func getVal(key string) string {
	val, err := redisClient.Get(key).Result()
	if err != nil {
		log.Printf("error getting value for key %v. error is: %#v", key, err)
	}
	return val
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Connecting to redis server!")
	fmt.Fprintf(w, "\nGetting value for %s", r.URL.Path[1:])
	// TODO: escape input? Does redis have same risks as SQL injection?
	getVal(r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
