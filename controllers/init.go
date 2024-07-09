package controllers

import (
	"go-gin/consts"
	"go-gin/internal/ginx/httpx"
	"go-gin/middlewares"

	"github.com/gin-gonic/gin"
)

func Init(route *gin.Engine) {

	route.NoMethod(func(ctx *gin.Context) {
		httpx.Error(ctx, consts.ErrMethodNotAllowed)
	})
	route.NoRoute(func(ctx *gin.Context) {
		httpx.Error(ctx, consts.ErrNoRoute)
	})

	r := route.Group("/")
	r.GET("/", UserController.Index)
	r.Use(middlewares.TokenCheck())

	// 用户信息
	user_router := r.Group("/user")
	user_router.GET("/list", UserController.List)
	user_router.GET("/adduser", UserController.AddUser)

	// 登录注册
	login_router := r.Group("/")
	route.GET("/login", LoginController.Login)
	login_router.GET("/logout", LoginController.LoginOut)

}
