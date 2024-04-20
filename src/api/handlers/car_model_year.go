package handlers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/salarSb/car-sales/api/dto"
	_ "github.com/salarSb/car-sales/api/helper"
	"github.com/salarSb/car-sales/config"
	"github.com/salarSb/car-sales/services"
)

type CarModelYearHandler struct {
	service *services.CarModelYearService
}

func NewCarModelYearHandler(cfg *config.Config) *CarModelYearHandler {
	return &CarModelYearHandler{service: services.NewCarModelYearService(cfg)}
}

// Create CarModelYear godoc
// @Summary Create a car model year
// @Description Create a car model year
// @Tags CarModelYears
// @Accept json
// @Produce json
// @Param Request body dto.CreateCarModelYearRequest true "Create a car model year"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.CarModelYearResponse} "CarModelYear response"
// @Failure 422 {object} helper.BaseHttpResponse "Unprocessable Entity"
// @Failure 409 {object} helper.BaseHttpResponse "Status Conflict"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Failure 500 {object} helper.BaseHttpResponse "Internal Server Error"
// @Router /v1/car-model-years/ [post]
// @Security AuthBearer
func (h *CarModelYearHandler) Create(c *gin.Context) {
	Create(c, h.service.Create)
}

// Update CarModelYear godoc
// @Summary Update a car model year
// @Description Update a car model year
// @Tags CarModelYears
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Param Request body dto.UpdateCarModelYearRequest true "Update a car model year"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.CarModelYearResponse} "CarModelYear response"
// @Failure 422 {object} helper.BaseHttpResponse "Unprocessable Entity"
// @Failure 409 {object} helper.BaseHttpResponse "Status Conflict"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Failure 404 {object} helper.BaseHttpResponse "Not Found"
// @Failure 500 {object} helper.BaseHttpResponse "Internal Server Error"
// @Router /v1/car-model-years/{id} [put]
// @Security AuthBearer
func (h *CarModelYearHandler) Update(c *gin.Context) {
	Update(c, h.service.Update)
}

// Delete CarModelYear godoc
// @Summary Delete a car model year
// @Description Delete a car model year
// @Tags CarModelYears
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} helper.BaseHttpResponse "response"
// @Failure 409 {object} helper.BaseHttpResponse "Status Conflict"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Failure 404 {object} helper.BaseHttpResponse "Not Found"
// @Failure 500 {object} helper.BaseHttpResponse "Internal Server Error"
// @Router /v1/car-model-years/{id} [delete]
// @Security AuthBearer
func (h *CarModelYearHandler) Delete(c *gin.Context) {
	Delete(c, h.service.Delete)
}

// GetById CarModelYear godoc
// @Summary Get a car model year
// @Description Get a car model year
// @Tags CarModelYears
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.CarModelYearResponse} "CarModelYear response"
// @Failure 404 {object} helper.BaseHttpResponse "Not Found"
// @Failure 500 {object} helper.BaseHttpResponse "Internal Server Error"
// @Router /v1/car-model-years/{id} [get]
// @Security AuthBearer
func (h *CarModelYearHandler) GetById(c *gin.Context) {
	GetById(c, h.service.GetById)
}

// GetByFilter CarModelYear godoc
// @Summary Get car model years
// @Description Get car model years
// @Tags CarModelYears
// @Accept json
// @Produce json
// @Param Request  body dto.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.PagedList[dto.CarModelYearResponse]} "CarModelYear response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Failure 500 {object} helper.BaseHttpResponse "Internal Server Error"
// @Router /v1/car-model-years/get-by-filter [post]
// @Security AuthBearer
func (h *CarModelYearHandler) GetByFilter(c *gin.Context) {
	GetByFilter(c, h.service.GetByFilter)
}
