package handlers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/salarSb/car-sales/api/dto"
	_ "github.com/salarSb/car-sales/api/helper"
	"github.com/salarSb/car-sales/config"
	"github.com/salarSb/car-sales/services"
)

type CarModelColorHandler struct {
	service *services.CarModelColorService
}

func NewCarModelColorHandler(cfg *config.Config) *CarModelColorHandler {
	return &CarModelColorHandler{service: services.NewCarModelColorService(cfg)}
}

// Create CarModelColor godoc
// @Summary Create a car model color
// @Description Create a car model color
// @Tags CarModelColors
// @Accept json
// @Produce json
// @Param Request body dto.CreateCarModelColorRequest true "Create a car model color"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.CarModelColorResponse} "CarModelColor response"
// @Failure 422 {object} helper.BaseHttpResponse "Unprocessable Entity"
// @Failure 409 {object} helper.BaseHttpResponse "Status Conflict"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Failure 500 {object} helper.BaseHttpResponse "Internal Server Error"
// @Router /v1/car-model-colors/ [post]
// @Security AuthBearer
func (h *CarModelColorHandler) Create(c *gin.Context) {
	Create(c, h.service.Create)
}

// Update CarModelColor godoc
// @Summary Update a car model color
// @Description Update a car model color
// @Tags CarModelColors
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Param Request body dto.UpdateCarModelColorRequest true "Update a car model color"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.CarModelColorResponse} "CarModelColor response"
// @Failure 422 {object} helper.BaseHttpResponse "Unprocessable Entity"
// @Failure 409 {object} helper.BaseHttpResponse "Status Conflict"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Failure 404 {object} helper.BaseHttpResponse "Not Found"
// @Failure 500 {object} helper.BaseHttpResponse "Internal Server Error"
// @Router /v1/car-model-colors/{id} [put]
// @Security AuthBearer
func (h *CarModelColorHandler) Update(c *gin.Context) {
	Update(c, h.service.Update)
}

// Delete CarModelColor godoc
// @Summary Delete a car model color
// @Description Delete a car model color
// @Tags CarModelColors
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} helper.BaseHttpResponse "response"
// @Failure 409 {object} helper.BaseHttpResponse "Status Conflict"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Failure 404 {object} helper.BaseHttpResponse "Not Found"
// @Failure 500 {object} helper.BaseHttpResponse "Internal Server Error"
// @Router /v1/car-model-colors/{id} [delete]
// @Security AuthBearer
func (h *CarModelColorHandler) Delete(c *gin.Context) {
	Delete(c, h.service.Delete)
}

// GetById CarModelColor godoc
// @Summary Get a car model color
// @Description Get a car model color
// @Tags CarModelColors
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.CarModelColorResponse} "CarModelColor response"
// @Failure 404 {object} helper.BaseHttpResponse "Not Found"
// @Failure 500 {object} helper.BaseHttpResponse "Internal Server Error"
// @Router /v1/car-model-colors/{id} [get]
// @Security AuthBearer
func (h *CarModelColorHandler) GetById(c *gin.Context) {
	GetById(c, h.service.GetById)
}

// GetByFilter CarModelColor godoc
// @Summary Get car model colors
// @Description Get car model colors
// @Tags CarModelColors
// @Accept json
// @Produce json
// @Param Request  body dto.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.PagedList[dto.CarModelColorResponse]} "CarModelColor response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Failure 500 {object} helper.BaseHttpResponse "Internal Server Error"
// @Router /v1/car-model-colors/get-by-filter [post]
// @Security AuthBearer
func (h *CarModelColorHandler) GetByFilter(c *gin.Context) {
	GetByFilter(c, h.service.GetByFilter)
}
