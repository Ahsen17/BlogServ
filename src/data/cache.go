package data

import (
	"BlogServ/config"
	"fmt"
	"github.com/go-redis/redis"
)

//cache 选择redis

var (
	client *redis.Client
)

func InitClient(c *config.Cache) *redis.Client {
	if client != nil {
		return client
	}

	lock.Lock()
	defer lock.Unlock()

	client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", c.Host, c.Port),
		Password: c.Password,
		DB:       c.Database,
	})

	return client
}
