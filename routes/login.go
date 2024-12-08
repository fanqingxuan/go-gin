package routes

import (
	"go-gin/controllers"
	"go-gin/internal/httpx"
)

func RegisterLoginRoutes(r *httpx.RouterGroup) {
	// 退出登录
	r.GET("/logout", controllers.LoginController.LoginOut)
}
