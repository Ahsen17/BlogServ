package data

import (
	"fmt"

	"github.com/ahsen17/BlogServ/config"
	"github.com/go-redis/redis"
)

//cache 选择redis

var (
	cache *redis.Client
)

func init() {
	c := config.CacheConfig()

	cache = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", c.Host, c.Port),
		Password: c.Password,
		DB:       c.Database,
	})
}

func CacheClient() *redis.Client {
	return cache
}
