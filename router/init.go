package router

import (
	"go-gin/controller"
	"go-gin/internal/components/db"
	"go-gin/internal/components/redisx"
	"go-gin/internal/errorx"
	"go-gin/internal/g"
	"go-gin/internal/httpx"
	"go-gin/util"
)

func Init(route *httpx.Engine) {
	route.NoMethod(func(ctx *httpx.Context) (any, error) {
		return nil, errorx.ErrMethodNotAllowed
	})
	route.NoRoute(func(ctx *httpx.Context) (any, error) {
		return nil, errorx.ErrNoRoute
	})
	// 健康检测
	route.GET("/status", func(ctx *httpx.Context) (any, error) {
		db_err := db.WithContext(ctx).Ping()
		redis_err := redisx.GetInstance().Ping(ctx).Err()
		return g.MapStrStr{
			"database": util.Conditional(db_err == nil, "ok", "failed"),
			"redis":    util.Conditional(redis_err == nil, "ok", "failed"),
		}, nil
	})
	route.GET("/login", controller.LoginController.Login)
	route.Group("/").GET("/", controller.UserController.Index)
	RegisterUserRoutes(route.Group("/user"))
	RegisterLoginRoutes(route.Group("/"))
	RegisterApiRoutes(route.Group("/api"))
	RegisterDemoRoutes(route.Group("/demo"))
}
