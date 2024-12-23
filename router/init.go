package router

import (
	"go-gin/controller"
	"go-gin/internal/errorx"
	"go-gin/internal/httpx"
)

func Init(route *httpx.Engine) {
	route.NoMethod(func(ctx *httpx.Context) (any, error) {
		return nil, errorx.ErrMethodNotAllowed
	})
	route.NoRoute(func(ctx *httpx.Context) (any, error) {
		return nil, errorx.ErrNoRoute
	})
	route.GET("/login", controller.LoginController.Login)
	route.Group("/").GET("/", controller.UserController.Index)
	RegisterUserRoutes(route.Group("/user"))
	RegisterLoginRoutes(route.Group("/"))
	RegisterApiRoutes(route.Group("/api"))
	RegisterDemoRoutes(route.Group("/demo"))
}
