package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/salarSb/car-sales/api/dto"
	"github.com/salarSb/car-sales/api/helper"
	"github.com/salarSb/car-sales/config"
	"github.com/salarSb/car-sales/services"
	"net/http"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(cfg *config.Config) *UserHandler {
	userService := services.NewUserService(cfg)
	return &UserHandler{userService: userService}
}

// SendOtp godoc
// @Summary Send otp to user
// @description Send otp to user
// @Tags Users
// @Accept json
// @Produce json
// @Param Request body dto.GetOtpRequest true "GetOtpRequest"
// @Success 201 {object} helper.BaseHttpResponse "Success"
// @Failure 422 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 500 {object} helper.BaseHttpResponse "Failed"
// Router /v1/users/send-otp [post]
func (h *UserHandler) SendOtp(c *gin.Context) {
	req := new(dto.GetOtpRequest)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusUnprocessableEntity,
			helper.GenerateBaseResponseWithValidationError(nil, false, -1, err),
		)
		return
	}
	err = h.userService.SendOtp(req)
	if err != nil {
		c.AbortWithStatusJSON(
			helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, -1, err),
		)
		return
	}
	// Call internal sms service
	c.JSON(http.StatusCreated, helper.GenerateBaseResponse("otp sent", true, 0))
}
