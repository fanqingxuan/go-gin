package middleware

import (
	"fmt"
	"go-gin/internal/httpx"
)

func BeforeSampleB() httpx.HandlerFunc {
	return func(ctx *httpx.Context) (any, error) {
		fmt.Println("BeforeSampleB")
		return nil, nil
	}
}
