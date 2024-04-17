package handlers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/salarSb/car-sales/api/dto"
	_ "github.com/salarSb/car-sales/api/helper"
	"github.com/salarSb/car-sales/config"
	"github.com/salarSb/car-sales/services"
)

type CarModelHandler struct {
	service *services.CarModelService
}

func NewCarModelHandler(cfg *config.Config) *CarModelHandler {
	return &CarModelHandler{service: services.NewCarModelService(cfg)}
}

// Create CarModel godoc
// @Summary Create a car model
// @Description Create a car model
// @Tags CarModels
// @Accept json
// @Produce json
// @Param Request body dto.CreateCarModelRequest true "Create a car model"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.CarModelResponse} "CarModel response"
// @Failure 422 {object} helper.BaseHttpResponse "Unprocessable Entity"
// @Failure 409 {object} helper.BaseHttpResponse "Status Conflict"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Failure 500 {object} helper.BaseHttpResponse "Internal Server Error"
// @Router /v1/car-models/ [post]
// @Security AuthBearer
func (h *CarModelHandler) Create(c *gin.Context) {
	Create(c, h.service.Create)
}

// Update CarModel godoc
// @Summary Update a car model
// @Description Update a car model
// @Tags CarModels
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Param Request body dto.UpdateCarModelRequest true "Update a car model"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.CarModelResponse} "CarModel response"
// @Failure 422 {object} helper.BaseHttpResponse "Unprocessable Entity"
// @Failure 409 {object} helper.BaseHttpResponse "Status Conflict"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Failure 404 {object} helper.BaseHttpResponse "Not Found"
// @Failure 500 {object} helper.BaseHttpResponse "Internal Server Error"
// @Router /v1/car-models/{id} [put]
// @Security AuthBearer
func (h *CarModelHandler) Update(c *gin.Context) {
	Update(c, h.service.Update)
}

// Delete CarModel godoc
// @Summary Delete a car model
// @Description Delete a car model
// @Tags CarModels
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} helper.BaseHttpResponse "response"
// @Failure 409 {object} helper.BaseHttpResponse "Status Conflict"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Failure 404 {object} helper.BaseHttpResponse "Not Found"
// @Failure 500 {object} helper.BaseHttpResponse "Internal Server Error"
// @Router /v1/car-models/{id} [delete]
// @Security AuthBearer
func (h *CarModelHandler) Delete(c *gin.Context) {
	Delete(c, h.service.Delete)
}

// GetById CarModel godoc
// @Summary Get a car model
// @Description Get a car model
// @Tags CarModels
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.CarModelResponse} "CarModel response"
// @Failure 404 {object} helper.BaseHttpResponse "Not Found"
// @Failure 500 {object} helper.BaseHttpResponse "Internal Server Error"
// @Router /v1/car-models/{id} [get]
// @Security AuthBearer
func (h *CarModelHandler) GetById(c *gin.Context) {
	GetById(c, h.service.GetById)
}

// GetByFilter CarModel godoc
// @Summary Get car models
// @Description Get car models
// @Tags CarModels
// @Accept json
// @Produce json
// @Param Request  body dto.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.PagedList[dto.CarModelResponse]} "CarModel response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Failure 500 {object} helper.BaseHttpResponse "Internal Server Error"
// @Router /v1/car-models/get-by-filter [post]
// @Security AuthBearer
func (h *CarModelHandler) GetByFilter(c *gin.Context) {
	GetByFilter(c, h.service.GetByFilter)
}
