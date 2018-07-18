package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

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
	tearDownRedis()
	tearDownCache()
}

func TestHandlerAcceptsMaxConnections(t *testing.T) {
	req, err := http.NewRequest("GET", "", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	hf := http.HandlerFunc(HomeHandler)
	// newServer := httptest.NewServer(hf)

	hf.ServeHTTP(rr, req)

	// // set MaxConnections to 5
	// // open 10 connections
	// rr := httptest.NewRecorder()
	// hf := http.HandlerFunc(HomeHandler)
	// for index := 0; index < 150; index++ {
	// 	hf.ServeHTTP(rr, req)
	// }

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
	tearDownRedis()
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
	tearDownRedis()
	tearDownCache()
}

func TestCacheEmptyWithShortTTL(t *testing.T) {
	testCache := newCache(50)
	m, _ := time.ParseDuration("5s")
	testVal := valWithTTL{
		Value:     "testVal1",
		TTL:       m,
		CreatedAt: time.Now().Add(time.Second * -30)}

	nt := []struct {
		routeKey string
		routeVal interface{}
	}{
		{"testKey1", testVal},
		{"testKey2", testVal},
		{"testKey3", testVal},
		{"testKey4", testVal},
	}
	for _, tc := range nt {
		testCache.client.Add(tc.routeKey, tc.routeVal)
		result, _ := testCache.getCVal(tc.routeKey)
		actual := result
		if actual != nil {
			t.Errorf("handler returned unexpected body: got %v want nil", actual)
		}
	}
	tearDownRedis()
	tearDownCache()
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
	tearDownRedis()
	tearDownCache()
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

func tearDownRedis() {
	myRedisClient.client.FlushDB()
}
