package database

import (
	"fmt"

	"gin_init/pkg/logger"
	"gin_init/pkg/setting"
	"os"

	"github.com/jinzhu/gorm"
	// mysql drive
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

// Setup 初始化
func Setup() {
	var err error

	username := setting.DatabaseSetting.User
	password := setting.DatabaseSetting.Password
	host := setting.DatabaseSetting.Host

	source := fmt.Sprintf("%s:%s@%s", username, password, host)

	db, err = gorm.Open("mysql", source)
	if err != nil {
		logger.Error("db error", err.Error())
	}

	if os.Getenv("GINENV") != "production" {
		db.LogMode(true)
	}

	// 全局禁用表名复数
	db.SingularTable(true) // 如果设置为true,`User`的默认表名为`user`,使用`TableName`设置的表名不受影响

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(200)
}

// CloseDB 关闭
func CloseDB() {
	defer db.Close()
}

//Conn MysqlConn 返回mysql连
func Conn() *gorm.DB {
	return db
}
