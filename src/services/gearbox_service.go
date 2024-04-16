package services

import (
	"context"
	"github.com/salarSb/car-sales/api/dto"
	"github.com/salarSb/car-sales/config"
	"github.com/salarSb/car-sales/data/models"
)

type GearboxService struct {
	base *BaseService[models.Gearbox, dto.GearboxRequest, dto.GearboxRequest, dto.GearboxResponse]
}

func NewGearboxService(cfg *config.Config) *GearboxService {
	return &GearboxService{
		base: NewBaseService[models.Gearbox, dto.GearboxRequest, dto.GearboxRequest, dto.GearboxResponse](
			cfg,
			[]preload{},
		),
	}
}

func (s *GearboxService) Create(ctx context.Context, req *dto.GearboxRequest) (*dto.GearboxResponse, error) {
	return s.base.Create(ctx, req)
}

func (s *GearboxService) Update(ctx context.Context, id int, req *dto.GearboxRequest) (*dto.GearboxResponse, error) {
	return s.base.Update(ctx, id, req)
}

func (s *GearboxService) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

func (s *GearboxService) GetById(id int) (*dto.GearboxResponse, error) {
	return s.base.GetById(id)
}

func (s *GearboxService) GetByFilter(req *dto.PaginationInputWithFilter) (*dto.PagedList[dto.GearboxResponse], error) {
	return s.base.GetByFilter(req)
}
