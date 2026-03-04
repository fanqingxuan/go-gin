package queue

import (
	"go-gin/internal/component/redisx"

	"github.com/hibiken/asynq"
)

var (
	srv *asynq.Server
	mux *asynq.ServeMux
)

const (
	QueueHigh   = "critical"
	QueueNormal = "default"
	QueueLow    = "low"
)

func InitServer(c redisx.Config) {
	srv = asynq.NewServer(
		RedisClientOpt(c),
		asynq.Config{
			// Specify how many concurrent workers to use
			Concurrency:    10,
			StrictPriority: true,
			// Optionally specify multiple queues with different priority.
			Queues: map[string]int{
				QueueHigh:   6,
				QueueNormal: 3,
				QueueLow:    1,
			},
			// See the godoc for other configuration options
		},
	)
	mux = asynq.NewServeMux()
}

func Start() error {
	return srv.Run(mux)
}
