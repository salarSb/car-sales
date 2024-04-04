package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/salarSb/car-sales/api/helper"
	"net/http"
)

// ErrorHandler  => only for development, don't show internal errors on production please
func ErrorHandler(c *gin.Context, err any) {
	if err, ok := err.(error); ok {
		httpResponse := helper.GenerateBaseResponseWithError(nil, false, helper.CustomRecovery, err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, httpResponse)
		return
	}
	httpResponse := helper.GenerateBaseResponseWithAnyError(nil, false, helper.CustomRecovery, err)
	c.AbortWithStatusJSON(http.StatusInternalServerError, httpResponse)
}
