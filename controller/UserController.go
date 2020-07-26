package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/sunyd/go-demo1/common"
	"github.com/sunyd/go-demo1/model"
	"github.com/sunyd/go-demo1/util"
	"log"
	"net/http"
)

func Register(ctx *gin.Context) {
	db := common.GetDB()
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
		username = util.RandomString(8)
	}
	log.Println(username, password, telephone)

	//验证手机是否已经存在，已经存在不允许注册
	if telephoneIsExist(telephone, db) {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 402, "msg": "用户已经存在"})
		return
	}

	//验证通过，新增该用户
	newUser := model.User{
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
}

/**
查询手机号是否已经注册
*/
func telephoneIsExist(telephone string, db *gorm.DB) bool {

	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
