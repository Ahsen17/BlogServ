/*
  Package models
  @Author: Ahsen17
  @Github: https://github.com/Ahsen17
  @Time: 2024/5/1 上午5:28
  @Description: 文档管理
*/

package models

import "gorm.io/gorm"

const (
	TableDocument = "document"
)

type Document struct {
	gorm.Model
}

type DocumentData struct {
}

type DocumentMgr struct {
	Document *Document

	DBClient *gorm.DB
}

func (d Document) TableName() string {
	return TableDocument
}
