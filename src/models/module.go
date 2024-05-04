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

const (
	TableModule = "module"
)

type Module struct {
	gorm.Model
	Name   string `json:"name" gorm:"column:name"`
	Root   bool   `json:"root" gorm:"column:root"`     // 是否为根节点
	Parent uint   `json:"parent" gorm:"column:parent"` // 父节点
	URL    string `json:"url" gorm:"column:url type:text not null"`
	Level  int    `json:"level" gorm:"column:level type:int not null"` // 模块访问等级
}

type ModuleMgr struct {
	Module *Module

	DBClient *gorm.DB
}

func (m Module) TableName() string {
	return TableModule
}

// Exists 检查模块是否存在
func (m *ModuleMgr) Exists() bool {
	var count int64
	m.DBClient.Table(TableModule).Where("username = ?", m.Module.Name).Count(&count)
	return count > 0
}

// Access 检查访问权限
func (m *ModuleMgr) Access(role Role) bool {
	if !m.Exists() {
		logger.Errorf("模块[%s]不存在", m.Module.Name)
		return false
	}
	// 判断是否有访问权限
	return m.Module.Level > role.Level
}

// Register 注册模块
func (m *ModuleMgr) Register() bool {
	if m.Exists() {
		logger.Errorf("模块[%s]已存在", m.Module.Name)
		return false
	}
	if err := m.DBClient.Table(TableModule).Create(&m.Module).Error; err != nil {
		logger.Errorf("注册[%s]模块失败: %s", m.Module.Name, err)
		return false
	}
	logger.Infof("注册[%s]模块成功", m.Module.Name)
	return true
}

// Parent 获取父模块
func (m *ModuleMgr) Parent() *Module {
	if m.Module.Root == true {
		return nil
	}
	var module *Module
	if err := m.DBClient.Table(TableModule).Where("id = ?", m.Module.Parent).First(&module).Error; err != nil {
		logger.Errorf("获取父模块失败: %s", err)
		return nil
	}
	return module
}

// Children 获取子模块
func (m *ModuleMgr) Children() []*Module {
	var modules []*Module
	if err := m.DBClient.Table(TableModule).Where("parent = ?", m.Module.ID).Find(&modules).Error; err != nil {
		logger.Errorf("获取子模块失败: %s", err)
		return nil
	}
	return modules
}
