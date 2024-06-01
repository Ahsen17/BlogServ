package data

import (
	"fmt"
	"github.com/ahsen17/BlogServ/logger"

	"github.com/ahsen17/BlogServ/config"
	"github.com/go-redis/redis"
)

//cache 选择redis

var (
	cache *redis.Client

	serverKey = fmt.Sprintf("%s@%s:%d",
		"blogserv",
		config.ServerConfig().Address,
		config.ServerConfig().Port,
	)
)

func init() {
	c := config.CacheConfig()

	cache = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", c.Host, c.Port),
		Password: c.Password,
		DB:       c.Database,
	})

	if err := cache.DebugObject(serverKey).Err(); err != nil {
		//if err := cache.Set(serverKey, 1, 0).Err(); err != nil {
		logger.Fatalf("缓存连接池初始化失败: %s", err)
		panic(err)
	}

	logger.Info("缓存连接池初始化成功")
}

func CacheClient() *redis.Client {
	return cache
}
