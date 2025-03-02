package router

import (
	"go-gin/internal/httpx"
)

func Init(route *httpx.Engine) {
	RegisterCommonRoutes(route)
	RegisterUserRoutes(route.Group("/user"))
	RegisterLoginRoutes(route.Group("/"))
	RegisterApiRoutes(route.Group("/api"))
	RegisterDemoRoutes(route.Group("/demo"))
}
