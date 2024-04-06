package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/salarSb/car-sales/api/handlers"
	"github.com/salarSb/car-sales/config"
)

func Country(r *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewCountryHandler(cfg)
	r.POST("/", h.Create)
	r.PUT("/:id", h.Update)
	r.DELETE("/:id", h.Delete)
	r.GET("/:id", h.GetById)
}
