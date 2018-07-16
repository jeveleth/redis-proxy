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

// setRVal sets a value in Redis from a given key
func (r *redisClient) setRVal(key string, value string) {
	err := r.client.Set(key, value, 0).Err()
	if err != nil {
		log.Printf("Error setting Value: %v for Key: %v. Error is: %v.", value, key, err)
	}
}

// getRVal returns a key's value from Redis; otherwise it returns a nil error
func (r *redisClient) getRVal(key string) (string, error) {
	// Note: As the project seeks only a Redis GET,
	// This retrieves only strings, not sets, lists, hashes, or other Redis types.
	val, err := r.client.Get(key).Result()

	if err != nil {
		log.Printf("Redis can't find Key: %v. Error is: %v.", key, err)
		return "", err
	}

	return val, nil
}
