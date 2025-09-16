package lruKit

import (
	"fmt"
	"testing"
	"time"
)

func TestNewCache(t *testing.T) {
	cache, err := NewCache[int, any](8)
	if err != nil {
		panic(err)
	}

	for i := 0; i < 16; i++ {
		cache.Add(i, nil)
	}
	fmt.Println(cache.Len()) // 8
}

func TestNewExpirableCache(t *testing.T) {
	// make cache with 10ms TTL and 5 max keys
	cache := NewExpirableCache[string, string](5, nil, time.Millisecond*10)

	// set value under key1.
	cache.Add("key1", "val1")

	// get value under key1
	r, ok := cache.Get("key1")

	// check for OK value
	if ok {
		fmt.Printf("value before expiration is found: %v, value: %q, length: %d\n", ok, r, cache.Len())
	}

	// wait for cache to expire
	time.Sleep(time.Millisecond * 12)

	// get value under key1 after key expiration
	r, ok = cache.Get("key1")
	fmt.Printf("value after expiration is found: %v, value: %q, length: %d\n", ok, r, cache.Len())

	// set value under key2, would evict old entry because it is already expired.
	cache.Add("key2", "val2")

	fmt.Printf("Cache len: %d\n", cache.Len())
}
