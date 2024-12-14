package middlewares

import (
	"fmt"
	"go-gin/internal/httpx"
)

func BeforeSampleB() httpx.HandlerFunc {
	return func(ctx *httpx.Context) (interface{}, error) {
		fmt.Println("BeforeSampleB")
		return nil, nil
	}
}
