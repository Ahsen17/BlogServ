package data

import (
	"BlogServ/config"
	"fmt"
	"github.com/go-redis/redis"
)

//cache 选择redis

var (
	cache *redis.Client
)

func InitCache(c *config.Cache) *redis.Client {
	if cache != nil {
		return cache
	}

	lock.Lock()
	defer lock.Unlock()

	cache = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", c.Host, c.Port),
		Password: c.Password,
		DB:       c.Database,
	})

	return cache
}
