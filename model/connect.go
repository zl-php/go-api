package model

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DB 全局变量
var DB *gorm.DB

// 初始化mysql链接
func Mysql(connString string) {
	// 初始化GORM日志配置
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level(这里记得根据需求改一下)
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)

	db, err := gorm.Open(mysql.Open(connString), &gorm.Config{
		Logger: newLogger,
	})

	// Error
	if connString == "" || err != nil {
		panic(err)
	}

	// GORM 使用 database/sql 维护连接池 https://pkg.go.dev/database/sql
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	//设置连接池
	//SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(20)
	//SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	//超时
	sqlDB.SetConnMaxLifetime(time.Second * 30)

	DB = db

	// 迁移数据库
	migration()
}
