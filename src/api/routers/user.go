package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/salarSb/car-sales/api/handlers"
	"github.com/salarSb/car-sales/config"
)

func User(router *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewUserHandler(cfg)
	router.POST("/send-otp", h.SendOtp)
}
