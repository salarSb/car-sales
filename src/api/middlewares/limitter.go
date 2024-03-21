package middlewares

import (
	"github.com/didip/tollbooth"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LimitByRequestMiddleware() gin.HandlerFunc {
	lmt := tollbooth.NewLimiter(1, nil)
	return func(context *gin.Context) {
		err := tollbooth.LimitByRequest(lmt, context.Writer, context.Request)
		if err != nil {
			context.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": err.Error(),
			})
			return
		}
		context.Next()
	}
}
