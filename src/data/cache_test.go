package data

import (
	"BlogServ/logger"
	"testing"
)

func TestCache(t *testing.T) {
	client := CacheClient()
	defer client.Close()

	//client.Set("key", "11111", 0).Err()
	if val, err := client.Get("key").Result(); err == nil {
		logger.Infof("%s", val)
	} else {
		logger.Errorf("%s ===", err)
	}

	client.Del("key").Err()
}
