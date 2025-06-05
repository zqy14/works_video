package utils

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"video1/video-rpc/global"
)

var (
	db  *gorm.DB
	err error
)

func GlobalMysql() *gorm.DB {
	dsn := "root:root@tcp(127.0.0.1:3306)/video?charset=utf8mb4&parseTime=True&loc=Local"
	//dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.GlobalMysql.Mysql.User, config.GlobalMysql.Mysql.Password, config.GlobalMysql.Mysql.Host, config.GlobalMysql.Mysql.Port, config.GlobalMysql.Mysql.Database)
	global.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	log.Println("mysql init success")
	return db
}
