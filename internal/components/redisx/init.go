package redisx

import (
	"context"
	"go-gin/internal/components/logx"

	"github.com/redis/go-redis/v9"
)

var (
	instance *redis.Client
	conf     Config
)

type Config struct {
	Addr     string `yaml:"addr"`
	Username string `yaml:"username"`
	Password string `yaml:"password"` // no password set
	DB       int    `yaml:"db"`       // use default DB
}

func InitConfig(c Config) {
	conf = c
}

func Init() {
	options := &redis.Options{
		Addr:     conf.Addr,
		Username: conf.Username,
		Password: conf.Password,
		DB:       conf.DB,
	}
	rdb := redis.NewClient(options)
	rdb.AddHook(&LogHook{})
	rdb.AddHook(&ErrHook{})
	err := rdb.Ping(context.Background()).Err()
	if err != nil {
		logx.WithContext(context.Background()).Error("redis", err)
	}
	instance = rdb
}

func Client() *redis.Client {
	return instance
}
