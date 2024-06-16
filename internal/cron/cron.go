package cron

import (
	"context"
	"fmt"
	"go-gin/internal/components/logx"
	"go-gin/internal/traceid"
	"time"

	"github.com/robfig/cron"
)

type MYCron struct {
	cron cron.Cron
}

func New() *MYCron {
	return &MYCron{
		cron: *cron.New(),
	}
}

func (c *MYCron) Run() *MYCron {
	c.cron.Start()
	// 程序结束时停止 cron 定时器（可选）
	defer c.cron.Stop()
	// 主程序可以保持运行状态，等待 cron 任务执行
	select {}
}

func (c *MYCron) AddJob(spec string, cmd Job) {
	err := c.cron.AddFunc(spec, func() {
		ctx := context.WithValue(context.Background(), traceid.TraceIdFieldName, traceid.New())
		logx.WithContext(ctx).Info("定时任务", fmt.Sprintf("%s开始执行", cmd.Name()))

		start := time.Now()

		err := cmd.Handle(ctx)

		TimeStamp := time.Now()
		Cost := TimeStamp.Sub(start)
		if Cost > time.Minute {
			Cost = Cost.Truncate(time.Second)
		}
		if err != nil {
			logx.WithContext(ctx).Error("定时任务", fmt.Sprintf("%s执行结束,spec:%s,cost:%s,error=%s", cmd.Name(), spec, Cost.String(), err.Error()))
		} else {
			logx.WithContext(ctx).Info("定时任务", fmt.Sprintf("%s执行结束,spec:%s,cost:%s", cmd.Name(), spec, Cost.String()))
		}
	})
	if err != nil {
		logx.WithContext(context.Background()).Error("定时任务", fmt.Sprintf("%s添加失败,spec:%s", cmd.Name(), spec))
	}
}
