/*
  Package models
  @Author: Ahsen17
  @Github: https://github.com/Ahsen17
  @Time: 2024/5/1 上午5:34
  @Description: ....
*/

package models

import (
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

const (
	TableTag = "tag"
)

type Tag struct {
	gorm.Model
}

type TagMgr struct {
	Tag *Tag

	DBClient *gorm.DB
	Cache    *redis.Client
}

func (t Tag) TableName() string {
	return TableTag
}
