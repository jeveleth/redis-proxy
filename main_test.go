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
		routeKey string
		routeVal string
	}{
		{"key1", "val1"},
		{"key2", "val2"},
		{"key3", "val3"},
		{"key4", "val4"},
		{"key5", "val5"},
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
		expected := fmt.Sprintf("From Redis: %v => %v", tc.routeKey, tc.routeVal)
		actual := rr.Body.String()
		if actual != expected {
			t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
		}
	}
	tearDownCache()
}

func TestGetValueFromLocalCacheSuccess(t *testing.T) {
	setupKeysInCache(t)

	nt := []struct {
		routeKey string
		routeVal string
	}{
		{"key1", "val1"},
		{"key2", "val2"},
	}

	for _, tc := range nt {
		path := fmt.Sprintf("/getval/%s", tc.routeKey)
		req, err := http.NewRequest("GET", path, nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc("/getval/{key}", GetValueFromKeyHandler)
		router.ServeHTTP(rr, req)
		expected := fmt.Sprintf("From cache: %v => %v", tc.routeKey, tc.routeVal)
		actual := rr.Body.String()
		if actual != expected {
			t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
		}
	}
}

func TestCacheEvictsWithLRU(t *testing.T) {
	testCache := newCache(5)
	nt := []struct {
		routeKey string
		routeVal string
	}{
		{"testKey1", "testVal1"},
		{"testKey2", "testVal2"},
		{"testKey3", "testVal3"},
		{"testKey4", "testVal4"},
		{"testKey5", "testVal5"},
		{"testKey6", "testVal6"},
	}
	for _, tc := range nt {
		testCache.client.Add(tc.routeKey, tc.routeVal)
	}
	for _, key := range testCache.client.Keys() {
		if key == "testKey1" {
			t.Errorf("key %v should be evicted", key)
		}
	}
}

func setupKeysInCache(t *testing.T) {
	tt := []struct {
		routeKey string
		routeVal string
	}{
		{"key1", "val1"},
		{"key2", "val2"},
	}

	for _, tc := range tt {
		localCache.setCVal(tc.routeKey, tc.routeVal)
	}

}

func tearDownCache() {
	localCache.client.Purge()
}

// TODO: Test that getting value:
// --> gets a value from the cache

// --> evicts a key with expiry time
// 	Fixed key size
// 	Sequential concurrent processing
// 	Configuration
// 	Parallel concurrent processing
// 	Redis client protocol # ran out of time
