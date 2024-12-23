package middleware

import (
	"fmt"
	"go-gin/internal/httpx"
)

func AfterSampleB() httpx.HandlerFunc {
	return func(ctx *httpx.Context) (any, error) {
		fmt.Println("AfterSampleB")
		return nil, nil
	}
}
