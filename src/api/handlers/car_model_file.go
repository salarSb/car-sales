package handlers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/salarSb/car-sales/api/dto"
	_ "github.com/salarSb/car-sales/api/helper"
	"github.com/salarSb/car-sales/config"
	"github.com/salarSb/car-sales/services"
)

type CarModelFileHandler struct {
	service *services.CarModelFileService
}

func NewCarModelFileHandler(cfg *config.Config) *CarModelFileHandler {
	return &CarModelFileHandler{service: services.NewCarModelFileService(cfg)}
}

// Create CarModelFile godoc
// @Summary Create a car model file
// @Description Create a car model file
// @Tags CarModelFiles
// @Accept json
// @Produce json
// @Param Request body dto.CreateCarModelFileRequest true "Create a car model file"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.CarModelFileResponse} "CarModelFile response"
// @Failure 422 {object} helper.BaseHttpResponse "Unprocessable Entity"
// @Failure 409 {object} helper.BaseHttpResponse "Status Conflict"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Failure 500 {object} helper.BaseHttpResponse "Internal Server Error"
// @Router /v1/car-model-files/ [post]
// @Security AuthBearer
func (h *CarModelFileHandler) Create(c *gin.Context) {
	Create(c, h.service.Create)
}

// Update CarModelFile godoc
// @Summary Update a car model file
// @Description Update a car model file
// @Tags CarModelFiles
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Param Request body dto.UpdateCarModelFileRequest true "Update a car model file"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.CarModelFileResponse} "CarModelFile response"
// @Failure 422 {object} helper.BaseHttpResponse "Unprocessable Entity"
// @Failure 409 {object} helper.BaseHttpResponse "Status Conflict"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Failure 404 {object} helper.BaseHttpResponse "Not Found"
// @Failure 500 {object} helper.BaseHttpResponse "Internal Server Error"
// @Router /v1/car-model-files/{id} [put]
// @Security AuthBearer
func (h *CarModelFileHandler) Update(c *gin.Context) {
	Update(c, h.service.Update)
}

// Delete CarModelFile godoc
// @Summary Delete a car model file
// @Description Delete a car model file
// @Tags CarModelFiles
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} helper.BaseHttpResponse "response"
// @Failure 409 {object} helper.BaseHttpResponse "Status Conflict"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Failure 404 {object} helper.BaseHttpResponse "Not Found"
// @Failure 500 {object} helper.BaseHttpResponse "Internal Server Error"
// @Router /v1/car-model-files/{id} [delete]
// @Security AuthBearer
func (h *CarModelFileHandler) Delete(c *gin.Context) {
	Delete(c, h.service.Delete)
}

// GetById CarModelFile godoc
// @Summary Get a car model file
// @Description Get a car model file
// @Tags CarModelFiles
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.CarModelFileResponse} "CarModelFile response"
// @Failure 404 {object} helper.BaseHttpResponse "Not Found"
// @Failure 500 {object} helper.BaseHttpResponse "Internal Server Error"
// @Router /v1/car-model-files/{id} [get]
// @Security AuthBearer
func (h *CarModelFileHandler) GetById(c *gin.Context) {
	GetById(c, h.service.GetById)
}

// GetByFilter CarModelFile godoc
// @Summary Get car model files
// @Description Get car model files
// @Tags CarModelFiles
// @Accept json
// @Produce json
// @Param Request  body dto.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.PagedList[dto.CarModelFileResponse]} "CarModelFile response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Failure 500 {object} helper.BaseHttpResponse "Internal Server Error"
// @Router /v1/car-model-files/get-by-filter [post]
// @Security AuthBearer
func (h *CarModelFileHandler) GetByFilter(c *gin.Context) {
	GetByFilter(c, h.service.GetByFilter)
}
