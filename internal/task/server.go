package task

import (
	"context"
	"fmt"
	"go-gin/internal/components/logx"
	"go-gin/internal/components/redisx"
	"go-gin/internal/traceid"
	"time"

	"github.com/hibiken/asynq"
)

type Server struct {
	server *asynq.Server
}

var (
	srv *asynq.Server
	mux *asynq.ServeMux
)

func InitServer(c redisx.Config) {
	srv = asynq.NewServer(
		asynq.RedisClientOpt{
			Addr:     c.Addr,
			Username: c.Username,
			Password: c.Password,
			DB:       c.DB,
		},
		asynq.Config{
			// Specify how many concurrent workers to use
			Concurrency: 10,
			// Optionally specify multiple queues with different priority.
			Queues: map[string]int{
				"critical": 6,
				"default":  3,
				"low":      1,
			},
			// See the godoc for other configuration options
		},
	)
	mux = asynq.NewServeMux()
}

func Handle(h *TaskHandler) {
	mux.HandleFunc(h.typename, func(ctx context.Context, t *asynq.Task) error {
		new_ctx := context.WithValue(ctx, traceid.TraceIdFieldName, traceid.New())
		logx.WithContext(new_ctx).Info("队列", fmt.Sprintf("开始执行,task:%s", t.Type()))
		start := time.Now()

		err := h.handler(new_ctx, t)

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
func Start() error {
	return srv.Run(mux)
}
