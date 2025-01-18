package queue

import (
	"go-gin/internal/components/redisx"

	"github.com/hibiken/asynq"
)

var (
	srv *asynq.Server
	mux *asynq.ServeMux
)

var (
	HIGH   = "critical"
	NORMAL = "default"
	LOW    = "low"
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
			Concurrency:    10,
			StrictPriority: true,
			// Optionally specify multiple queues with different priority.
			Queues: map[string]int{
				HIGH:   6,
				NORMAL: 3,
				LOW:    1,
			},
			// See the godoc for other configuration options
		},
	)
	mux = asynq.NewServeMux()
}

func Start() error {
	return srv.Run(mux)
}
