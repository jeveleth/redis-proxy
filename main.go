// // TODO: Clients interface to the Redis proxy through HTTP, with the Redis “GET” command mapped to the HTTP “GET” method.
// // -- Note that the proxy still uses the Redis protocol to communicate with the backend Redis server.

// // TODO: How do I query the redis server to confirm that info was set? Not just in the cache.
// // TODO: How do I build for testing?
// // TODO: What am I trying to test?
// // TODO: Initialize a struct that the functions act on so they don't have to pass in the client each time

package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
)

var redisClient = newRedisClient()

func newRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Password: os.Getenv("REDIS_PW"),
		DB:       0,
	})
}

func setVal(key string, value string) {
	err := redisClient.Set(key, value, 0).Err()
	if err != nil {
		log.Printf("error setting value %v for key %v. error is: %#v", value, key, err)
	}
	log.Printf("Setting %v to store %v", key, value)
}

func getVal(key string) string {
	val, err := redisClient.Get(key).Result()
	if err != nil {
		log.Printf("error getting value for key %v. error is: %#v", key, err)
	}
	return val
}

// HomeHandler serves as the index page/ status page
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the proxy!\n"))
	// TODO: escape input? Does redis have same risks as SQL injection?
}

// SetValueFromKeyHandler will SET value in Redis
func SetValueFromKeyHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	setVal(vars["key"], vars["value"])
}

// GetValueFromKeyHandler will GET value from Redis
func GetValueFromKeyHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Printf("value for %v is %v", vars["key"], getVal(vars["key"]))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/getter/{key}", GetValueFromKeyHandler)
	// TODO: restrict getter to HTTP GET Method
	r.HandleFunc("/setter/{key:[0-9a-zA-Z]+}/{value:[0-9a-zA-Z]+}", SetValueFromKeyHandler)
	// TODO: Better 404 handling?
	log.Fatal(http.ListenAndServe(":8080", r))
}
