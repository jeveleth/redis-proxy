package main

import (
	"flag"
	"time"
)

type Config struct {
	RedisAddr       string
	CacheExpiryTime time.Duration
	CacheCapacity   int
	ProxyPort       int
}

func MustLoadConfig() Config {
	config := Config{}
	flag.StringVar(&config.RedisAddr, "redis-addr", "redis:6379", "Address of the backing Redis")
	flag.DurationVar(&config.CacheExpiryTime, "cache-expiry-time", time.Minute, "Cache expiry time")
	flag.IntVar(&config.CacheCapacity, "cache-capacity", 2, "Capacity (number of keys)")
	flag.IntVar(&config.ProxyPort, "proxy-port", 8080, "TCP/IP port number the proxy listens on")
	flag.Parse()
	return config
}