package persistence

import (
	"testing"
	"time"
)

// These tests require memcached running on localhost:11211 (the default)
const testServer = "localhost:11211"

var newMemcachedStore = func(t *testing.T, defaultExpiration time.Duration) CacheStore {
	flushMemcached(t, testServer)
	return NewMemcachedStore([]string{testServer}, defaultExpiration)
}

func TestMemcachedCache_TypicalGetSet(t *testing.T) {
	typicalGetSet(t, newMemcachedStore)
}

func TestMemcachedCache_IncrDecr(t *testing.T) {
	incrDecr(t, newMemcachedStore)
}

func TestMemcachedCache_Expiration(t *testing.T) {
	expiration(t, newMemcachedStore)
}

func TestMemcachedCache_EmptyCache(t *testing.T) {
	emptyCache(t, newMemcachedStore)
}

func TestMemcachedCache_Replace(t *testing.T) {
	testReplace(t, newMemcachedStore)
}

func TestMemcachedCache_Add(t *testing.T) {
	testAdd(t, newMemcachedStore)
}
