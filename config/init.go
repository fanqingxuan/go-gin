package config

import (
	"fmt"
	filex "go-gin/internal/file"
)

var instance Config

func Init(filename string) {
	filex.MustLoad(filename, &instance)
	fmt.Println(instance)
}

func IsDebugMode() bool {
	return instance.App.Debug
}

func Port() string {
	return instance.App.Port
}
