package api

import (
	"github.com/ahsen17/BlogServ/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

func init() {

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
