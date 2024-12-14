package httpx

import (
	"go-gin/internal/components/logx"
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
		TimeStamp := time.Now()
		Cost := TimeStamp.Sub(start)
		if Cost > time.Minute {
			Cost = Cost.Truncate(time.Second)
		}

		logx.AccessLoggerInstance.Info().Ctx(c).
			Str("path", path).
			Str("method", c.Request.Method).
			Str("ip", c.ClientIP()).
			Str("cost", Cost.String()).
			Int("status", c.Writer.Status()).
			Str("proto", c.Request.Proto).
			Str("user_agent", c.Request.UserAgent()).
			Send()
	}
}
