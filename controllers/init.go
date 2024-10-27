package controllers

import (
	"go-gin/events"
	"go-gin/internal/errorx"
	"go-gin/internal/ginx/httpx"
	"go-gin/models"
	"go-gin/tasks"
	"time"

	"github.com/gin-gonic/gin"
)

func Init(route *gin.Engine) {

	route.NoMethod(func(ctx *gin.Context) {
		httpx.Error(ctx, errorx.ErrMethodNotAllowed)
	})
	route.NoRoute(func(ctx *gin.Context) {
		httpx.Error(ctx, errorx.ErrNoRoute)
	})
	notNeedAuthRouteList(route)
	needAuthRouteList(route)
}

// 需要登录的路由
func needAuthRouteList(route *gin.Engine) {
	r := route.Group("")
	// r.Use(middlewares.TokenCheck())
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
	r.GET("/task", func(ctx *gin.Context) {
		// err := task.DispatchNow(tasks.NewSampleTask("测试1234"))
		// fmt.Println(err)
		// err = task.Dispatch(tasks.NewSampleBTask("测试1234"), time.Second)

		// fmt.Println(err)
		tasks.NewSampleTask("测试3333").DispatchNow()
		tasks.NewSampleTask("测试3333").Dispatch(5 * time.Second)
		ctx.String(200, "hello world")
	})
	r.GET("/event", func(ctx *gin.Context) {
		// event.Fire(ctx, events.NewSampleEvent("hello 测试"))
		events.NewSampleEvent("333").Fire(ctx)
		events.NewDemoEvent(&models.User{Name: "hello"}).Fire(ctx)
		ctx.String(200, "hello world")
	})

	// api测试
	api_router := r.Group("/api")
	api_router.GET("/", ApiController.Index)
	api_router.GET("/indexa", ApiController.IndexA)
	api_router.GET("/loginapi", ApiController.IndexB)
	api_router.GET("/mylogin", ApiController.IndexC)

	api_router.Any("/list", ApiController.List)
}
