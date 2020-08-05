package common

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"github.com/sunyd/go-demo1/model"
)

var DB *gorm.DB

/**
数据库初始化.
*/
func InitSqlDb() *gorm.DB {
	driveName := viper.GetString("datasource.driverName")
	host := viper.GetString("datasource.host")
	port := viper.GetInt("datasource.port")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	charset := viper.GetString("datasource.charset")

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
