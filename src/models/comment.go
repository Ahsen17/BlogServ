/*
  Package models
  @Author: Ahsen17
  @Github: https://github.com/Ahsen17
  @Time: 2024/5/1 上午5:29
  @Description: 评论管理
*/

package models

import "gorm.io/gorm"

const (
	TableComment = "comment"
)

type Comment struct {
	gorm.Model
}

type CommentMgr struct {
	Comment *Comment

	DBClient *gorm.DB
}

func (c Comment) TableName() string {
	return TableComment
}
