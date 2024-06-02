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
	"github.com/gin-gonic/gin"
)

// RegisterAccountRequest 注册平台账户
func RegisterAccountRequest(ctx *gin.Context) {
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

	if err := ctx.ShouldBindJSON(&accMgr.Account); err != nil {
		resp.ERROR("获取账户信息异常", nil)
		return
	}

	if ok, msg := accMgr.Login(ctx); !ok {
		resp.ERROR(msg, nil)
	} else {
		resp.OK(msg, &accMgr.Account)
	}
}

// OnLogoutRequest 注销登录
func OnLogoutRequest(ctx *gin.Context) {
	accMgr := models.AccountMgr{DBClient: dbClient, Cache: cache}
	resp := com.ResponseMgr{Ctx: ctx}

	if err := ctx.ShouldBindJSON(&accMgr.Account); err != nil {
		resp.ERROR("获取账户信息异常", nil)
		return
	}

	if ok := accMgr.Logout(); !ok {
		resp.ERROR("退出登录异常", nil)
	} else {
		resp.OK("已注销", &accMgr.Account)
	}
}
