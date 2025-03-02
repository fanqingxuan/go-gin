package router

import (
	"go-gin/controller"
	"go-gin/internal/httpx"
)

func RegisterUserRoutes(r *httpx.RouterGroup) {
	// r.Use(middlewares.TokenCheck())
	// 用户信息
	r.GET("/", controller.UserController.Index)
	r.GET("/list", controller.UserController.List)
	r.Match([]httpx.HttpMethod{httpx.MethodGet, httpx.MethodPost}, "/adduser", controller.UserController.AddUser)
	r.Any("/multiadduser", controller.UserController.MultiUserAdd)
}
