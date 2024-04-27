/*
  Package router
  @Author: Ahsen17
  @Github: https://github.com/Ahsen17
  @Time:
  @Description: 平台路由管理
*/

package router

import (
	"net/http"

	"github.com/ahsen17/BlogServ/src/api"
	"github.com/gin-gonic/gin"

	swaggerDocs "github.com/ahsen17/BlogServ/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// 参考下开源项目RouterGroup的用法

// InitRouter 初始化应用路由
func InitRouter(engine *gin.Engine) {
	// 系统路由
	mainGrp := engine.Group("")
	{
		mainGrp.GET("/", MainHandler)
		mainGrp.GET("/404", NotFoundHandler)
		mainGrp.GET("/500", ServerErrorHandler)
	}

	// 应用接口
	apiV1Grp := "/api/v1"
	swaggerDocs.SwaggerInfo.BasePath = apiV1Grp
	// 接口version 1
	v1Grp := engine.Group(apiV1Grp)
	{
		v1Grp.GET("/example", api.SwaggerExample)
	}

	// 其他路由
	{
		// swagger接口文档
		engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
}

// MainHandler 主页
func MainHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"success": true,
		"message": "无内容",
		"data":    nil,
	})
}

// NotFoundHandler 404处理
func NotFoundHandler(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"code":    http.StatusNotFound,
		"success": false,
		"message": "未找到相关内容",
		"data":    nil,
	})
}

func ServerErrorHandler(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"code":    http.StatusInternalServerError,
		"success": false,
		"message": "系统错误",
		"data":    nil,
	})
}
