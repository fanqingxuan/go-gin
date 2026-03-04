# Cron 定时任务

定时任务定义在 `cron/` 目录，在 `cron/init.go` 中注册。

## 结构体方式（实现 cronx.Job 接口）

```go
// cron/db_check.go
package cron

import (
    "context"
    "go-gin/internal/component/db"
)

type DBCheckJob struct{}

func (j *DBCheckJob) Handle(ctx context.Context) error {
    if db.IsConnected() {
        return nil
    }
    return db.Connect()
}
```

## 函数方式

```go
// cron/sample_func.go
package cron

import "context"

func SampleFunc(ctx context.Context) error {
    // 业务逻辑
    return nil
}
```

## 注册任务 (cron/init.go)

```go
package cron

import "go-gin/internal/cronx"

func Init() {
    // cron 表达式方式
    cronx.AddJob("@every 3s", &DBCheckJob{})
    cronx.AddFunc("@every 5s", SampleFunc)

    // 流式 API 方式
    cronx.Schedule(&SampleJob{}).EveryMinute()
    cronx.Schedule(&SampleJob{}).DailyAt("08:30")
    cronx.ScheduleFunc(SampleFunc).EveryFiveMinutes()
}
```

## 流式调度方法

- 秒级: `EverySecond`, `EveryTwoSeconds`, `EveryFiveSeconds`, `EveryTenSeconds`, `EveryFifteenSeconds`, `EveryThirtySeconds`
- 分钟级: `EveryMinute`, `EveryTwoMinutes`, `EveryThreeMinutes`, `EveryFiveMinutes`, `EveryTenMinutes`, `EveryFifteenMinutes`, `EveryThirtyMinutes`
- 小时级: `Hourly`, `HourlyAt(minute)`, `EveryTwoHours`, `EveryThreeHours`, `EveryFourHours`, `EverySixHours`
- 天级: `Daily`, `DailyAt("HH:mm")`, `TwiceDaily(h1, h2)`, `TwiceDailyAt(h1, h2, min)`
- 周级: `Weekly`, `WeeklyOn(day, "HH:mm")`, `Weekdays`, `Weekends`, `Mondays`~`Sundays`
- 月级: `Monthly`, `MonthlyOn(day, "HH:mm")`, `TwiceMonthly(d1, d2, "HH:mm")`, `LastDayOfMonth("HH:mm")`
- 其他: `Quarterly`, `Yearly`, `Cron("expression")`
