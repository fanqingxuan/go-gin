package router

import (
	"fmt"
	"go-gin/controller"
	"go-gin/event"
	"go-gin/internal/eventbus"
	"go-gin/internal/excelx"
	"go-gin/internal/g"
	"go-gin/internal/httpx"
	"go-gin/internal/queue"
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
		// queue.NewOption().Queue(queue.HIGH).TaskID("test").Dispatch(task.NewSampleBTask("hello"))
		// queue.NewOption().Queue(queue.HIGH).TaskID("test").DispatchIf(true, task.NewSampleBTask("hello"))
		queue.NewOption().LowQueue().TaskID("test").DispatchIf(true, task.NewSampleBTask("hello"))
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
		return g.MapStrInt{"hello": 333}, nil
	})

	// 测试导出 Excel
	r.GET("/export/excel", func(ctx *httpx.Context) (any, error) {
		type User struct {
			ID   int
			Name string
			Age  int
		}
		users := []User{
			{ID: 1, Name: "张三", Age: 25},
			{ID: 2, Name: "李四", Age: 30},
			{ID: 3, Name: "王五", Age: 28},
		}

		headers := []string{"ID", "姓名", "年龄"}
		return nil, excelx.Download(ctx.Context, "用户列表.xlsx", headers, excelx.StructsToRows(users))
	})

	// 测试导出 CSV
	r.GET("/export/csv", func(ctx *httpx.Context) (any, error) {
		type User struct {
			ID   int
			Name string
			Age  int
		}
		users := []User{
			{ID: 1, Name: "张三", Age: 25},
			{ID: 2, Name: "李四", Age: 30},
			{ID: 3, Name: "王五", Age: 28},
		}

		headers := []string{"ID", "姓名", "年龄"}
		return nil, excelx.DownloadCSV(ctx.Context, "用户列表.csv", headers, excelx.StructsToStringRows(users))
	})
}
