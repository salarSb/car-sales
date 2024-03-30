package middlewares

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/salarSb/car-sales/api/helper"
	"github.com/salarSb/car-sales/config"
	"github.com/salarSb/car-sales/pkg/limiter"
	"golang.org/x/time/rate"
	"net/http"
	"time"
)

func OtpLimiter(cfg *config.Config) gin.HandlerFunc {
	var ipLimiter = limiter.NewIpRateLimiter(rate.Every(cfg.Otp.Limiter*time.Second), 1)
	return func(context *gin.Context) {
		ipLimiter := ipLimiter.GetLimiter(context.Request.RemoteAddr)
		if !ipLimiter.Allow() {
			context.AbortWithStatusJSON(
				http.StatusTooManyRequests,
				helper.GenerateBaseResponseWithError(
					nil,
					false,
					int(helper.OtpLimiterError),
					errors.New("not allowed"),
				),
			)
			context.Abort()
		} else {
			context.Next()
		}
	}
}
