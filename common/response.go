/*
  Package common
  @Author: Ahsen17
  @Github: https://github.com/Ahsen17
  @Time:
  @Description: Json响应体
*/

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
	Code    int         `json:"code"`
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponseMgr struct {
	Ctx      *gin.Context
	Response Response
}

func (mgr ResponseMgr) BASE(code int, success bool, msg string, data interface{}) {
	mgr.Response.Code = code
	mgr.Response.Success = success
	mgr.Response.Message = msg
	mgr.Response.Data = data
	mgr.Ctx.JSON(code, &mgr.Response)
}

// OK 访问服务成功
func (mgr ResponseMgr) OK(msg string, data interface{}) {
	if msg == "" {
		msg = SuccessAccess
	}
	mgr.BASE(http.StatusOK, true, msg, data)
}

// FAIL 访问服务失败
func (mgr ResponseMgr) FAIL(msg string, data interface{}) {
	if msg == "" {
		msg = FailedAccess
	}
	mgr.BASE(http.StatusBadRequest, false, msg, data)
}

// ERROR 应用程序内部异常
func (mgr ResponseMgr) ERROR(msg string, data interface{}) {
	if msg == "" {
		msg = ServerError
	}
	mgr.BASE(http.StatusInternalServerError, false, msg, data)
}
