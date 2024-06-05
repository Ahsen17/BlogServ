/*
  Package router
  @Author: Ahsen17
  @Github: https://github.com/Ahsen17
  @Time:
  @Description: 平台路由管理
*/

package router

import (
	com "github.com/ahsen17/BlogServ/common"
	swaggerDocs "github.com/ahsen17/BlogServ/docs"
	"github.com/ahsen17/BlogServ/src/api"
	"github.com/ahsen17/BlogServ/src/teleport"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type RoutesMgr struct {
	Engine *gin.Engine
}

// 参考下开源项目RouterGroup的用法

// NoContentHandler 返回无内容
func NoContentHandler(c *gin.Context) {
	com.ResponseMgr{Ctx: c}.OK("无内容", nil)
}

// NotFoundHandler 资源未找到
func NotFoundHandler(c *gin.Context) {
	com.ResponseMgr{Ctx: c}.NOTFOUND("", nil)
}

// ServerErrorHandler 系统错误
func ServerErrorHandler(c *gin.Context) {
	com.ResponseMgr{Ctx: c}.ERROR("", nil)
}

// CollectRouters 初始化应用路由
func (rm RoutesMgr) CollectRouters() {
	engine := rm.Engine
	// 系统路由
	mainGrp := engine.Group("")
	{
		mainGrp.GET("/", NoContentHandler)
		mainGrp.GET("/404", NotFoundHandler)
		mainGrp.GET("/500", ServerErrorHandler)
	}

	collectTeleportRouters(engine)
	collectDataRouters(engine)
	collectSearchRouters(engine)
	collectModuleRouters(engine)
	collectApiRouters(engine)
}

func (rm RoutesMgr) RegisterMiddlewares(middleware ...gin.HandlerFunc) {
	engine := rm.Engine
	engine.Use(middleware...)
}

func collectTeleportRouters(e *gin.Engine) {
	// 通信服务路由
	teleportGrp := e.Group("/teleport")
	{
		teleportGrp.GET("", NoContentHandler)
	}

	// 账户操作路由
	accountGrp := teleportGrp.Group("/account")
	{
		accountGrp.POST("/register", teleport.OnRegisterRequest)
		accountGrp.POST("/login", teleport.OnLoginRequest)
		accountGrp.POST("/logout", teleport.OnLogoutRequest)
	}
}

func collectDataRouters(e *gin.Engine) {
	// 数据服务路由
	dataGrp := e.Group("/data")
	{
		dataGrp.GET("", NoContentHandler)
		dataGrp.POST("/storage/init", teleport.InitStorage)
	}
}

func collectSearchRouters(e *gin.Engine) {
	// 搜索服务路由
	searchGrp := e.Group("/search")
	{
		searchGrp.GET("", NoContentHandler)
	}
}

func collectModuleRouters(e *gin.Engine) {
	// 模块服务路由
	moduleGrp := e.Group("/module")
	{
		moduleGrp.GET("", NoContentHandler)
	}
}

func collectApiRouters(e *gin.Engine) {
	// API路由
	// 接口文档: /api/docs/index.html
	e.GET("/api/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// 应用接口
	apiV1Grp := "/api/v1"
	swaggerDocs.SwaggerInfo.BasePath = apiV1Grp
	// 接口version 1st.
	v1Grp := e.Group(apiV1Grp)
	v1Grp.GET("/example", api.SwaggerExample)
}
