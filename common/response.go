package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	SuccessAccess = "访问成功"
	FailedAccess  = "访问失败"
	ServerError   = "服务器响应失败"
)

type Response struct {
	Ctx     *gin.Context
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (r Response) BASE(code int, msg string, data interface{}) {
	r.Code = code
	r.Message = msg
	r.Data = data
	r.Ctx.JSON(code, r)
}

// OK 访问服务成功
func (r Response) OK(msg string, data interface{}) {
	if msg == "" {
		msg = SuccessAccess
	}
	r.BASE(http.StatusOK, msg, data)
}

// FAIL 访问服务失败
func (r Response) FAIL(msg string, data interface{}) {
	if msg == "" {
		msg = FailedAccess
	}
	r.BASE(http.StatusBadRequest, msg, data)
}

// ERROR 应用程序内部异常
func (r Response) ERROR(msg string, data interface{}) {
	if msg == "" {
		msg = ServerError
	}
	r.BASE(http.StatusInternalServerError, msg, data)
}
