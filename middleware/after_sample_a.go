package middleware

import (
	"fmt"
	"go-gin/internal/httpx"
)

func AfterSampleA() httpx.HandlerFunc {
	return func(ctx *httpx.Context) (interface{}, error) {
		fmt.Println("AfterSampleA")
		return nil, nil
	}
}
