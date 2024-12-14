package routes

import (
	"go-gin/controllers"
	"go-gin/events"
	"go-gin/internal/g"
	"go-gin/internal/httpx"
	"go-gin/internal/task"
	"go-gin/middlewares"
	"go-gin/models"
	"go-gin/tasks"
	"time"
)

func RegisterDemoRoutes(r *httpx.RouterGroup) {

	r.GET("/", middlewares.AfterSampleA(), controllers.UserController.Index)

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
}
