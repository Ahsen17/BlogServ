package data

import (
	"BlogServ/config"
	"BlogServ/logger"
	"testing"
)

func TestCache_InitClient(t *testing.T) {
	c := config.NewConfig()
	client := InitCache(&c.Cache)
	defer client.Close()

	//client.Set("key", "11111", 0).Err()
	if val, err := client.Get("key").Result(); err == nil {
		logger.Infof("%s", val)
	} else {
		logger.Errorf("%s ===", err)
	}

	client.Del("key").Err()
}
