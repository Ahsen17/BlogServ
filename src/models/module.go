/*
  Package models
  @Author: Ahsen17
  @Github: https://github.com/Ahsen17
  @Time:
  @Description: 功能模块管理
*/

package models

import (
	"github.com/ahsen17/BlogServ/logger"
	"gorm.io/gorm"
)

type Module struct {
	gorm.Model
	Name  string `json:"name" gorm:"column:name"`
	URL   string `json:"url" gorm:"column:url type:text not null"`
	Level int    `json:"level" gorm:"column:level type:int not null"`
}

type ModuleMgr struct {
	Module *Module

	DBClient *gorm.DB
}

func (m *ModuleMgr) Exists() bool {
	return m.DBClient.Where("url = ?", m.Module.URL).First(&m.Module).RowsAffected > 0
}

func (m *ModuleMgr) Access(role Role) bool {
	if !m.Exists() {
		logger.Errorf("模块[%s]不存在", m.Module.Name)
		return false
	}
	// 判断是否有访问权限
	return m.Module.Level > role.Level
}

func (m *ModuleMgr) Register() bool {
	if m.Exists() {
		logger.Errorf("模块[%s]已存在", m.Module.Name)
		return false
	}
	if err := m.DBClient.Create(&m.Module).Error; err != nil {
		logger.Errorf("注册[%s]模块失败: %s", m.Module.Name, err)
		return false
	}
	logger.Infof("注册[%s]模块成功", m.Module.Name)
	return true
}
