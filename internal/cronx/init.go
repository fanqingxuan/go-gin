package cronx

import (
	"context"
	"go-gin/internal/components/logx"
	"go-gin/internal/traceid"
	"reflect"
	"runtime"
	"strings"
	"time"

	"github.com/robfig/cron/v3"
)

var c *cron.Cron

func New() {
	c = cron.New()
}

func Run() {
	c.Start()
	// 程序结束时停止 cron 定时器（可选）
	defer c.Stop()
	// 主程序可以保持运行状态，等待 cron 任务执行
	select {}
}

// Schedule 创建定时任务
func Schedule(job Job) *JobBuilder {
	if c == nil {
		panic("please call cronx.New() first")
	}
	return NewJobBuilder(job)
}

func AddJob(spec string, cmd Job) {
	jobName := getStructName(cmd)
	_, err := c.AddFunc(spec, func() {
		ctx := context.WithValue(context.Background(), traceid.TraceIdFieldName, traceid.New())
		logx.CronLoggerInstance.Info().Ctx(ctx).Str("cron", jobName).Str("spec", spec).Str("keywords", "开始执行").Send()

		start := time.Now()

		err := cmd.Handle(ctx)

		TimeStamp := time.Now()
		Cost := TimeStamp.Sub(start)
		if Cost > time.Minute {
			Cost = Cost.Truncate(time.Second)
		}
		if err != nil {
			logx.CronLoggerInstance.Error().Ctx(ctx).Str("cron", jobName).Str("spec", spec).Str("keywords", "执行结束").Str("cost", Cost.String()).Str("err", err.Error()).Send()
		} else {
			logx.CronLoggerInstance.Info().Ctx(ctx).Str("cron", jobName).Str("spec", spec).Str("keywords", "执行结束").Str("cost", Cost.String()).Send()
		}
	})
	if err != nil {
		logx.CronLoggerInstance.Info().Ctx(context.Background()).Str("cron", jobName).Str("spec", spec).Str("keywords", "添加失败").Send()
	}
}

// ScheduleFunc 创建函数类型的定时任务
func ScheduleFunc(fn JobFunc) *JobBuilder {
	if c == nil {
		panic("please call cronx.New() first")
	}
	return NewJobBuilder(fn)
}

func AddFunc(spec string, cmd JobFunc) {
	AddJob(spec, cmd)
}

// 添加一个工具函数来获取结构体名称
func getStructName(v interface{}) string {
	t := reflect.TypeOf(v)

	// 处理函数类型
	if t.Kind() == reflect.Func {
		// 获取函数的完整路径名
		fullName := runtime.FuncForPC(reflect.ValueOf(v).Pointer()).Name()
		// 提取最后一个点号后的函数名
		if lastDot := strings.LastIndex(fullName, "."); lastDot >= 0 {
			return fullName[lastDot+1:]
		}
		return fullName
	}

	// 处理指针类型
	if t.Kind() == reflect.Ptr {
		return t.Elem().Name()
	}

	// 处理普通类型
	return t.Name()
}
