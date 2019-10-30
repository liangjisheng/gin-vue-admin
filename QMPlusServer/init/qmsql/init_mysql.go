package qmsql

import (
	"gin-vue-admin/QMPlusServer/config"
	"log"

	"github.com/jinzhu/gorm"

	// mysql driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DEFAULTDB ...
var DEFAULTDB *gorm.DB

// InitMysql 初始化数据库并产生数据库全局变量
func InitMysql(admin config.Admin) *gorm.DB {
	if db, err := gorm.Open("mysql", admin.Username+":"+admin.Password+"@("+admin.Path+")/"+admin.DBName+"?"+admin.Config); err != nil {
		log.Printf("DEFAULTDB数据库启动异常%+v", err)
	} else {
		DEFAULTDB = db
		DEFAULTDB.DB().SetMaxIdleConns(10)
		DEFAULTDB.DB().SetMaxIdleConns(100)
	}
	return DEFAULTDB
}
