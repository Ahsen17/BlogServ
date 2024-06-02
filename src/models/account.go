/*
  Package models
  @Author: Ahsen17
  @Github: https://github.com/Ahsen17
  @Time:
  @Description: 账户管理
*/

package models

import (
	"fmt"
	"github.com/ahsen17/BlogServ/logger"
	"github.com/ahsen17/BlogServ/tools"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
	"strings"
	"time"
)

const (
	TableAccount = "account"
	ExpireTime   = 12 * 3600
)

const (
	ACTIVE = iota + 1
	DEACTIVATE
	REVOKE
	BANED = 9
)

type Account struct {
	gorm.Model
	Username      string `json:"username" gorm:"unique;not null;size:12"`
	Password      string `json:"password" gorm:"not null;size:32"`
	Status        int    `json:"status"`
	CreateBy      string `json:"create_by"`
	UpdateBy      string `json:"update_by"`
	LastLoginTime int64  `json:"login_time"`
	LastLoginIp   string `json:"login_ip"`
}

type AccountMgr struct {
	Account      *Account
	UserMgr      *UserMgr
	RoleMgr      *RoleMgr
	CustomLogMgr *CustomLogMgr // 用户操作日志

	DBClient *gorm.DB
	Cache    *redis.Client
}

func (a Account) TableName() string {
	return TableAccount
}

func (mgr *AccountMgr) Exists() bool {
	var count int64
	mgr.DBClient.Table(TableAccount).Where("username = ?", mgr.Account.Username).Count(&count)
	return count > 0
}

func (mgr *AccountMgr) CheckPassword() bool {
	return mgr.DBClient.Table(TableAccount).Where(
		"username = ? AND password = ?",
		mgr.Account.Username,
		mgr.Account.Password,
	).First(&mgr.Account).RowsAffected > 0
}

func (mgr *AccountMgr) IfLoginIn() bool {
	exists, err := mgr.Cache.Exists(mgr.Account.Username).Result()
	return err == nil && exists != 0
}

func (mgr *AccountMgr) Register() bool {
	if mgr.Exists() {
		logger.Errorf("用户[%s]已存在", mgr.Account.Username)
		return false
	}
	mgr.Account.Status = DEACTIVATE // 刚注册默认未激活
	if err := mgr.DBClient.Table(TableAccount).Create(&mgr.Account).Error; err != nil {
		logger.Errorf("注册用户[%s]失败: %s", mgr.Account.Username, err)
		return false
	}
	return true
}

func (mgr *AccountMgr) Login(ctx *gin.Context) (bool, string) {
	username := mgr.Account.Username

	// 判断账户是否存在/密码是否正确
	if !mgr.Exists() || !mgr.CheckPassword() {
		return false, fmt.Sprintf("账户[%s]不存在或密码错误", username)
	}

	// 判断是否已登录
	if mgr.IfLoginIn() {
		return false, "请勿重复登录"
	}

	// 判断是否可用
	if mgr.Account.Status != ACTIVE {
		info := fmt.Sprintf("账户[%s]已被禁用", username)
		logger.Error(info)
		return false, info
	}

	// 生成访问密钥
	servTool := tools.ServTool{}
	clientIP := servTool.FetchRemoteIp(ctx)
	token, err := servTool.GenerateAccessKey(username, clientIP)
	if err != nil {
		info := fmt.Sprintf("生成登录秘钥失败: %s", err)
		logger.Error(info)
		return false, info
	}
	// ExpireTime 单位被识别成微秒，导致一注册值后立马过期
	//mgr.Cache.Set(username, authString, ExpireTime)
	mgr.Cache.Set(username, token, time.Minute*ExpireTime)

	return true, token
}

func (mgr *AccountMgr) Logout(ctx *gin.Context) bool {
	// 检查登录状态
	if !mgr.IfLoginIn() {
		logger.Error("请先登录")
		return false
	}

	// TODO: 验证访问秘钥（需要在鉴权中间件中实现）
	token := ctx.Request.Header.Get("Authorization")
	if result, err := mgr.Cache.Get(mgr.Account.Username).Result(); err != nil || token != result {
		// 缓存获取token失败或token匹配失败
		return false
	}
	// 校验token中的IP与当前访问的实际IP
	servTool := tools.ServTool{}
	info, _ := servTool.DecryptAccessKey(token)
	remoteIP := servTool.FetchRemoteIp(ctx)
	infoDetail := strings.Split(info, "@")
	if infoDetail[0] != mgr.Account.Username || infoDetail[1] != remoteIP {
		// token校验失败
		return false
	}

	// 注销登录
	if err := mgr.Cache.Del(mgr.Account.Username).Err(); err != nil {
		logger.Errorf("注销登录失败")
		return false
	}
	return true
}

func (mgr *AccountMgr) Kick(username []string) string {
	if username == nil {
		return ""
	}

	var strs []string
	for _, str := range username {
		if err := mgr.Cache.Del(str).Err(); err != nil {
			logger.Errorf("强制[%s]下线失败", str)
			strs = append(strs, str)
		}
	}

	if len(strs) > 0 {
		return fmt.Sprintf("用户[%s]下线失败", strings.Join(strs, ","))
	}
	return "强制下线成功"
}

func (mgr *AccountMgr) Edit() error {
	return nil
}

func (mgr *AccountMgr) Ban(username []string) (bool, string) {
	if username == nil {
		return false, ""
	}
	if mgr.DBClient.Table(TableAccount).Where("username IN ?", username).Update("status", BANED).Error != nil {
		return false, fmt.Sprintf("封禁账户[%s]失败", strings.Join(username, ","))
	}
	return true, "封禁成功"
}

func (mgr *AccountMgr) Unban(username []string) (bool, string) {
	if username == nil {
		return false, ""
	}
	if mgr.DBClient.Table(TableAccount).Where("username IN ?", username).Update("status", ACTIVE).Error != nil {
		return false, fmt.Sprintf("解禁账户[%s]失败", strings.Join(username, ","))
	}
	return true, "解除封禁成功"
}

func (mgr *AccountMgr) Revoke() (bool, string) {
	if !mgr.Exists() {
		return false, fmt.Sprintf("账户[%s]不存在,注销失败", mgr.Account.Username)
	}
	if mgr.DBClient.Table(TableAccount).Where("username = ?", mgr.Account.Username).Update("status", REVOKE).Error != nil {
		return false, "注销失败,未知错误"
	}
	return true, "注销成功"
}
