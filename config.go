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
	// MaxConnections  int
}

func MustLoadConfig() Config {
	config := Config{}
	flag.StringVar(&config.RedisAddr, "redis-addr", "redis:6379", "Address of the backing Redis")
	flag.DurationVar(&config.CacheExpiryTime, "cache-expiry-time", time.Hour, "Cache expiry time")
	flag.IntVar(&config.CacheCapacity, "cache-capacity", 1000, "Capacity (number of keys)")
	flag.IntVar(&config.ProxyPort, "proxy-port", 8080, "TCP/IP port number the proxy listens on")
	// flag.IntVar(&config.MaxConnections, "max-conn", 8080, "Maximum limit of connections to proxy")
	flag.Parse()
	return config
}
