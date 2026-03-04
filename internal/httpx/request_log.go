package httpx

import (
	"go-gin/internal/component/logx"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
)

func RequestLog() gin.HandlerFunc {

	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path
		rawQuery := c.Request.URL.RawQuery
		raw, _ := url.QueryUnescape(rawQuery)
		// Process request
		c.Next()

		if raw != "" {
			path = path + "?" + raw
		}
		timestamp := time.Now()
		cost := timestamp.Sub(start)
		if cost > time.Minute {
			cost = cost.Truncate(time.Second)
		}

		logx.AccessLogger.Info().Ctx(c).
			Str("path", path).
			Str("method", c.Request.Method).
			Str("ip", c.ClientIP()).
			Str("cost", cost.String()).
			Int("status", c.Writer.Status()).
			Str("proto", c.Request.Proto).
			Str("user_agent", c.Request.UserAgent()).
			Send()
	}
}
