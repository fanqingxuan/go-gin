package middleware

import (
	"fmt"
	"go-gin/internal/httpx"
)

func Xxx() httpx.HandlerFunc {
	return func(ctx *httpx.Context) (any, error) {
		return nil, nil
	}
}
