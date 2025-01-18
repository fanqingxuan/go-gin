package queue

import (
	"go-gin/internal/components/redisx"

	"github.com/hibiken/asynq"
)

var (
	client *asynq.Client
)

func Init(c redisx.Config) {
	client = asynq.NewClient(RedisClientOpt(c))
}

func Close() {
	defer client.Close()
}
