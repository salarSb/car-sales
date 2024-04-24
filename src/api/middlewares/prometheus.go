package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/salarSb/car-sales/pkg/metrics"
	"strconv"
	"time"
)

func Prometheus() gin.HandlerFunc {
	return func(context *gin.Context) {
		start := time.Now()
		path := context.FullPath()
		method := context.Request.Method
		context.Next()
		status := context.Writer.Status()
		metrics.HttpDuration.WithLabelValues(path, method, strconv.Itoa(status)).Observe(float64(time.Since(start) / time.Millisecond))
	}
}
