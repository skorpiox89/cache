package persistence

import (
	"net"
	"testing"
	"time"
)

func requireCacheService(t *testing.T, address string) {
	t.Helper()

	conn, err := net.DialTimeout("tcp", address, time.Second)
	if err != nil {
		t.Skipf("cache service unavailable at %s: %v", address, err)
	}
	_ = conn.Close()
}

func flushMemcached(t *testing.T, address string) {
	t.Helper()

	conn, err := net.DialTimeout("tcp", address, time.Second)
	if err != nil {
		t.Skipf("cache service unavailable at %s: %v", address, err)
	}
	defer conn.Close()

	if _, err := conn.Write([]byte("flush_all\r\n")); err != nil {
		t.Fatalf("flush memcached at %s: %v", address, err)
	}
}
