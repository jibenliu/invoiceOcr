package models

import (
	"github.com/jinzhu/gorm"
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DB 数据库链接单例
var DB *gorm.DB

// Database 在中间件中初始化mysql链接
func Database(connString string) {
	db, err := gorm.Open("mysql", connString)
	if err != nil {
		panic("数据库链接失败！")
	}
	db.LogMode(true)
	db.DB().SetMaxIdleConns(50)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(time.Second * 30)
	DB = db
}
