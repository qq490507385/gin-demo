package config

import (
	"doc-system/util"
	"fmt"

	"github.com/jinzhu/gorm"

	// log
	"github.com/lexkong/log"

	// Mysql driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DB  数据库
var db *gorm.DB

// InitDB 默认初始化
func InitDB() *gorm.DB {
	db = Connect()
	return db
}

// Connect 连接默认数据库
func Connect() *gorm.DB {
	return ConnectSQL(c.connectDefault)
}

// ConnectSQL 连接指定数据库
func ConnectSQL(dbKey string) *gorm.DB {
	s := c.connections[dbKey]
	var (
		driver   = util.FlectVal(s, "driver")
		host     = util.FlectVal(s, "host")
		port     = util.FlectVal(s, "port")
		database = util.FlectVal(s, "database")
		username = util.FlectVal(s, "username")
		password = util.FlectVal(s, "password")
		charset  = util.FlectVal(s, "charset")
	)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=Local", username, password, host, port, database, charset)
	db, err := gorm.Open(driver.String(), dsn)
	if err != nil {
		log.Warnf("failed to connect database")
		fmt.Println("failed to connect database")
	}
	return db
}

// Close 关闭数据库
func Close(db *gorm.DB) {
	db.Close()
}
