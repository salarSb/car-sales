package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/salarSb/car-sales/api/handlers"
)

func Health(routerGroup *gin.RouterGroup) {
	handler := handlers.NewHealthHandler()
	routerGroup.GET("/", handler.Health)
}
