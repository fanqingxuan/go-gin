package middlewares

import "go-gin/internal/httpx"

func Init(r *httpx.Engine) {
	r.Use(dbCheck())
	r.Use(recoverLog())
}
