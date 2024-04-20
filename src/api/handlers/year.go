package handlers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/salarSb/car-sales/api/dto"
	_ "github.com/salarSb/car-sales/api/helper"
	"github.com/salarSb/car-sales/config"
	"github.com/salarSb/car-sales/services"
)

type YearHandler struct {
	service *services.YearService
}

func NewYearHandler(cfg *config.Config) *YearHandler {
	return &YearHandler{service: services.NewYearService(cfg)}
}

// Create Year godoc
// @Summary Create a Year
// @Description Create a Year
// @Tags Years
// @Accept json
// @Produce json
// @Param Request body dto.CreateYearRequest true "Create a year"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.YearResponse} "Year response"
// @Failure 422 {object} helper.BaseHttpResponse "Unprocessable Entity"
// @Failure 409 {object} helper.BaseHttpResponse "Status Conflict"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Failure 500 {object} helper.BaseHttpResponse "Internal Server Error"
// @Router /v1/years/ [post]
// @Security AuthBearer
func (h *YearHandler) Create(c *gin.Context) {
	Create(c, h.service.Create)
}

// Update year godoc
// @Summary Update a year
// @Description Update a year
// @Tags Years
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Param Request body dto.UpdateYearRequest true "Update a year"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.YearResponse} "Year response"
// @Failure 422 {object} helper.BaseHttpResponse "Unprocessable Entity"
// @Failure 409 {object} helper.BaseHttpResponse "Status Conflict"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Failure 404 {object} helper.BaseHttpResponse "Not Found"
// @Failure 500 {object} helper.BaseHttpResponse "Internal Server Error"
// @Router /v1/years/{id} [put]
// @Security AuthBearer
func (h *YearHandler) Update(c *gin.Context) {
	Update(c, h.service.Update)
}

// Delete Year godoc
// @Summary Delete a year
// @Description Delete a year
// @Tags Years
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} helper.BaseHttpResponse "response"
// @Failure 409 {object} helper.BaseHttpResponse "Status Conflict"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Failure 404 {object} helper.BaseHttpResponse "Not Found"
// @Failure 500 {object} helper.BaseHttpResponse "Internal Server Error"
// @Router /v1/years/{id} [delete]
// @Security AuthBearer
func (h *YearHandler) Delete(c *gin.Context) {
	Delete(c, h.service.Delete)
}

// GetById Year godoc
// @Summary Get a year
// @Description Get a year
// @Tags Years
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.YearResponse} "Year response"
// @Failure 404 {object} helper.BaseHttpResponse "Not Found"
// @Failure 500 {object} helper.BaseHttpResponse "Internal Server Error"
// @Router /v1/years/{id} [get]
// @Security AuthBearer
func (h *YearHandler) GetById(c *gin.Context) {
	GetById(c, h.service.GetById)
}

// GetByFilter Year godoc
// @Summary Get years
// @Description Get years
// @Tags Years
// @Accept json
// @Produce json
// @Param Request  body dto.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.PagedList[dto.YearResponse]} "Year response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Failure 500 {object} helper.BaseHttpResponse "Internal Server Error"
// @Router /v1/years/get-by-filter [post]
// @Security AuthBearer
func (h *YearHandler) GetByFilter(c *gin.Context) {
	GetByFilter(c, h.service.GetByFilter)
}
