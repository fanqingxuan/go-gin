package task

import (
	"go-gin/internal/components/redisx"

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
	mux.HandleFunc(h.typename, h.handler)
}
func Start() error {
	return srv.Run(mux)
}
