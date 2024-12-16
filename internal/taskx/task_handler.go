package taskx

import (
	"context"
	"fmt"
	"go-gin/internal/components/logx"
	"go-gin/internal/traceid"
	"time"

	"github.com/hibiken/asynq"
)

type TaskHandler struct {
	taskName string
	handler  func(context.Context, []byte) error
}

func NewTaskHandler(taskName string, handler func(context.Context, []byte) error) *TaskHandler {
	return &TaskHandler{
		taskName: taskName,
		handler:  handler,
	}
}

func AddHandler(h *TaskHandler) {
	mux.HandleFunc(h.taskName, func(ctx context.Context, t *asynq.Task) error {
		new_ctx := context.WithValue(ctx, traceid.TraceIdFieldName, traceid.New())
		logx.WithContext(new_ctx).Info("队列", fmt.Sprintf("开始执行,task:%s,payload:%s", t.Type(), string(t.Payload())))
		start := time.Now()

		err := h.handler(new_ctx, t.Payload())

		TimeStamp := time.Now()
		Cost := TimeStamp.Sub(start)
		if Cost > time.Minute {
			Cost = Cost.Truncate(time.Second)
		}
		if err != nil {
			logx.WithContext(new_ctx).Error("队列", fmt.Sprintf("执行结束,task:%s,cost:%s,error=%s", t.Type(), Cost.String(), err.Error()))
		} else {
			logx.WithContext(new_ctx).Info("队列", fmt.Sprintf("执行结束,task:%s,cost:%s", t.Type(), Cost.String()))
		}
		return err
	})
}
