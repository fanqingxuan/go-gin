package router

import (
	"go-gin/controller"
	"go-gin/internal/httpx"
)

func RegisterLoginRoutes(r *httpx.RouterGroup) {
	r.GET("/login", controller.LoginController.Login)

	// 退出登录
	r.GET("/logout", controller.LoginController.LoginOut)
}
