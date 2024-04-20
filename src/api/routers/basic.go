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
	r.POST("/get-by-filter", h.GetByFilter)
}

func City(r *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewCityHandler(cfg)
	r.POST("/", h.Create)
	r.PUT("/:id", h.Update)
	r.DELETE("/:id", h.Delete)
	r.GET("/:id", h.GetById)
	r.POST("/get-by-filter", h.GetByFilter)
}

func File(r *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewFileHandler(cfg)
	r.POST("/", h.Create)
	r.PUT("/:id", h.Update)
	r.DELETE("/:id", h.Delete)
	r.GET("/:id", h.GetById)
	r.POST("/get-by-filter", h.GetByFilter)
}

func Color(r *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewColorHandler(cfg)
	r.POST("/", h.Create)
	r.PUT("/:id", h.Update)
	r.DELETE("/:id", h.Delete)
	r.GET("/:id", h.GetById)
	r.POST("/get-by-filter", h.GetByFilter)
}

func CarModelColor(r *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewCarModelColorHandler(cfg)
	r.POST("/", h.Create)
	r.PUT("/:id", h.Update)
	r.DELETE("/:id", h.Delete)
	r.GET("/:id", h.GetById)
	r.POST("/get-by-filter", h.GetByFilter)
}

func Year(r *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewYearHandler(cfg)
	r.POST("/", h.Create)
	r.PUT("/:id", h.Update)
	r.DELETE("/:id", h.Delete)
	r.GET("/:id", h.GetById)
	r.POST("/get-by-filter", h.GetByFilter)
}
