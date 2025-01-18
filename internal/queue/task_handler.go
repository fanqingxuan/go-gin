package queue

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
		fmt.Println("---" + string(t.Payload()))
		logx.QueueLoggerInstance.Info().Ctx(new_ctx).Str("task", t.Type()).Str("keywords", "开始执行").Any("payload", string(t.Payload())).Send()

		start := time.Now()
		err := h.handler(new_ctx, t.Payload())

		TimeStamp := time.Now()
		Cost := TimeStamp.Sub(start)
		if Cost > time.Minute {
			Cost = Cost.Truncate(time.Second)
		}
		if err != nil {
			logx.QueueLoggerInstance.Error().Ctx(new_ctx).Str("task", t.Type()).Str("keywords", "执行结束").Str("cost", Cost.String()).Str("err", err.Error()).Send()

		} else {
			logx.QueueLoggerInstance.Info().Ctx(new_ctx).Str("task", t.Type()).Str("keywords", "执行结束").Str("cost", Cost.String()).Send()

		}
		return err
	})
}
