package services

import (
	"context"
	"github.com/salarSb/car-sales/api/dto"
	"github.com/salarSb/car-sales/config"
	"github.com/salarSb/car-sales/data/models"
)

type CarTypeService struct {
	base *BaseService[models.CarType, dto.CarTypeRequest, dto.CarTypeRequest, dto.CarTypeResponse]
}

func NewCarTypeService(cfg *config.Config) *CarTypeService {
	return &CarTypeService{
		base: NewBaseService[models.CarType, dto.CarTypeRequest, dto.CarTypeRequest, dto.CarTypeResponse](
			cfg,
			[]preload{},
		),
	}
}

func (s *CarTypeService) Create(ctx context.Context, req *dto.CarTypeRequest) (*dto.CarTypeResponse, error) {
	return s.base.Create(ctx, req)
}

func (s *CarTypeService) Update(ctx context.Context, id int, req *dto.CarTypeRequest) (*dto.CarTypeResponse, error) {
	return s.base.Update(ctx, id, req)
}

func (s *CarTypeService) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

func (s *CarTypeService) GetById(id int) (*dto.CarTypeResponse, error) {
	return s.base.GetById(id)
}

func (s *CarTypeService) GetByFilter(req *dto.PaginationInputWithFilter) (*dto.PagedList[dto.CarTypeResponse], error) {
	return s.base.GetByFilter(req)
}
