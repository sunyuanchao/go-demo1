package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sunyd/go-demo1/controller"
	"github.com/sunyd/go-demo1/middleware"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	r.GET("/api/user", middleware.AuthMiddleware(), controller.Info)
	return r
}
