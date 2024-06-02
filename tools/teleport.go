package tools

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

type ServTool struct {
}

// FetchRemoteIp 获取访问IP
func (st ServTool) FetchRemoteIp(ctx *gin.Context) string {
	reqIP := ctx.ClientIP()
	if reqIP == "::1" {
		reqIP = "127.0.0.1"
	}
	return reqIP
}

// GenerateAccessKey 生成访问密钥
func (st ServTool) GenerateAccessKey(username string, ipAddress string) (string, error) {
	plainText := fmt.Sprintf("%s@%s", username, ipAddress)
	return EncryptAES(plainText)
}

// DecryptAccessKey 解密访问密钥
func (st ServTool) DecryptAccessKey(token string) (string, error) {
	authStr, err := DecryptAES(token)
	if err != nil {
		return "", err
	}
	return strings.Split(authStr, "_")[0], nil
}
