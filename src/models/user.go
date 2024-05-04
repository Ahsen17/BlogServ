/*
  Package models
  @Author: Ahsen17
  @Github: https://github.com/Ahsen17
  @Time:
  @Description: 用户管理
*/

package models

import (
	"gorm.io/gorm"
	"time"
)

const (
	TableUser = "user"
)

type User struct {
	gorm.Model
	Info UserInfo `json:"info"`
	Data UserData `json:"datasource"`
}

type UserInfo struct {
	Avatar   string     `json:"avatar"`
	Nickname string     `json:"nickname"`
	Sex      uint       `json:"sex"`
	Birth    *time.Time `json:"birth"`
	Email    string     `json:"email"`
	Address  string     `json:"address"`
	Site     string     `json:"site"`
}

type UserData struct {
}

type UserMgr struct {
	User     *User
	UserData *UserData

	DBClient *gorm.DB
}

func (user User) TableName() string {
	return TableUser
}

// Exist 检查用户是否存在
func (m *UserMgr) Exist() bool {
	return m.DBClient.Table(TableAccount).Where(
		"username = ?", m.User.Info.Nickname,
	).First(&m.User).RowsAffected > 0
}

// Create 创建用户
func (m *UserMgr) Create() bool {
	if m.Exist() {
		return false
	}
	if err := m.DBClient.Table(TableAccount).Create(&m.User).Error; err != nil {
		return false
	}
	return true
}

// Edit 编辑用户
func (m *UserMgr) Edit() bool {
	if !m.Exist() {
		return false
	}
	if err := m.DBClient.Table(TableAccount).Updates(&m.User).Error; err != nil {
		return false
	}
	return true
}

// Delete 删除用户
func (m *UserMgr) Delete() error {
	return m.DBClient.Table(TableAccount).Delete(&m.User).Error
}

// Single 获取用户
// RetrieveSingle 获取单个用户
func (m *UserMgr) Single() bool {
	if !m.Exist() {
		return false
	}
	if err := m.DBClient.Table(TableAccount).First(&m.User).Error; err != nil {
		return false
	}
	return true
}

// Batch 批量获取用户
// RetrieveBatch 获取批量用户，通过nickname关键字搜索
func (m *UserMgr) Batch(keyword string) []*User {
	var users []*User
	if err := m.DBClient.Table(TableAccount).Where("username = ?", keyword).Find(&users).Error; err != nil {
		return nil
	}
	return users
}

// Following 获取用户收藏
func (m *UserMgr) Following() []*User {
	// TODO: 获取当前用户的所有关注用户
	return nil
}

// Follower 获取用户粉丝
func (m *UserMgr) Follower() []*User {
	// TODO: 获取当前用户的所有粉丝
	return nil
}
