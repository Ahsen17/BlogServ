package data

import (
	"BlogServ/config"
	"BlogServ/logger"
	"BlogServ/models"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//database 选择mysql

// InitDB 初始化数据库连接
func InitDB(c *config.Config) *gorm.DB {
	dbConf := c.Database
	dsnRaw := "%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local"
	dsn := fmt.Sprintf(dsnRaw, dbConf.Username, dbConf.Password, dbConf.Host, dbConf.Port, "")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Fatal("连接数据库失败")
	}

	db.Exec("CREATE DATABASE IF NOT EXISTS " + dbConf.Database)
	dsn = fmt.Sprintf(dsnRaw, dbConf.Username, dbConf.Password, dbConf.Host, dbConf.Port, dbConf.Database)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Fatal("连接数据库失败")
	}

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		logger.Fatal("数据库表迁移失败")
	}

	return db
}

// InitDBPool 初始化数据库连接池
func InitDBPool(c *config.Config) {
	return
}
