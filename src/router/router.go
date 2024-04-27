/*
  Package router
  @Author: Ahsen17
  @Github: https://github.com/Ahsen17
  @Time:
  @Description: 平台路由管理
*/

package router

import (
	"github.com/ahsen17/BlogServ/src/api"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	swaggerDocs "github.com/ahsen17/BlogServ/docs"
)

// 参考下开源项目RouterGroup的用法

// InitRouter 初始化应用路由
func InitRouter(engine *gin.Engine) {
	// 系统路由
	mainGrp := engine.Group("")
	{
		mainGrp.GET("/", api.NoContentHandler)
		mainGrp.GET("/404", api.NotFoundHandler)
		mainGrp.GET("/500", api.ServerErrorHandler)
	}

	// 数据路由
	dataGrp := engine.Group("/data")
	{
		dataGrp.GET("", api.NoContentHandler)

		dataGrp.Group("/account")

		dataGrp.Group("/log")

		dataGrp.Group("/module")

		dataGrp.Group("/role")

		userGrp := dataGrp.Group("/user")
		userGrp.GET("/detail", api.NoContentHandler)
	}

	// 模块路由
	moduleGrp := engine.Group("/module")
	{
		moduleGrp.GET("", api.NoContentHandler)
	}

	// 搜索路由
	searchGrp := engine.Group("/search")
	{
		searchGrp.GET("", api.NoContentHandler)
	}

	// 通信路由
	teleportGrp := engine.Group("/teleport")
	{
		teleportGrp.GET("", api.NoContentHandler)
	}

	// 接口文档: /api/docs/index.html
	engine.GET("/api/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// 应用接口
	apiV1Grp := "/api/v1"
	swaggerDocs.SwaggerInfo.BasePath = apiV1Grp
	// 接口version 1st.
	v1Grp := engine.Group(apiV1Grp)
	v1Grp.GET("/example", api.SwaggerExample)
}
