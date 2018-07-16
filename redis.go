package main

import (
	"log"
	"os"

	"github.com/go-redis/redis"
)

type redisClient struct {
	client *redis.Client
}

func newRedisClient(redisAddr string) *redisClient {
	client := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: os.Getenv("REDIS_PW"),
		DB:       0,
	})
	return &redisClient{client: client}
}

func (r *redisClient) setRVal(key string, value string) {
	err := r.client.Set(key, value, 0).Err()
	if err != nil {
		log.Printf("error setting value %v for key %v. error is: %#v", value, key, err)
	}
	log.Printf("Setting %v to store %v", key, value)
}

func (r *redisClient) getRVal(key string) (string, error) {
	// Note: As the assignment asks for only a Redis GET,
	// I'm dealing only with retrieving strings,
	// not sets, lists, hashes, or other Redis types.
	val, err := r.client.Get(key).Result()
	log.Printf("checking redis for getting value for key %v.", key)

	// if key == nil {
	// log.Printf("key is empty key %v.", key)
	// }

	if err != nil {
		log.Printf("error getting value for key %v. error is: %#v", key, err)
		return "", err
	}
	return val, nil
}
