package services

import (
	"context"
	"github.com/salarSb/car-sales/api/dto"
	"github.com/salarSb/car-sales/config"
	"github.com/salarSb/car-sales/data/models"
)

type ColorService struct {
	base *BaseService[models.Color, dto.CreateColorRequest, dto.UpdateColorRequest, dto.ColorResponse]
}

func NewColorService(cfg *config.Config) *ColorService {
	return &ColorService{
		base: NewBaseService[models.Color, dto.CreateColorRequest, dto.UpdateColorRequest, dto.ColorResponse](
			cfg,
			[]preload{},
		),
	}
}

func (s *ColorService) Create(ctx context.Context, req *dto.CreateColorRequest) (*dto.ColorResponse, error) {
	return s.base.Create(ctx, req)
}

func (s *ColorService) Update(ctx context.Context, id int, req *dto.UpdateColorRequest) (*dto.ColorResponse, error) {
	return s.base.Update(ctx, id, req)
}

func (s *ColorService) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

func (s *ColorService) GetById(id int) (*dto.ColorResponse, error) {
	return s.base.GetById(id)
}

func (s *ColorService) GetByFilter(req *dto.PaginationInputWithFilter) (*dto.PagedList[dto.ColorResponse], error) {
	return s.base.GetByFilter(req)
}
