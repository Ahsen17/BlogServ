package main

import (
	"github.com/ahsen17/BlogServ/logger"
	"github.com/ahsen17/BlogServ/src"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.SetTrustedProxies([]string{"localhost"})

	src.InitRouter(r)
	err := r.Run(":9090")
	if err != nil {
		logger.Errorf("服务启动失败: %s", err)
		panic(err)
	}
}
