package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var instance *redis.Client

func Init() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "host.docker.internal:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	err := rdb.Ping(context.Background()).Err()
	if err != nil {
		panic(err)
	}
	instance = rdb
}

func GetInstance() *redis.Client {
	return instance
}
