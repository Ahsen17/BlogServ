/*
  Package models
  @Author: Ahsen17
  @Github: https://github.com/Ahsen17
  @Time: 2024/5/1 上午5:11
  @Description: 消息管理
*/

package models

import (
	"gorm.io/gorm"
)

const (
	TableNotify = "notify"

	AnonymousMsg = iota
	PlatformMsg
	PlatformNotify
	SystemNotify
)

type Notify struct {
	gorm.Model
	Content string `json:"content"`
	Type    int    `json:"type"`
	Level   int    `json:"level"`
	Dst     uint   `json:"destiny"`
	Sor     uint   `json:"source"`
}

type NotifyMgr struct {
	Notify *Notify

	DBClient *gorm.DB
}

func (n Notify) TableName() string {
	return TableNotify
}

func (mgr *NotifyMgr) Add2Queue() bool {
	// TODO: 添加消息到队列
	return false
}
