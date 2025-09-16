package lruKit

import (
	"time"

	lru "github.com/hashicorp/golang-lru/v2"
	"github.com/hashicorp/golang-lru/v2/expirable"
)

// NewCache 缓存（不带TTL）.
func NewCache[K comparable, V any](size int) (*lru.Cache[K, V], error) {
	return lru.New[K, V](size)
}

// NewExpirableCache 缓存（带TTL）.
/*
@param size (1) 缓存大小（max keys）
			(2) 0: unlimited size
@param ttl	Providing 0 TTL turns expiring off.
*/
func NewExpirableCache[K comparable, V any](size int, onEvict expirable.EvictCallback[K, V], ttl time.Duration) *expirable.LRU[K, V] {
	return expirable.NewLRU[K, V](size, onEvict, ttl)
}
