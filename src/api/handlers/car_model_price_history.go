package handlers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/salarSb/car-sales/api/dto"
	_ "github.com/salarSb/car-sales/api/helper"
	"github.com/salarSb/car-sales/config"
	"github.com/salarSb/car-sales/services"
)

type CarModelPriceHistoryHandler struct {
	service *services.CarModelPriceHistoryService
}

func NewCarModelPriceHistoryHandler(cfg *config.Config) *CarModelPriceHistoryHandler {
	return &CarModelPriceHistoryHandler{service: services.NewCarModelPriceHistoryService(cfg)}
}

// Create CarModelPriceHistory godoc
// @Summary Create a carModelPriceHistory
// @Description Create a carModelPriceHistory
// @Tags CarModelPriceHistories
// @Accept json
// @Produce json
// @Param Request body dto.CreateCarModelPriceHistoryRequest true "Create a carModelPriceHistory"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.CarModelPriceHistoryResponse} "CarModelPriceHistory response"
// @Failure 422 {object} helper.BaseHttpResponse "Unprocessable Entity"
// @Failure 409 {object} helper.BaseHttpResponse "Status Conflict"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Failure 500 {object} helper.BaseHttpResponse "Internal Server Error"
// @Router /v1/car-model-price-histories/ [post]
// @Security AuthBearer
func (h *CarModelPriceHistoryHandler) Create(c *gin.Context) {
	Create(c, h.service.Create)
}

// Update CarModelPriceHistory godoc
// @Summary Update a carModelPriceHistory
// @Description Update a carModelPriceHistory
// @Tags CarModelPriceHistories
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Param Request body dto.UpdateCarModelPriceHistoryRequest true "Update a carModelPriceHistory"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.CarModelPriceHistoryResponse} "CarModelPriceHistory response"
// @Failure 422 {object} helper.BaseHttpResponse "Unprocessable Entity"
// @Failure 409 {object} helper.BaseHttpResponse "Status Conflict"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Failure 404 {object} helper.BaseHttpResponse "Not Found"
// @Failure 500 {object} helper.BaseHttpResponse "Internal Server Error"
// @Router /v1/car-model-price-histories/{id} [put]
// @Security AuthBearer
func (h *CarModelPriceHistoryHandler) Update(c *gin.Context) {
	Update(c, h.service.Update)
}

// Delete CarModelPriceHistory godoc
// @Summary Delete a carModelPriceHistory
// @Description Delete a carModelPriceHistory
// @Tags CarModelPriceHistories
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} helper.BaseHttpResponse "response"
// @Failure 409 {object} helper.BaseHttpResponse "Status Conflict"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Failure 404 {object} helper.BaseHttpResponse "Not Found"
// @Failure 500 {object} helper.BaseHttpResponse "Internal Server Error"
// @Router /v1/car-model-price-histories/{id} [delete]
// @Security AuthBearer
func (h *CarModelPriceHistoryHandler) Delete(c *gin.Context) {
	Delete(c, h.service.Delete)
}

// GetById CarModelPriceHistory godoc
// @Summary Get a carModelPriceHistory
// @Description Get a carModelPriceHistory
// @Tags CarModelPriceHistories
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.CarModelPriceHistoryResponse} "CarModelPriceHistory response"
// @Failure 404 {object} helper.BaseHttpResponse "Not Found"
// @Failure 500 {object} helper.BaseHttpResponse "Internal Server Error"
// @Router /v1/car-model-price-histories/{id} [get]
// @Security AuthBearer
func (h *CarModelPriceHistoryHandler) GetById(c *gin.Context) {
	GetById(c, h.service.GetById)
}

// GetByFilter CarModelPriceHistory godoc
// @Summary Get car-model-price-histories
// @Description Get car-model-price-histories
// @Tags CarModelPriceHistories
// @Accept json
// @Produce json
// @Param Request  body dto.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.PagedList[dto.CarModelPriceHistoryResponse]} "CarModelPriceHistory response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Failure 500 {object} helper.BaseHttpResponse "Internal Server Error"
// @Router /v1/car-model-price-histories/get-by-filter [post]
// @Security AuthBearer
func (h *CarModelPriceHistoryHandler) GetByFilter(c *gin.Context) {
	GetByFilter(c, h.service.GetByFilter)
}
