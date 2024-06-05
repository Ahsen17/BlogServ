/*
  Package src
  @Author: Ahsen17
  @Github: https://github.com/Ahsen17
  @Time: 2024/5/1 上午4:24
  @Description: 中间件
*/

package src

import (
	"github.com/ahsen17/BlogServ/logger"
	"github.com/ahsen17/BlogServ/src/data"
	"github.com/gin-gonic/gin"
)

var (
	cache    = data.CacheClient()
	dbClient = data.DBClient()
)

var (
	// 平台内权限用户操作
	loginUri = "/teleport/account/login"

	// 非平台匿名用户操作

	// 无需鉴权URI，包括匿名用户URI
	passUris = []string{
		loginUri,
	}
)

func uriBingo(uri string) bool {
	for _, v := range passUris {
		if v == uri {
			return true
		}
	}
	return false
}

// GlobalUserAuthMiddleware 平台用户鉴权
func GlobalUserAuthMiddleware(ctx *gin.Context) {
	requestUri := ctx.Request.RequestURI
	if uriBingo(requestUri) {
		// 无需鉴权
		return
	}

	logger.Info("执行鉴权操作")
}
