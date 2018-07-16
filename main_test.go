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
	tt := []struct {
		routeVariable string
		shouldPass    bool
	}{
		{"key1", true},
		// {"key2", true},
		// {"key3", true},
	}

	for _, tc := range tt {
		path := fmt.Sprintf("/getter/%s", tc.routeVariable)
		req, err := http.NewRequest("GET", path, nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc("/getter/{type}", GetValueFromKeyHandler)
		router.ServeHTTP(rr, req)

		if rr.Code == http.StatusOK && !tc.shouldPass {
			t.Errorf("handler should have failed on routeVariable %s: got %v want %v",
				tc.routeVariable, rr.Code, http.StatusOK)
		}
	}
}

// func TestSetValueFromKeyHandler(t *testing.T) {
// 	myRedisClient.setRVal("key1", "val1")
// 	tt := []struct {
// 		routeKey   string
// 		routeVal   string
// 		shouldPass bool
// 	}{
// 		{"key1", "value1", true},
// 		{"key2", "value2", true},
// 		{"key3", "value3", true},
// 	}

// 	for _, tc := range tt {
// 		path := fmt.Sprintf("/setter/%s/%s", tc.routeKey, tc.routeVal)
// 		req, err := http.NewRequest("GET", path, nil)
// 		if err != nil {
// 			t.Fatal(err)
// 		}

// 		rr := httptest.NewRecorder()
// 		router := mux.NewRouter()
// 		router.HandleFunc("/setter/{type}", setValueFromKeyHandler)
// 		router.ServeHTTP(rr, req)

// 		if rr.Code == http.StatusOK && !tc.shouldPass {
// 			t.Errorf("handler should have failed on routeVariable %s: got %v want %v",
// 				tc.routeKey, rr.Code, http.StatusOK)
// 		}
// 	}
// }

// TODO: Test that getting value:
// --> gets a value from the cache
// --> gets a value from redis if cache is empty
// --> evicts a key with expiry time
// --> evicts a key with lru
// 	# curl localhost:8080/getter/p
// 	Fixed key size
// 	Sequential concurrent processing
// 	Configuration
// 	Parallel concurrent processing
// 	Redis client protocol # ran out of time
