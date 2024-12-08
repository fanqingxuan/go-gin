package main

import (
	"fmt"
	"go-gin/internal/environment"
	"go-gin/internal/httpx"
	"go-gin/internal/httpx/validators"
)

type Config struct {
	Port string
}

var conf Config

func InitConfig(c Config) {
	conf = c
}

func InitValidators() {
	validators.Init()

}

func InitServer() *httpx.Engine {
	if environment.IsDebugMode() {
		httpx.SetDebugMode()
	} else {
		httpx.SetReleaseMode()
	}
	engine := httpx.Default()
	return engine
}

func StartServer(engine *httpx.Engine) {
	fmt.Printf("Starting server at localhost%s...\n", conf.Port)
	if err := engine.Run(conf.Port); err != nil {
		fmt.Printf("Start server error,err=%v", err)
	}
}
