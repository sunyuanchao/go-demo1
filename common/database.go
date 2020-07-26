package common

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/sunyd/go-demo1/model"
)

var DB *gorm.DB

/**
数据库初始化.
*/
func InitSqlDb() *gorm.DB {
	driveName := "mysql"
	host := "localhost"
	port := 3306
	database := "test"
	username := "root"
	password := "root"
	charset := "utf8"

	args := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)
	fmt.Println(args)
	db, err := gorm.Open(driveName, args)
	if err != nil {
		panic("fail to connect database,err:= " + err.Error())
	}

	//自动创建数据表
	db.AutoMigrate(&model.User{})
	DB = db
	return db

}

func GetDB() *gorm.DB {
	return DB
}
