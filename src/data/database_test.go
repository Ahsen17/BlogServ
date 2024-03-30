package data

import (
	"BlogServ/config"
	"BlogServ/models"
	"testing"
)

var c *config.Config

//func TestInitDB(t *testing.T) {
//	c = config.FetchConfig()
//	//fmt.Println(c)
//	db := InitDB(c)
//	//db.Create(&models.User{Username: "test", Password: "test"})
//	var user models.User
//	db.First(&user, 1)
//	fmt.Println(user.Username, user.Password)
//}

func TestInitDBPool(t *testing.T) {
	c = config.NewConfig()
	InitDBClient(&c.Database)
}

func TestMigrate(t *testing.T) {
	c = config.NewConfig()
	dbPool := InitDBClient(&c.Database)
	// 测试数据库表迁移
	Migrate(dbPool, &models.User{})
}
