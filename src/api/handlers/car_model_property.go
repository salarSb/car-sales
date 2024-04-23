package handlers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/salarSb/car-sales/api/dto"
	_ "github.com/salarSb/car-sales/api/helper"
	"github.com/salarSb/car-sales/config"
	"github.com/salarSb/car-sales/services"
)

type CarModelPropertyHandler struct {
	service *services.CarModelPropertyService
}

func NewCarModelPropertyHandler(cfg *config.Config) *CarModelPropertyHandler {
	return &CarModelPropertyHandler{service: services.NewCarModelPropertyService(cfg)}
}

// Create CarModelProperty godoc
// @Summary Create a car model property
// @Description Create a car model property
// @Tags CarModelProperties
// @Accept json
// @Produce json
// @Param Request body dto.CreateCarModelPropertyRequest true "Create a car model property"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.CarModelPropertyResponse} "CarModelProperty response"
// @Failure 422 {object} helper.BaseHttpResponse "Unprocessable Entity"
// @Failure 409 {object} helper.BaseHttpResponse "Status Conflict"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Failure 500 {object} helper.BaseHttpResponse "Internal Server Error"
// @Router /v1/car-model-properties/ [post]
// @Security AuthBearer
func (h *CarModelPropertyHandler) Create(c *gin.Context) {
	Create(c, h.service.Create)
}

// Update CarModelProperty godoc
// @Summary Update a car model property
// @Description Update a car model property
// @Tags CarModelProperties
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Param Request body dto.UpdateCarModelPropertyRequest true "Update a car model property"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.CarModelPropertyResponse} "CarModelProperty response"
// @Failure 422 {object} helper.BaseHttpResponse "Unprocessable Entity"
// @Failure 409 {object} helper.BaseHttpResponse "Status Conflict"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Failure 404 {object} helper.BaseHttpResponse "Not Found"
// @Failure 500 {object} helper.BaseHttpResponse "Internal Server Error"
// @Router /v1/car-model-properties/{id} [put]
// @Security AuthBearer
func (h *CarModelPropertyHandler) Update(c *gin.Context) {
	Update(c, h.service.Update)
}

// Delete CarModelProperty godoc
// @Summary Delete a car model property
// @Description Delete a car model property
// @Tags CarModelProperties
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} helper.BaseHttpResponse "response"
// @Failure 409 {object} helper.BaseHttpResponse "Status Conflict"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Failure 404 {object} helper.BaseHttpResponse "Not Found"
// @Failure 500 {object} helper.BaseHttpResponse "Internal Server Error"
// @Router /v1/car-model-properties/{id} [delete]
// @Security AuthBearer
func (h *CarModelPropertyHandler) Delete(c *gin.Context) {
	Delete(c, h.service.Delete)
}

// GetById CarModelProperty godoc
// @Summary Get a car model property
// @Description Get a car model property
// @Tags CarModelProperties
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.CarModelPropertyResponse} "CarModelProperty response"
// @Failure 404 {object} helper.BaseHttpResponse "Not Found"
// @Failure 500 {object} helper.BaseHttpResponse "Internal Server Error"
// @Router /v1/car-model-properties/{id} [get]
// @Security AuthBearer
func (h *CarModelPropertyHandler) GetById(c *gin.Context) {
	GetById(c, h.service.GetById)
}

// GetByFilter CarModelProperty godoc
// @Summary Get car model properties
// @Description Get car model properties
// @Tags CarModelProperties
// @Accept json
// @Produce json
// @Param Request  body dto.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.PagedList[dto.CarModelPropertyResponse]} "CarModelProperty response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Failure 500 {object} helper.BaseHttpResponse "Internal Server Error"
// @Router /v1/car-model-properties/get-by-filter [post]
// @Security AuthBearer
func (h *CarModelPropertyHandler) GetByFilter(c *gin.Context) {
	GetByFilter(c, h.service.GetByFilter)
}
