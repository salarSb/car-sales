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
// @Summary SendOtp
// @description SendOtp
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
			helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, err),
		)
		return
	}
	err = h.userService.SendOtp(req)
	if err != nil {
		c.AbortWithStatusJSON(
			helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err),
		)
		return
	}
	// Call internal sms service
	c.JSON(http.StatusCreated, helper.GenerateBaseResponse("otp sent", true, helper.Success))
}

// LoginByUsername godoc
// @Summary LoginByUsername
// @description LoginByUsername
// @Tags Users
// @Accept json
// @Produce json
// @Param Request body dto.LoginByUsernameRequest true "LoginByUsernameRequest"
// @Success 201 {object} helper.BaseHttpResponse "Success"
// @Failure 422 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 500 {object} helper.BaseHttpResponse "Failed"
// Router /v1/users/login-by-username [post]
func (h *UserHandler) LoginByUsername(c *gin.Context) {
	req := new(dto.LoginByUsernameRequest)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusUnprocessableEntity,
			helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, err),
		)
		return
	}
	token, err := h.userService.LoginByUsername(req)
	if err != nil {
		c.AbortWithStatusJSON(
			helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err),
		)
		return
	}
	c.JSON(http.StatusCreated, helper.GenerateBaseResponse(token, true, helper.Success))
}

// RegisterByUsername godoc
// @Summary RegisterByUsername
// @description RegisterByUsername
// @Tags Users
// @Accept json
// @Produce json
// @Param Request body dto.RegisterByUsernameRequest true "RegisterByUsernameRequest"
// @Success 201 {object} helper.BaseHttpResponse "Success"
// @Failure 422 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 500 {object} helper.BaseHttpResponse "Failed"
// Router /v1/users/register-by-username [post]
func (h *UserHandler) RegisterByUsername(c *gin.Context) {
	req := new(dto.RegisterByUsernameRequest)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusUnprocessableEntity,
			helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, err),
		)
		return
	}
	err = h.userService.RegisterByUsername(req)
	if err != nil {
		c.AbortWithStatusJSON(
			helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err),
		)
		return
	}
	c.JSON(http.StatusCreated, helper.GenerateBaseResponse(nil, true, helper.Success))
}

// RegisterLoginByMobileNumber godoc
// @Summary RegisterLoginByMobileNumber
// @description RegisterLoginByMobileNumber
// @Tags Users
// @Accept json
// @Produce json
// @Param Request body dto.RegisterLoginByMobileRequest true "RegisterLoginByMobileRequest"
// @Success 201 {object} helper.BaseHttpResponse "Success"
// @Failure 422 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 500 {object} helper.BaseHttpResponse "Failed"
// Router /v1/users/login-by-mobile [post]
func (h *UserHandler) RegisterLoginByMobileNumber(c *gin.Context) {
	req := new(dto.RegisterLoginByMobileRequest)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusUnprocessableEntity,
			helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, err),
		)
		return
	}
	token, err := h.userService.RegisterLoginByMobileNumber(req)
	if err != nil {
		c.AbortWithStatusJSON(
			helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err),
		)
		return
	}
	c.JSON(http.StatusCreated, helper.GenerateBaseResponse(token, true, helper.Success))
}
