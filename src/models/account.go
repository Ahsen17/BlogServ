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
	"github.com/go-redis/redis"
	"gorm.io/gorm"
	"strings"
)

const (
	TABLE = "account"

	ACTIVE     = 1
	DEACTIVATE = 2
	REVOKE     = 3
	BANED      = 9

	DefaultAccount = SYSTEM

	ExpireTime = 12 * 3600
)

type Account struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
	Status   int    `json:"status"`
	CreateBy string `json:"create_by"`
	UpdateBy string `json:"update_by"`
}

type AccountData struct {
	LastLoginTime int64
	LastLoginIp   string
}

type AccountMgr struct {
	Account     *Account
	AccountData *AccountData

	UserMgr      *UserMgr
	RoleMgr      *RoleMgr
	CustomLogMgr *CustomLogMgr // 用户操作日志

	DBClient *gorm.DB
	Cache    *redis.Client
}

func (a Account) TableName() string {
	return TABLE
}

func (mgr *AccountMgr) Exists() bool {
	return mgr.DBClient.Where("username = ?", mgr.Account.Username).First(&mgr.Account).RowsAffected > 0
}

func (mgr *AccountMgr) CheckPassword() bool {
	return mgr.DBClient.Where(
		"username = ?, password = ?",
		mgr.Account.Username,
		mgr.Account.Password,
	).First(&mgr.Account).RowsAffected > 0
}

func (mgr *AccountMgr) IfLoginIn() bool {
	exists, err := mgr.Cache.Exists(mgr.Account.Username).Result()
	if err != nil || exists != 0 {
		return true
	}
	return false
}

func (mgr *AccountMgr) Register() bool {
	if mgr.Exists() {
		logger.Errorf("用户[%s]已存在", mgr.Account.Username)
		return false
	}
	mgr.Account.Status = DEACTIVATE // 刚注册默认未激活
	if err := mgr.DBClient.Create(&mgr.Account).Error; err != nil {
		logger.Errorf("注册用户[%s]失败: %s", mgr.Account.Username, err)
		return false
	}
	return true
}

func (mgr *AccountMgr) Login() (bool, string) {
	username := mgr.Account.Username

	// 判断账户是否存在/密码是否正确
	if !mgr.Exists() || !mgr.CheckPassword() {
		info := fmt.Sprintf("账户[%s]不存在或密码错误", username)
		return false, info
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

	// TODO: 生成访问密钥
	var authString string
	mgr.Cache.Set(username, authString, ExpireTime)

	return true, authString
}

func (mgr *AccountMgr) Logout() bool {
	// 检查登录状态
	if !mgr.IfLoginIn() {
		logger.Error("请先登录")
		return false
	}

	// 注销登录
	err := mgr.Cache.Del(mgr.Account.Username).Err()
	if err != nil {
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
	if mgr.DBClient.Table(TABLE).Where("username IN ?", username).Update("status", BANED).Error != nil {
		return false, fmt.Sprintf("封禁账户[%s]失败", strings.Join(username, ","))
	}
	return true, "封禁成功"
}

func (mgr *AccountMgr) Unban(username []string) (bool, string) {
	if username == nil {
		return false, ""
	}
	if mgr.DBClient.Table(TABLE).Where("username IN ?", username).Update("status", ACTIVE).Error != nil {
		return false, fmt.Sprintf("解禁账户[%s]失败", strings.Join(username, ","))
	}
	return true, "解除封禁成功"
}

func (mgr *AccountMgr) Revoke() (bool, string) {
	if !mgr.Exists() {
		return false, fmt.Sprintf("账户[%s]不存在,注销失败", mgr.Account.Username)
	}
	if mgr.DBClient.Where("username = ?", mgr.Account.Username).Update("status", REVOKE).Error != nil {
		return false, "注销失败,未知错误"
	}
	return true, "注销成功"
}
