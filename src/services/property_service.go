package services

import (
	"context"
	"github.com/salarSb/car-sales/api/dto"
	"github.com/salarSb/car-sales/config"
	"github.com/salarSb/car-sales/data/models"
)

type PropertyService struct {
	base *BaseService[
		models.Property,
		dto.CreatePropertyRequest,
		dto.UpdatePropertyRequest,
		dto.PropertyResponse,
	]
}

func NewPropertyService(cfg *config.Config) *PropertyService {
	return &PropertyService{
		base: NewBaseService[
			models.Property,
			dto.CreatePropertyRequest,
			dto.UpdatePropertyRequest,
			dto.PropertyResponse,
		](
			cfg,
			[]preload{{string: "PropertyCategory"}},
		),
	}
}

func (s *PropertyService) Create(ctx context.Context, req *dto.CreatePropertyRequest) (*dto.PropertyResponse, error) {
	return s.base.Create(ctx, req)
}

func (s *PropertyService) Update(ctx context.Context, id int, req *dto.UpdatePropertyRequest) (*dto.PropertyResponse, error) {
	return s.base.Update(ctx, id, req)
}

func (s *PropertyService) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

func (s *PropertyService) GetById(id int) (*dto.PropertyResponse, error) {
	return s.base.GetById(id)
}

func (s *PropertyService) GetByFilter(req *dto.PaginationInputWithFilter) (*dto.PagedList[dto.PropertyResponse], error) {
	return s.base.GetByFilter(req)
}
