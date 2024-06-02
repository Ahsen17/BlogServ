package src

import (
	"fmt"
	"github.com/ahsen17/BlogServ/config"
	"github.com/ahsen17/BlogServ/logger"
	"github.com/ahsen17/BlogServ/router"
	"github.com/gin-gonic/gin"
)

type ServMgr struct {
}

func (mgr *ServMgr) RunServer() {
	engine := gin.Default()
	engine.SetTrustedProxies([]string{"localhost"})

	// 初始化系统路由
	router.CollectRouters(engine)

	// 启动服务
	serverConfig := config.ServerConfig()
	err := engine.Run(fmt.Sprintf("%s:%d", serverConfig.Address, serverConfig.Port))
	if err != nil {
		logger.Errorf("服务启动失败: %s", err)
		panic(err)
	}
}
