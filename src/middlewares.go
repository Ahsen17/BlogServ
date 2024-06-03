/*
  Package src
  @Author: Ahsen17
  @Github: https://github.com/Ahsen17
  @Time: 2024/5/1 上午4:24
  @Description: 中间件
*/

package src

import (
	"github.com/ahsen17/BlogServ/common"
	"github.com/ahsen17/BlogServ/src/data"
	"github.com/ahsen17/BlogServ/tools"
	"github.com/gin-gonic/gin"
	"strings"
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

func requestAuthenticate(ctx *gin.Context) {

}

// GlobalLoginStatusMiddleware 登录状态检查中间件
func GlobalLoginStatusMiddleware(ctx *gin.Context) {
	requestUri := ctx.Request.URL.Path
	token := ctx.Request.Header.Get("Authorization")

	// 无需检查是否登录
	if uriBingo(requestUri) {
		return
	}

	if cache.Get(token).Err() != nil {
		// 在线状态失效，重定向登录
		ctx.Redirect(302, loginUri)
	}
}

// GlobalUserAuthMiddleware 平台用户鉴权
func GlobalUserAuthMiddleware(ctx *gin.Context) {
	requestUri := ctx.Request.URL.Path
	// 无需鉴权
	if uriBingo(requestUri) {
		return
	}

	servTool := tools.ServTool{}
	resp := common.ResponseMgr{Ctx: ctx}
	token := ctx.Request.Header.Get("Authorization")

	// 判断请求来源IP与token中的IP是否相符合
	requestIp := servTool.FetchRemoteIp(ctx)
	usernameWithIP, _ := servTool.DecryptAccessKey(token)
	cacheIp := strings.Split(usernameWithIP, "@")[1]
	if cacheIp != requestIp {
		ctx.Abort()
		resp.FAIL("非法访问", nil)
	}

	// TODO: 获取用户身份，判断是否有权限调用该uri
	// TODO: 用户登录还有点问题，需要验证不同IP的非法访问拦截
}
