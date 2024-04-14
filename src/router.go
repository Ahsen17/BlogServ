package src

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 参考下开源项目RouterGroup的用法

var (
	router *Router
)

type Router struct {
	routes *[]map[string][]string
}

func InitRouter(e *gin.Engine) {
	// 主页
	mainGrp := e.Group("/")
	mainGrp.Any("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "主页无内容",
			"data":    nil,
		})
	})
}
