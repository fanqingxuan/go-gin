# Event 事件系统

事件定义在 `event/` 目录，监听器定义在 `event/listener/` 目录。

## 定义事件 (event/)

```go
// event/sample_event.go
package event

import "go-gin/internal/eventbus"

var SampleEventName eventbus.EventName = "event.sample"

func NewSampleEvent(user string) *eventbus.Event {
    return eventbus.NewEvent(SampleEventName, user)
}
```

事件 payload 支持任意类型：

```go
// event/demo_event.go
var DemoEventName eventbus.EventName = "event.demo"

func NewDemoEvent(u *entity.User) *eventbus.Event {
    return eventbus.NewEvent(DemoEventName, u)
}
```

## 定义监听器 (event/listener/)

监听器实现 `eventbus.Listener` 接口：`Handle(context.Context, *eventbus.Event) error`

```go
// event/listener/sample_a_listener.go
package listener

import (
    "context"
    "go-gin/internal/eventbus"
)

type SampleAListener struct{}

func (l *SampleAListener) Handle(ctx context.Context, e *eventbus.Event) error {
    user := eventbus.PayloadAs[string](e)  // 泛型提取 payload
    // 业务逻辑
    return nil
}
```

结构体类型 payload 提取：

```go
// event/listener/demo_a_listener.go
func (l *DemoAListener) Handle(ctx context.Context, e *eventbus.Event) error {
    user := eventbus.PayloadAs[*entity.User](e)
    // 使用 user.Name 等
    return nil
}
```

## 注册监听器 (event/init.go)

```go
package event

import (
    "go-gin/event/listener"
    "go-gin/internal/eventbus"
)

func Init() {
    // 一个事件可绑定多个监听器
    eventbus.AddListener(SampleEventName, &listener.SampleAListener{}, &listener.SampleBListener{})
    eventbus.AddListener(DemoEventName, &listener.DemoAListener{})
}
```

## 触发事件

```go
// 同步触发（监听器依次执行，遇 error 停止）
eventbus.Fire(ctx, event.NewSampleEvent("张三"))

// 异步触发（每个监听器独立 goroutine）
eventbus.FireAsync(ctx, event.NewDemoEvent(user))

// 条件触发
eventbus.FireIf(ctx, condition, event.NewSampleEvent("test"))
eventbus.FireAsyncIf(ctx, condition, event.NewDemoEvent(user))

// 也可通过事件对象直接触发
event.NewSampleEvent("test").Fire(ctx)
event.NewSampleEvent("test").FireAsync(ctx)
```
