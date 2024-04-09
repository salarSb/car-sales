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

type CityHandler struct {
	service *services.CityService
}

func NewCityHandler(cfg *config.Config) *CityHandler {
	return &CityHandler{service: services.NewCityService(cfg)}
}

// Create City godoc
// @Summary Create a City
// @Description Create a City
// @Tags Cities
// @Accept json
// @Produce json
// @Param Request body dto.CreateCityRequest true "Create a city"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.CityResponse} "City response"
// @Failure 422 {object} helper.BaseHttpResponse "Unprocessable Entity"
// @Failure 409 {object} helper.BaseHttpResponse "Status Conflict"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Failure 500 {object} helper.BaseHttpResponse "Internal Server Error"
// @Router /v1/cities/ [post]
// @Security AuthBearer
func (h *CityHandler) Create(c *gin.Context) {
	req := new(dto.CreateCityRequest)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusUnprocessableEntity,
			helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, err),
		)
		return
	}
	res, err := h.service.Create(c, req)
	if err != nil {
		c.AbortWithStatusJSON(
			helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err),
		)
		return
	}
	c.JSON(http.StatusCreated, helper.GenerateBaseResponse(res, true, helper.Success))
}

// Update city godoc
// @Summary Update a city
// @Description Update a city
// @Tags Cities
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Param Request body dto.UpdateCityRequest true "Update a city"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.CityResponse} "City response"
// @Failure 422 {object} helper.BaseHttpResponse "Unprocessable Entity"
// @Failure 409 {object} helper.BaseHttpResponse "Status Conflict"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Failure 404 {object} helper.BaseHttpResponse "Not Found"
// @Failure 500 {object} helper.BaseHttpResponse "Internal Server Error"
// @Router /v1/cities/{id} [put]
// @Security AuthBearer
func (h *CityHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	req := new(dto.UpdateCityRequest)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusUnprocessableEntity,
			helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, err),
		)
		return
	}
	res, err := h.service.Update(c, id, req)
	if err != nil {
		c.AbortWithStatusJSON(
			helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err),
		)
		return
	}
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(res, true, helper.Success))
}

// Delete City godoc
// @Summary Delete a city
// @Description Delete a city
// @Tags Cities
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} helper.BaseHttpResponse "response"
// @Failure 409 {object} helper.BaseHttpResponse "Status Conflict"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Failure 404 {object} helper.BaseHttpResponse "Not Found"
// @Failure 500 {object} helper.BaseHttpResponse "Internal Server Error"
// @Router /v1/cities/{id} [delete]
// @Security AuthBearer
func (h *CityHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	err := h.service.Delete(c, id)
	if err != nil {
		c.AbortWithStatusJSON(
			helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err),
		)
		return
	}
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(nil, true, helper.Success))
}

// GetById City godoc
// @Summary Get a city
// @Description Get a city
// @Tags Cities
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.CityResponse} "City response"
// @Failure 404 {object} helper.BaseHttpResponse "Not Found"
// @Failure 500 {object} helper.BaseHttpResponse "Internal Server Error"
// @Router /v1/cities/{id} [get]
// @Security AuthBearer
func (h *CityHandler) GetById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	res, err := h.service.GetById(id)
	if err != nil {
		c.AbortWithStatusJSON(
			helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err),
		)
		return
	}
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(res, true, helper.Success))
}

// GetByFilter City godoc
// @Summary Get cities
// @Description Get cities
// @Tags Cities
// @Accept json
// @Produce json
// @Param Request  body dto.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.PagedList[dto.CityResponse]} "City response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Failure 500 {object} helper.BaseHttpResponse "Internal Server Error"
// @Router /v1/cities/get-by-filter [post]
// @Security AuthBearer
func (h *CityHandler) GetByFilter(c *gin.Context) {
	req := dto.PaginationInputWithFilter{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err),
		)
		return
	}
	res, err := h.service.GetByFilter(&req)
	if err != nil {
		c.AbortWithStatusJSON(
			helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err),
		)
		return
	}
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(res, true, helper.Success))
}
