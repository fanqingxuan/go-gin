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
	notNeedAuthRouteList(route)
	needAuthRouteList(route)
}

// 需要登录的路由
func needAuthRouteList(route *gin.Engine) {
	r := route.Group("")
	r.Use(middlewares.TokenCheck())
	// 用户信息
	user_router := r.Group("/user")
	user_router.GET("/list", UserController.List)
	user_router.GET("/adduser", UserController.AddUser)

	// 退出登录
	login_router := r.Group("/")
	login_router.GET("/logout", LoginController.LoginOut)

}

// 不需要登录的路由
func notNeedAuthRouteList(route *gin.Engine) {
	route.GET("/login", LoginController.Login)

	r := route.Group("/")
	r.GET("/", UserController.Index)

	// 登录注册

	// api测试
	api_router := r.Group("/api")
	api_router.GET("/", ApiController.Index)
	api_router.GET("/indexa", ApiController.IndexA)
	api_router.GET("/indexb", ApiController.IndexB)
	api_router.POST("/list", ApiController.List)
}
