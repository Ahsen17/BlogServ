package data

import (
	"testing"
)

func TestCache(t *testing.T) {
	client := CacheClient()
	defer client.Close()

	client.Del("key").Err()
}
