package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sunyd/go-demo1/common"
	"github.com/sunyd/go-demo1/model"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 402, "msg": "权限不足"})
			ctx.Abort()
			return
		}

		tokenString = tokenString[7:]

		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 402, "msg": "权限不足"})
			ctx.Abort()
			return
		}

		// token合法，取出token的用户id信息，然后通过userid查询出用户信息
		userId := claims.UserId
		db := common.GetDB()
		var user model.User
		db.First(&user, userId)

		// 用户不存在
		if user.ID == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 402, "msg": "权限不足"})
			ctx.Abort()
			return
		}

		// 用户存在，将用户信息写入上下文
		ctx.Set("user", user)

		// 继续向下处理
		ctx.Next()

	}
}
