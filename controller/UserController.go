package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/sunyd/go-demo1/common"
	"github.com/sunyd/go-demo1/model"
	"github.com/sunyd/go-demo1/util"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

/**
用户注册.
*/
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

	//加密用户密码
	hasedPasswd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "密码加密失败"})
		return
	}

	//验证通过，新增该用户
	newUser := model.User{
		Username:  username,
		Password:  string(hasedPasswd),
		Telephone: telephone,
	}

	db.Create(&newUser)

	//
	ctx.JSON(200, gin.H{
		"code": 200,
		"msg":  "注册成功",
	})
}

func Login(ctx *gin.Context) {

	db := common.GetDB()

	//获取参数
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")

	//验证参数
	if len(telephone) != 11 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 402, "msg": "手机号不足11位"})
		return
	}
	if len(password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 402, "msg": "密码长度不足6位"})
		return
	}

	//数据库查询出用户
	var user model.User
	db.Where("telephone =?", telephone).First(&user)
	if user.ID == 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户不存在"})
		return
	}

	//判断用户密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码不正确"})
		return
	}

	//用户登录成功，返回token
	token := "111"

	//返回结果
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": "登录成功", "data": gin.H{"token": token}})

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
