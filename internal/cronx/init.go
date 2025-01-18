package cronx

import (
	"context"
	"fmt"
	"go-gin/internal/components/logx"
	"go-gin/internal/traceid"
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
	_, err := c.AddFunc(spec, func() {
		ctx := context.WithValue(context.Background(), traceid.TraceIdFieldName, traceid.New())
		logx.WithContext(ctx).Info("定时任务", fmt.Sprintf("开始执行,cron:%s", cmd.Name()))

		start := time.Now()

		err := cmd.Handle(ctx)

		TimeStamp := time.Now()
		Cost := TimeStamp.Sub(start)
		if Cost > time.Minute {
			Cost = Cost.Truncate(time.Second)
		}
		if err != nil {
			logx.WithContext(ctx).Error("定时任务", fmt.Sprintf("执行结束,cron:%s,spec:%s,cost:%s,error=%s", cmd.Name(), spec, Cost.String(), err.Error()))
		} else {
			logx.WithContext(ctx).Info("定时任务", fmt.Sprintf("执行结束,cron:%s,spec:%s,cost:%s", cmd.Name(), spec, Cost.String()))
		}
	})
	if err != nil {
		logx.WithContext(context.Background()).Error("定时任务", fmt.Sprintf("添加失败,cron:%s,spec:%s", cmd.Name(), spec))
	}
}
