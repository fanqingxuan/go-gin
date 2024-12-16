package taskx

import (
	"go-gin/internal/components/redisx"

	"github.com/hibiken/asynq"
)

var (
	client *asynq.Client
)

func Init(c redisx.Config) {
	client = asynq.NewClient(asynq.RedisClientOpt{
		Addr:     c.Addr,
		Username: c.Username,
		Password: c.Password,
		DB:       c.DB})
}

func Close() {
	defer client.Close()
}
