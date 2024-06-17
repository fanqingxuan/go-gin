package redisx

import (
	"context"
	"go-gin/internal/components/logx"

	"github.com/redis/go-redis/v9"
)

var instance *redis.Client

func Init(options *redis.Options) {

	rdb := redis.NewClient(options)
	rdb.AddHook(&LogHook{})
	err := rdb.Ping(context.Background()).Err()
	if err != nil {
		logx.WithContext(context.Background()).Error("redis", err)
	}
	instance = rdb
}

func GetInstance() *redis.Client {
	return instance
}
