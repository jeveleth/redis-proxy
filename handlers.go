package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// HomeHandler serves as the index page/ status page
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the proxy!"))
}

// setValueFromKeyHandler will SET value in Redis
// Although this isn't explicitly asked for, it would be difficult to test
// GETs without confirming SETs. This should work only in tests
// func setValueFromKeyHandler(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	myRedisClient.setRVal(vars["key"], vars["value"])
// }

// GetValueFromKeyHandler will GET value from Redis
func GetValueFromKeyHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	val := localCache.getCVal(vars["key"])
	// log.Printf("DEBUG ==> hello getting %v, vars %v", vars["key"], val)
	// if non-empty, return value
	// If empty, call getVal()
	// If calling getVal(), then call cacheSetVal(),
	// return value
	// val, err := myRedisClient.getRVal
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusServiceUnavailable)
	// }

	switch v := val.(type) {
	case string:
		// return v
		w.Write([]byte(v))
	}
	// return w.Write([]byte(v))
}
