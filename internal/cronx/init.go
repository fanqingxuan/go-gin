package cronx

import (
	"context"
	"go-gin/internal/component/logx"
	"go-gin/internal/traceid"
	"os"
	"os/signal"
	"reflect"
	"runtime"
	"strings"
	"syscall"
	"time"

	"github.com/robfig/cron/v3"
)

var (
	c         *cron.Cron
	parentCtx context.Context
	cancel    context.CancelFunc
)

func New() {
	c = cron.New()
	parentCtx, cancel = context.WithCancel(context.Background())
}

// Run starts cron and blocks until receiving a termination signal (SIGINT/SIGTERM).
func Run() {
	c.Start()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	cancel()
	c.Stop()
}

// Stop cancels all running jobs and stops the cron scheduler.
func Stop() {
	cancel()
	c.Stop()
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
		ctx := context.WithValue(parentCtx, traceid.TraceIdFieldName, traceid.New())
		logx.CronLogger.Info().Ctx(ctx).Str("cron", jobName).Str("spec", spec).Str("keywords", "开始执行").Send()

		start := time.Now()

		err := cmd.Handle(ctx)

		timestamp := time.Now()
		cost := timestamp.Sub(start)
		if cost > time.Minute {
			cost = cost.Truncate(time.Second)
		}
		if err != nil {
			logx.CronLogger.Error().Ctx(ctx).Str("cron", jobName).Str("spec", spec).Str("keywords", "执行结束").Str("cost", cost.String()).Str("err", err.Error()).Send()
		} else {
			logx.CronLogger.Info().Ctx(ctx).Str("cron", jobName).Str("spec", spec).Str("keywords", "执行结束").Str("cost", cost.String()).Send()
		}
	})
	if err != nil {
		logx.CronLogger.Info().Ctx(context.Background()).Str("cron", jobName).Str("spec", spec).Str("keywords", "添加失败").Send()
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
	if v == nil {
		return "unknown"
	}
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
