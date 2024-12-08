package controllers

import (
	"go-gin/events"
	"go-gin/internal/errorx"
	"go-gin/internal/g"
	"go-gin/internal/httpx"
	"go-gin/internal/task"
	"go-gin/models"
	"go-gin/tasks"
	"time"
)

func Init(route *httpx.Engine) {

	route.NoMethod(func(ctx *httpx.Context) (interface{}, error) {
		return nil, errorx.ErrMethodNotAllowed
	})
	route.NoRoute(func(ctx *httpx.Context) (interface{}, error) {
		return nil, errorx.ErrNoRoute
	})
	notNeedAuthRouteList(route)
	needAuthRouteList(route)
}

// 需要登录的路由
func needAuthRouteList(route *httpx.Engine) {
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
func notNeedAuthRouteList(route *httpx.Engine) {
	route.GET("/login", LoginController.Login)

	r := route.Group("/")
	r.GET("/", UserController.Index)
	r.GET("/task", func(ctx *httpx.Context) (interface{}, error) {
		// err := task.DispatchNow(tasks.NewSampleTask("测试1234"))
		// fmt.Println(err)
		// err = task.Dispatch(tasks.NewSampleBTask("测试1234"), time.Second)

		// fmt.Println(err)
		tasks.NewSampleTask("测试1").DispatchNow()
		tasks.NewSampleTask("测试2").Dispatch(5 * time.Second)
		task.NewOption().Queue(task.HIGH).TaskID("test").Dispatch(tasks.NewSampleBTask("hello"))
		return "hello world", nil
	})
	r.GET("/event", func(ctx *httpx.Context) (interface{}, error) {
		// event.Fire(ctx, events.NewSampleEvent("hello 测试"))
		events.NewSampleEvent("333").Fire(ctx)
		events.NewDemoEvent(&models.User{Name: "hello"}).Fire(ctx)
		return "hello world", nil
	})
	r.GET("/test", func(ctx *httpx.Context) (interface{}, error) {
		return g.MapStrInt{"hello": 333}, nil
	})

	// api测试
	api_router := r.Group("/api")
	api_router.GET("/", ApiController.Index)
	api_router.GET("/indexa", ApiController.IndexA)
	api_router.GET("/loginapi", ApiController.IndexB)
	api_router.GET("/mylogin", ApiController.IndexC)

	api_router.Any("/list", ApiController.List)
}
