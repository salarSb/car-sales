package services

import (
	"context"
	"github.com/salarSb/car-sales/api/dto"
	"github.com/salarSb/car-sales/config"
	"github.com/salarSb/car-sales/data/models"
)

type CarModelColorService struct {
	base *BaseService[models.CarModelColor, dto.CreateCarModelColorRequest, dto.UpdateCarModelColorRequest, dto.CarModelColorResponse]
}

func NewCarModelColorService(cfg *config.Config) *CarModelColorService {
	return &CarModelColorService{
		base: NewBaseService[models.CarModelColor, dto.CreateCarModelColorRequest, dto.UpdateCarModelColorRequest, dto.CarModelColorResponse](
			cfg,
			[]preload{{string: "Color"}},
		),
	}
}

func (s *CarModelColorService) Create(ctx context.Context, req *dto.CreateCarModelColorRequest) (*dto.CarModelColorResponse, error) {
	return s.base.Create(ctx, req)
}

func (s *CarModelColorService) Update(ctx context.Context, id int, req *dto.UpdateCarModelColorRequest) (*dto.CarModelColorResponse, error) {
	return s.base.Update(ctx, id, req)
}

func (s *CarModelColorService) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

func (s *CarModelColorService) GetById(id int) (*dto.CarModelColorResponse, error) {
	return s.base.GetById(id)
}

func (s *CarModelColorService) GetByFilter(req *dto.PaginationInputWithFilter) (*dto.PagedList[dto.CarModelColorResponse], error) {
	return s.base.GetByFilter(req)
}
