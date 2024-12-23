package middleware

import (
	"fmt"
	"go-gin/internal/httpx"
)

func BeforeSampleA() httpx.HandlerFunc {
	return func(ctx *httpx.Context) (any, error) {
		fmt.Println("BeforeSampleA")
		return nil, nil
	}
}
