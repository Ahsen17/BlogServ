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
	data "github.com/ahsen17/BlogServ/src/data"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	dbClient = data.DBClient()
	cache    = data.CacheClient()
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
