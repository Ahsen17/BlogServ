/*
  Package teleport
  @Author: Ahsen17
  @Github: https://github.com/Ahsen17
  @Time: 2024/6/2 21:58
  @Description: ...
*/

package teleport

import (
	com "github.com/ahsen17/BlogServ/common"
	models "github.com/ahsen17/BlogServ/src/models"
	"github.com/gin-gonic/gin"
)

// InitStorage 初始化数据库表
func InitStorage(ctx *gin.Context) {
	resp := com.ResponseMgr{Ctx: ctx}

	// 初始化数据库
	if err := dbClient.AutoMigrate(
		&models.Account{},
		&models.User{},
		//&roleTbl,
	); err != nil {
		resp.ERROR("数据库表初始化失败", nil)
		return
	}

	resp.OK("数据库表初始化成功", nil)
}
