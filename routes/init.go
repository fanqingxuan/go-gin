package routes

import (
	"go-gin/controllers"
	"go-gin/internal/errorx"
	"go-gin/internal/httpx"
)

func Init(route *httpx.Engine) {
	route.NoMethod(func(ctx *httpx.Context) (interface{}, error) {
		return nil, errorx.ErrMethodNotAllowed
	})
	route.NoRoute(func(ctx *httpx.Context) (interface{}, error) {
		return nil, errorx.ErrNoRoute
	})
	route.GET("/login", controllers.LoginController.Login)
	route.Group("/").GET("/", controllers.UserController.Index)
	RegisterUserRoutes(route.Group("/user"))
	RegisterLoginRoutes(route.Group("/"))
	RegisterApiRoutes(route.Group("/api"))
	RegisterDemoRoutes(route.Group("/demo"))
}
