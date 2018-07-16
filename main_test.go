package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	hf := http.HandlerFunc(HomeHandler)

	hf.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `Welcome to the proxy!`
	actual := rr.Body.String()
	if actual != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}
}

func TestGetValueFromKeyHandler(t *testing.T) {
	myRedisClient.setRVal("key1", "val1")

	// TODO: Why is local cache not getting value from Redis in test?
	tt := []struct {
		routeKey   string
		routeVal   string
		shouldPass bool
	}{
		{"key1", "val1", true},
		{"key2", "val2", false},
	}

	for _, tc := range tt {
		path := fmt.Sprintf("/getval/%s", tc.routeKey)
		req, err := http.NewRequest("GET", path, nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc("/getval/{key}", GetValueFromKeyHandler)
		router.ServeHTTP(rr, req)
	}
}

// TODO: Test that getting value:
// --> gets a value from redis if cache is empty
// --> gets a value from the cache
// --> evicts a key with expiry time
// --> evicts a key with lru
// 	Fixed key size
// 	Sequential concurrent processing
// 	Configuration
// 	Parallel concurrent processing
// 	Redis client protocol # ran out of time
