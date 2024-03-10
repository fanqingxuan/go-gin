package middlewares

import (
	"fmt"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
)

func requestLog() gin.HandlerFunc {

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
		fmt.Println(raw)
		TimeStamp := time.Now()
		Cost := TimeStamp.Sub(start)
		if Cost > time.Minute {
			Cost = Cost.Truncate(time.Second)
		}

		requestMap := map[string]interface{}{
			"Path":     path,
			"Method":   c.Request.Method,
			"ClientIP": c.ClientIP(),
			"Cost":     Cost.String(),
			"Status":   c.Writer.Status(),
			"Proto":    c.Request.Proto,
		}

		fmt.Println(requestMap)
	}
}
