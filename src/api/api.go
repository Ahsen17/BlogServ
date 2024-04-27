/*
  Package api
  @Author: Ahsen17
  @Github: https://github.com/Ahsen17
  @Time:
  @Description: 平台接口管理
*/

package api

import (
	"github.com/ahsen17/BlogServ/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

func init() {

}

func NoContentHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"success": true,
		"message": "无内容",
		"data":    nil,
	})
}

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

// BasePath在同一个模块中定义一次即可

//	@BasePath	/api/v1

// SwaggerExample 接口测试文档
//
//	@Summary	接口测试
//	@Schemes
//	@Description	接口测试
//	@Tags			example
//	@Accept			json
//	@Produce		json
//	@Success		200	{string}	Hello world
//	@Router			/example/hello [get]
func SwaggerExample(g *gin.Context) {
	logger.Info("测试日志")
	g.JSON(http.StatusOK, "swagger example")
}
