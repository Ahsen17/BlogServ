package data

import (
	"BlogServ/config"
	"BlogServ/logger"
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//database 选择mysql

// InitDB 初始化数据库连接
//func InitDB(c *config.Config) *gorm.DB {
//	dbConf := c.Database
//	dsnRaw := "%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local"
//	// 若数据库不存在则创建()
//	//dsn := fmt.Sprintf(dsnRaw, dbConf.Username, dbConf.Password, dbConf.Host, dbConf.Port, "")
//	//
//	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
//	//if err != nil {
//	//	logger.Fatal("连接数据库失败")
//	//}
//	//
//	//db.Exec("CREATE DATABASE IF NOT EXISTS " + dbConf.Database)
//	dsn := fmt.Sprintf(dsnRaw, dbConf.Username, dbConf.Password, dbConf.Host, dbConf.Port, dbConf.Database)
//	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
//	if err != nil {
//		logger.Fatal("连接数据库失败")
//	}
//
//	err = db.AutoMigrate(&models.User{})
//	if err != nil {
//		logger.Fatal("数据库表迁移失败")
//	}
//
//	return db
//}

var (
	dbPool *sql.DB
)

// InitDBPool 初始化数据库连接池
func initDBPool(db *config.Database) *sql.DB {
	driverName := db.Driver
	dsn := "%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local"
	dataSourceName := fmt.Sprintf(dsn, db.Username, db.Password, db.Host, db.Port, db.Database)

	var err error
	if dbPool, err = sql.Open(driverName, dataSourceName); err != nil {
		logger.Errorf("创建数据库连接池失败: %s", err)
		panic(err)
	}

	// 设置最大连接数与最大空闲数
	dbPool.SetMaxOpenConns(db.MaxConn)
	dbPool.SetMaxIdleConns(db.IdleConn)

	//defer func(dbPool *sql.DB) {
	//	err := dbPool.Close()
	//	if err != nil {
	//		logger.Fatalf("关闭数据库连接池失败: %s", err)
	//	}
	//}(dbPool)

	logger.Info("数据库连接池初始化完成")
	return dbPool
}

// Migrate 迁移数据库表
func Migrate(sqlDB *sql.DB, objs ...any) error {
	gormDB, err := gorm.Open(mysql.New(mysql.Config{Conn: sqlDB}), &gorm.Config{})
	if err != nil {
		logger.Errorf("从连接池获取数据库Conn失败: %s", err)
		return err
	}

	// 迭代迁移数据库表
	for _, obj := range objs {
		err := gormDB.AutoMigrate(&obj)
		if err != nil {
			logger.Errorf("迁移数据库表失败: %s", err)
			return err
		}
	}

	logger.Info("数据库表迁移完成")
	return nil
}
