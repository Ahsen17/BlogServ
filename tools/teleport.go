package tools

import (
	"github.com/gin-gonic/gin"
)

type ServTool struct {
}

// FetchIpAddress 获取IP
func (st *ServTool) FetchIpAddress(ctx *gin.Context) string {
	return ctx.ClientIP()
}

func (st *ServTool) GenerateAccessKey() string {
	// TODO: 生成登录后访问秘钥
	return ""
}
