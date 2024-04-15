package services

import (
	"context"
	"github.com/salarSb/car-sales/api/dto"
	"github.com/salarSb/car-sales/config"
	"github.com/salarSb/car-sales/data/models"
)

type PropertyCategoryService struct {
	base *BaseService[
		models.PropertyCategory,
		dto.CreatePropertyCategoryRequest,
		dto.UpdatePropertyCategoryRequest,
		dto.PropertyCategoryResponse,
	]
}

func NewPropertyCategoryService(cfg *config.Config) *PropertyCategoryService {
	return &PropertyCategoryService{
		base: NewBseService[
			models.PropertyCategory,
			dto.CreatePropertyCategoryRequest,
			dto.UpdatePropertyCategoryRequest,
			dto.PropertyCategoryResponse,
		](
			cfg,
			[]preload{{string: "Properties"}},
		),
	}
}

func (s *PropertyCategoryService) Create(ctx context.Context, req *dto.CreatePropertyCategoryRequest) (*dto.PropertyCategoryResponse, error) {
	return s.base.Create(ctx, req)
}

func (s *PropertyCategoryService) Update(ctx context.Context, id int, req *dto.UpdatePropertyCategoryRequest) (*dto.PropertyCategoryResponse, error) {
	return s.base.Update(ctx, id, req)
}

func (s *PropertyCategoryService) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

func (s *PropertyCategoryService) GetById(id int) (*dto.PropertyCategoryResponse, error) {
	return s.base.GetById(id)
}

func (s *PropertyCategoryService) GetByFilter(req *dto.PaginationInputWithFilter) (*dto.PagedList[dto.PropertyCategoryResponse], error) {
	return s.base.GetByFilter(req)
}
