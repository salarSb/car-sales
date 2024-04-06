package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/salarSb/car-sales/api/dto"
	"github.com/salarSb/car-sales/api/helper"
	"github.com/salarSb/car-sales/config"
	"github.com/salarSb/car-sales/services"
	"net/http"
	"strconv"
)

type CountryHandler struct {
	countryService *services.CountryService
}

func NewCountryHandler(cfg *config.Config) *CountryHandler {
	return &CountryHandler{countryService: services.NewCountryService(cfg)}
}

// Create Country godoc
// @Summary Create a country
// @Description Create a country
// @Tags Countries
// @Accept json
// @Produce json
// @Param Request body dto.CountryRequest true "Create a country"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.CountryResponse} "Country response"
// @Failure 422 {object} helper.GenerateBaseResponseWithValidationError "Unprocessable Entity"
// @Failure 409 {object} helper.GenerateBaseResponseWithError "Status Conflict"
// @Failure 400 {object} helper.GenerateBaseResponseWithError "Bad Request"
// @Failure 500 {object} helper.GenerateBaseResponseWithError "Internal Server Error"
// @Router /v1/countries/ [post]
// @Security AuthBearer
func (h *CountryHandler) Create(c *gin.Context) {
	req := new(dto.CountryRequest)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusUnprocessableEntity,
			helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, err),
		)
		return
	}
	res, err := h.countryService.Create(c, req)
	if err != nil {
		c.AbortWithStatusJSON(
			helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err),
		)
		return
	}
	c.JSON(http.StatusCreated, helper.GenerateBaseResponse(res, true, helper.Success))
}

// Update Country godoc
// @Summary Update a country
// @Description Update a country
// @Tags Countries
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Param Request body dto.CountryRequest true "Update a country"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.CountryResponse} "Country response"
// @Failure 422 {object} helper.GenerateBaseResponseWithValidationError "Unprocessable Entity"
// @Failure 409 {object} helper.GenerateBaseResponseWithError "Status Conflict"
// @Failure 400 {object} helper.GenerateBaseResponseWithError "Bad Request"
// @Failure 404 {object} helper.GenerateBaseResponseWithError "Not Found"
// @Failure 500 {object} helper.GenerateBaseResponseWithError "Internal Server Error"
// @Router /v1/countries/{id} [put]
// @Security AuthBearer
func (h *CountryHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	req := new(dto.CountryRequest)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusUnprocessableEntity,
			helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, err),
		)
		return
	}
	res, err := h.countryService.Update(c, id, req)
	if err != nil {
		c.AbortWithStatusJSON(
			helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err),
		)
		return
	}
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(res, true, helper.Success))
}

// Delete Country godoc
// @Summary Delete a country
// @Description Delete a country
// @Tags Countries
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} helper.BaseHttpResponse "response"
// @Failure 409 {object} helper.GenerateBaseResponseWithError "Status Conflict"
// @Failure 400 {object} helper.GenerateBaseResponseWithError "Bad Request"
// @Failure 404 {object} helper.GenerateBaseResponseWithError "Not Found"
// @Failure 500 {object} helper.GenerateBaseResponseWithError "Internal Server Error"
// @Router /v1/countries/{id} [delete]
// @Security AuthBearer
func (h *CountryHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	err := h.countryService.Delete(c, id)
	if err != nil {
		c.AbortWithStatusJSON(
			helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err),
		)
		return
	}
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(nil, true, helper.Success))
}

// GetById Country godoc
// @Summary Update a country
// @Description Update a country
// @Tags Countries
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.CountryResponse} "Country response"
// @Failure 404 {object} helper.GenerateBaseResponseWithError "Not Found"
// @Failure 500 {object} helper.GenerateBaseResponseWithError "Internal Server Error"
// @Router /v1/countries/{id} [get]
// @Security AuthBearer
func (h *CountryHandler) GetById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	res, err := h.countryService.GetById(c, id)
	if err != nil {
		c.AbortWithStatusJSON(
			helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err),
		)
		return
	}
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(res, true, helper.Success))
}
