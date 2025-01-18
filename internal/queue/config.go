package queue

import (
	"go-gin/internal/components/redisx"

	"github.com/hibiken/asynq"
)

func RedisClientOpt(c redisx.Config) asynq.RedisClientOpt {
	return asynq.RedisClientOpt{
		Addr:     c.Addr,
		Username: c.Username,
		Password: c.Password,
		DB:       c.DB,
	}
}
