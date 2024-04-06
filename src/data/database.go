package data

import (
	"BlogServ/config"
	"BlogServ/logger"
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
)

var (
	client *gorm.DB
	lock   sync.Mutex
)

const (
	MYSQL  = "mysql"
	SQLITE = "sqlite"
)

// InitDBClient 初始化数据库连接池
func init() {
	dbConf := config.DBConfig()

	driverName := dbConf.Driver
	dsn := "%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local"
	dataSourceName := fmt.Sprintf(dsn, dbConf.Username, dbConf.Password,
		dbConf.Host, dbConf.Port, dbConf.Database)

	// 切换数据库驱动
	var conn gorm.Dialector
	switch driverName {
	case MYSQL:
		conn = mysql.Open(dataSourceName)
	}

	if conn != nil {
		logger.Fatal("数据库连接校验失败")
	}

	var connPool *sql.DB
	if client, err := gorm.Open(conn, &gorm.Config{}); err != nil {
		logger.Errorf("连接数据库失败: %s", err)
		panic(err)
	} else {
		connPool, err = client.DB()
		if err != nil {
			logger.Errorf("获取数据库连接池失败: %s", err)
			panic(err)
		}
	}

	// 设置最大连接数与最大空闲数
	connPool.SetMaxOpenConns(dbConf.MaxConn)
	connPool.SetMaxIdleConns(dbConf.IdleConn)

	logger.Info("数据库连接池初始化完成")
}

// Migrate 迁移数据库表
func Migrate(objs ...interface{}) error {
	// 迭代迁移数据库表
	if err := client.AutoMigrate(objs); err != nil {
		logger.Errorf("数据库表迁移失败")
		return err
	}

	logger.Info("数据库表迁移完成")
	return nil
}

func DBClient() *gorm.DB {
	return client
}
