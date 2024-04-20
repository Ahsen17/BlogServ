package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 参考下开源项目RouterGroup的用法

// InitRouter 初始化应用路由
func InitRouter(engine *gin.Engine) {
	mainGrp := engine.Group("")
	mainGrp.GET("/", MainHandler)
	mainGrp.GET("/404", NotFoundHandler)
	mainGrp.GET("/500", ServerErrorHandler)

	// 应用接口
	apiGrp := engine.Group("/api")
	apiGrp.Any("", NotFoundHandler)
}

// MainHandler 主页
func MainHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "无内容",
		"data":    nil,
	})
}

// NotFoundHandler 404处理
func NotFoundHandler(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"code":    http.StatusNotFound,
		"message": "未找到相关内容",
		"data":    nil,
	})
}

func ServerErrorHandler(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"code":    http.StatusInternalServerError,
		"message": "系统错误",
		"data":    nil,
	})
}
