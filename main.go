package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/gorilla/mux"
	"golang.org/x/net/netutil"
)

var thisConfig = MustLoadConfig()
var localCache = newCache(thisConfig.CacheCapacity)
var myRedisClient = newRedisClient(thisConfig.RedisAddr)

func getRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/getval/{key}", GetValueFromKeyHandler)
	return r
}

func main() {
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", thisConfig.ProxyPort))
	if err != nil {
		log.Fatalf("Listen: %v", err)
	}
	defer l.Close()
	l = netutil.LimitListener(l, thisConfig.MaxConnections)
	log.Fatal(http.Serve(l, getRoutes()))
}
