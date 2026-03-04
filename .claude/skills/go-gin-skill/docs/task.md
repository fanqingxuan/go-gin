# Task 异步队列任务

任务定义在 `task/` 目录，基于 Asynq (Redis) 实现。

## 定义任务

每个任务包含：任务名常量、创建函数 (`NewXxxTask`)、处理函数 (`NewXxxTaskHandler`)。

### 简单 payload（字符串）

```go
// task/sample.go
package task

import (
    "context"
    "go-gin/internal/component/logx"
    "go-gin/internal/queue"
)

const TypeSampleTask = "sample"

func NewSampleTask(p string) *queue.Task {
    return queue.NewTask(TypeSampleTask, p)
}

func NewSampleTaskHandler() *queue.TaskHandler {
    return queue.NewTaskHandler(TypeSampleTask, func(ctx context.Context, data []byte) error {
        logx.WithContext(ctx).Debug("sample_task", string(data))
        return nil
    })
}
```

### 结构体 payload（JSON 序列化/反序列化）

```go
// task/sampleB.go
package task

import (
    "context"
    "encoding/json"
    "go-gin/internal/component/logx"
    "go-gin/internal/queue"
)

const TypeSampleBTask = "sampleB"

type SampleBTaskPayload struct {
    UserId []string
}

func NewSampleBTask(p string) *queue.Task {
    return queue.NewTask(TypeSampleBTask, SampleBTaskPayload{UserId: []string{p}})
}

func NewSampleBTaskHandler() *queue.TaskHandler {
    return queue.NewTaskHandler(TypeSampleBTask, func(ctx context.Context, data []byte) error {
        var p SampleBTaskPayload
        if err := json.Unmarshal(data, &p); err != nil {
            logx.WithContext(ctx).Error("sampleB_task_unmarshal", err.Error())
            return err
        }
        logx.WithContext(ctx).Debug("sampleB_task", p.UserId)
        return nil
    })
}
```

## 注册任务处理器 (task/init.go)

```go
package task

import "go-gin/internal/queue"

func Init() {
    queue.AddHandler(NewSampleTaskHandler())
    queue.AddHandler(NewSampleBTaskHandler())
}
```

## 分发任务

```go
// 立即执行
task.NewSampleTask("hello").DispatchNow()

// 延迟执行
task.NewSampleTask("hello").Dispatch(5 * time.Minute)

// 条件分发
task.NewSampleTask("hello").DispatchIf(shouldProcess)
```
