/*
  Package server
  @Author: Ahsen17
  @Github: https://github.com/Ahsen17
  @Time: 2024/6/2 1:57
  @Description: ...
*/

package teleport

import (
	com "github.com/ahsen17/BlogServ/common"
	"github.com/ahsen17/BlogServ/src/models"
	"github.com/ahsen17/BlogServ/tools"
	"github.com/gin-gonic/gin"
)

// OnRegisterRequest 注册平台账户
func OnRegisterRequest(ctx *gin.Context) {
	accMgr := models.AccountMgr{DBClient: dbClient, Cache: cache}
	resp := com.ResponseMgr{Ctx: ctx}

	if err := ctx.ShouldBindJSON(&accMgr.Account); err != nil {
		resp.ERROR("获取JSON参数错误", nil)
		return
	}

	if ok := accMgr.Register(); !ok {
		resp.ERROR("注册失败", nil)
	} else {
		resp.OK("注册成功", &accMgr.Account)
	}
}

// OnLoginRequest 账户登录
func OnLoginRequest(ctx *gin.Context) {
	accMgr := models.AccountMgr{DBClient: dbClient, Cache: cache}
	resp := com.ResponseMgr{Ctx: ctx}

	token := ctx.Request.Header.Get("Authorization")
	if cache.Get(token).Err() == nil {
		// 已登录
		resp.FAIL("请勿重复登录", nil)
		return
	}

	if err := ctx.ShouldBindJSON(&accMgr.Account); err != nil {
		resp.ERROR("获取账户信息异常", nil)
		return
	}

	clientIP := tools.ServTool{}.FetchRemoteIp(ctx)
	if ok, msg := accMgr.Login(clientIP); !ok {
		resp.ERROR(msg, nil)
	} else {
		resp.OK("登陆成功", msg)
	}
}

// OnLogoutRequest 注销登录
func OnLogoutRequest(ctx *gin.Context) {
	resp := com.ResponseMgr{Ctx: ctx}
	token := ctx.Request.Header.Get("Authorization")
	if cache.Get(token).Err() != nil {
		resp.FAIL("请先登录", nil)
	} else if err := cache.Del(token).Err(); err != nil {
		resp.ERROR("注销失败", nil)
	} else {
		resp.OK("注销成功", nil)
	}
}
