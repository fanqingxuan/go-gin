package queue

import (
	"context"
	"go-gin/internal/component/logx"
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
		newCtx := context.WithValue(ctx, traceid.TraceIdFieldName, traceid.New())
		logx.QueueLoggerInstance.Info().Ctx(newCtx).Str("task", t.Type()).Str("keywords", "开始执行").Any("payload", string(t.Payload())).Send()

		start := time.Now()
		err := h.handler(newCtx, t.Payload())

		timestamp := time.Now()
		cost := timestamp.Sub(start)
		if cost > time.Minute {
			cost = cost.Truncate(time.Second)
		}
		if err != nil {
			logx.QueueLoggerInstance.Error().Ctx(newCtx).Str("task", t.Type()).Str("keywords", "执行结束").Str("cost", cost.String()).Str("err", err.Error()).Send()
		} else {
			logx.QueueLoggerInstance.Info().Ctx(newCtx).Str("task", t.Type()).Str("keywords", "执行结束").Str("cost", cost.String()).Send()
		}
		return err
	})
}
