package redisx

import "github.com/redis/go-redis/v9"

type Redisx struct {
	*redis.Client
}

func New() *Redisx {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "host.docker.internal:6379",
		Password: "", // 没有密码，默认值
		DB:       0,  // 默认DB 0
	})
	return &Redisx{
		rdb,
	}
}
