package src

import (
	"fmt"
	"github.com/ahsen17/BlogServ/config"
	"github.com/ahsen17/BlogServ/logger"
	"github.com/ahsen17/BlogServ/src/router"
	"github.com/gin-gonic/gin"
)

type ServMgr struct {
}

func (mgr *ServMgr) RunServer() {
	engine := gin.Default()
	engine.SetTrustedProxies([]string{"localhost"})

	routesMgr := router.RoutesMgr{Engine: engine}
	// 注册中间件
	routesMgr.RegisterMiddlewares(
		GlobalUserAuthMiddleware,
	)
	// 初始化路由
	routesMgr.CollectRouters()

	// 启动服务
	serverConfig := config.ServerConfig()
	err := engine.Run(fmt.Sprintf("%s:%d", serverConfig.Address, serverConfig.Port))
	if err != nil {
		logger.Errorf("服务启动失败: %s", err)
		panic(err)
	}
}
