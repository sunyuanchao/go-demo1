package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type User struct {
	gorm.Model
	Username  string `gorm:"varchar(20);not null"`
	Password  string `gorm:"size:255;not null"`
	Telephone string `gorm:"varchar(200);not null;unique"`
}

func main() {

	db := initSqlDb()
	defer db.Close()

	r := gin.Default()
	r.POST("/api/auth/register", func(ctx *gin.Context) {
		//获取请求参数
		username := ctx.PostForm("username")
		password := ctx.PostForm("password")
		telephone := ctx.PostForm("telephone")
		//验证参数
		if len(telephone) != 11 {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号长度不合法,长度为11位"})
			return
		}
		//验证密码，最小长度为10位
		if len(password) < 6 {
			ctx.JSON(http.StatusUnprocessableEntity, map[string]interface{}{"code": 422, "msg": "用户密码长度不合法，长度至少为6位"})
			return
		}
		//如果用户名没有，默认生成长度为8的用户名
		if len(username) == 0 {
			username = randomString(8)
		}
		log.Println(username, password, telephone)

		//验证手机是否已经存在，已经存在不允许注册
		if telephoneIsExist(telephone, db) {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 402, "msg": "用户已经存在"})
			return
		}

		//验证通过，新增该用户
		newUser := User{
			Username:  username,
			Password:  password,
			Telephone: telephone,
		}

		db.Create(&newUser)

		//
		ctx.JSON(200, gin.H{
			"code": 200,
			"msg":  "注册成功",
		})
	})
	r.Run() // listen a
}

/**
查询手机号是否已经注册
*/
func telephoneIsExist(telephone string, db *gorm.DB) bool {

	var user User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

/**
随机字符串生成.
*/
func randomString(n int) string {
	var letters = []byte("asdfghjklqwertyuiopzxcvbnmASDFGHJKLQWERTYUIOPZXCVBNM")
	result := make([]byte, n)
	rand.Seed(time.Now().Unix())
	for i, _ := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}


func initSqlDb() *gorm.DB {
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
	db.AutoMigrate(&User{})
	return db

}
