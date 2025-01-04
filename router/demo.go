package router

import (
	"fmt"
	"go-gin/controller"
	"go-gin/event"
	"go-gin/internal/eventbus"
	"go-gin/internal/g"
	"go-gin/internal/httpx"
	"go-gin/internal/taskx"
	"go-gin/middleware"
	"go-gin/model"
	"go-gin/task"
	"go-gin/util"
	"time"
)

func RegisterDemoRoutes(r *httpx.RouterGroup) {

	rr := r.Group("/")
	rr.After(middleware.AfterSampleA()).GET("/", controller.UserController.Index)

	r.GET("/task", func(ctx *httpx.Context) (any, error) {
		// err := task.DispatchNow(tasks.NewSampleTask("测试1234"))
		// fmt.Println(err)
		// err = task.Dispatch(tasks.NewSampleBTask("测试1234"), time.Second)

		// fmt.Println(err)
		task.NewSampleTask("测试1").DispatchNow()
		task.NewSampleTask("测试2").Dispatch(5 * time.Second)
		taskx.NewOption().Queue(taskx.HIGH).TaskID("test").Dispatch(task.NewSampleBTask("hello"))
		return "hello world", nil
	})
	r.GET("/event", func(ctx *httpx.Context) (any, error) {
		eventbus.Fire(ctx, event.NewSampleEvent("hello 测试"))
		event.NewSampleEvent("333").Fire(ctx)
		event.NewDemoEvent(&model.User{Name: "hello"}).Fire(ctx)
		return "hello world", nil
	})
	r.GET("/test", func(ctx *httpx.Context) (any, error) {

		a := 43
		fmt.Println(util.IfElse(a > 50, "hello", "default"))
		return g.MapStrInt{"hello": 333}, nil
	})
}
