package router

import (
	"go-gin/controller"
	"go-gin/internal/httpx"
)

func RegisterLoginRoutes(r *httpx.RouterGroup) {
	// 退出登录
	r.GET("/logout", controller.LoginController.LoginOut)
}
