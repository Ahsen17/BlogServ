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
	Data UserData `json:"data"`
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
