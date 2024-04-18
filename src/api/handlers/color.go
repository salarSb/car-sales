package handlers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/salarSb/car-sales/api/dto"
	_ "github.com/salarSb/car-sales/api/helper"
	"github.com/salarSb/car-sales/config"
	"github.com/salarSb/car-sales/services"
)

type ColorHandler struct {
	service *services.ColorService
}

func NewColorHandler(cfg *config.Config) *ColorHandler {
	return &ColorHandler{service: services.NewColorService(cfg)}
}

// Create Color godoc
// @Summary Create a color
// @Description Create a color
// @Tags Colors
// @Accept json
// @Produce json
// @Param Request body dto.CreateColorRequest true "Create a color"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.ColorResponse} "Color response"
// @Failure 422 {object} helper.BaseHttpResponse "Unprocessable Entity"
// @Failure 409 {object} helper.BaseHttpResponse "Status Conflict"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Failure 500 {object} helper.BaseHttpResponse "Internal Server Error"
// @Router /v1/colors/ [post]
// @Security AuthBearer
func (h *ColorHandler) Create(c *gin.Context) {
	Create(c, h.service.Create)
}

// Update Color godoc
// @Summary Update a color
// @Description Update a color
// @Tags Colors
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Param Request body dto.UpdateColorRequest true "Update a color"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.ColorResponse} "Color response"
// @Failure 422 {object} helper.BaseHttpResponse "Unprocessable Entity"
// @Failure 409 {object} helper.BaseHttpResponse "Status Conflict"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Failure 404 {object} helper.BaseHttpResponse "Not Found"
// @Failure 500 {object} helper.BaseHttpResponse "Internal Server Error"
// @Router /v1/colors/{id} [put]
// @Security AuthBearer
func (h *ColorHandler) Update(c *gin.Context) {
	Update(c, h.service.Update)
}

// Delete Color godoc
// @Summary Delete a color
// @Description Delete a color
// @Tags Colors
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} helper.BaseHttpResponse "response"
// @Failure 409 {object} helper.BaseHttpResponse "Status Conflict"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Failure 404 {object} helper.BaseHttpResponse "Not Found"
// @Failure 500 {object} helper.BaseHttpResponse "Internal Server Error"
// @Router /v1/colors/{id} [delete]
// @Security AuthBearer
func (h *ColorHandler) Delete(c *gin.Context) {
	Delete(c, h.service.Delete)
}

// GetById Color godoc
// @Summary Get a color
// @Description Get a color
// @Tags Colors
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.ColorResponse} "Color response"
// @Failure 404 {object} helper.BaseHttpResponse "Not Found"
// @Failure 500 {object} helper.BaseHttpResponse "Internal Server Error"
// @Router /v1/colors/{id} [get]
// @Security AuthBearer
func (h *ColorHandler) GetById(c *gin.Context) {
	GetById(c, h.service.GetById)
}

// GetByFilter Color godoc
// @Summary Get colors
// @Description Get colors
// @Tags Colors
// @Accept json
// @Produce json
// @Param Request  body dto.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.PagedList[dto.ColorResponse]} "Color response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Failure 500 {object} helper.BaseHttpResponse "Internal Server Error"
// @Router /v1/colors/get-by-filter [post]
// @Security AuthBearer
func (h *ColorHandler) GetByFilter(c *gin.Context) {
	GetByFilter(c, h.service.GetByFilter)
}
