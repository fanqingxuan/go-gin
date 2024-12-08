package routes

import (
	"go-gin/controllers"
	"go-gin/internal/httpx"
)

func RegisterUserRoutes(r *httpx.RouterGroup) {
	// r.Use(middlewares.TokenCheck())
	// 用户信息
	r.GET("/list", controllers.UserController.List)
	r.GET("/adduser", controllers.UserController.AddUser)

}
