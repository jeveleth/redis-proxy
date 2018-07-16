package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestHandlerWelcomeMessage(t *testing.T) {
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

func TestGetValueFromRedisWhenLocalCacheEmpty(t *testing.T) {
	tt := []struct {
		routeKey   string
		routeVal   string
		shouldPass bool
	}{
		{"key1", "val1", true},
		{"key2", "val2", false},
		{"key3", "val3", false},
		{"key4", "val4", false},
		{"key5", "val5", false},
	}

	for _, tc := range tt {
		myRedisClient.setRVal(tc.routeKey, tc.routeVal)
		path := fmt.Sprintf("/getval/%s", tc.routeKey)
		req, err := http.NewRequest("GET", path, nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc("/getval/{key}", GetValueFromKeyHandler)
		router.ServeHTTP(rr, req)
		expected := tc.routeVal
		actual := rr.Body.String()
		if actual != expected {
			t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
		}
	}
}
func TestGetValueFromLocalCacheSuccess(t *testing.T) {
	// TODO: Why is local cache not getting value from Redis in test?
	// --> cache gets cleared at test end
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
