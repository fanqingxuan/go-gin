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
		// task.NewSampleTask("测试1").DispatchNow()
		task.NewSampleTask("测试1").DispatchIf(true)
		task.NewSampleTask("测试2").Dispatch(5 * time.Second)
		// taskx.NewOption().Queue(taskx.HIGH).TaskID("test").Dispatch(task.NewSampleBTask("hello"))
		// taskx.NewOption().Queue(taskx.HIGH).TaskID("test").DispatchIf(true, task.NewSampleBTask("hello"))
		taskx.NewOption().LowQueue().TaskID("test").DispatchIf(true, task.NewSampleBTask("hello"))
		return "hello world", nil
	})
	r.GET("/event", func(ctx *httpx.Context) (any, error) {
		eventbus.Fire(ctx, event.NewSampleEvent("hello 测试"))
		event.NewSampleEvent("333").FireIf(ctx, true)
		// event.NewDemoEvent(&model.User{Name: "hello"}).Fire(ctx)
		return "hello world", nil
	})
	r.GET("/test", func(ctx *httpx.Context) (any, error) {

		a := 43
		fmt.Println(util.IsTrue(a))
		util.Unless(a > 40, func() {
			fmt.Println("hello", a)
		})
		return g.MapStrInt{"hello": 333}, nil
	})
}
