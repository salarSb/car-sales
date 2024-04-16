package services

import (
	"context"
	"github.com/salarSb/car-sales/api/dto"
	"github.com/salarSb/car-sales/config"
	"github.com/salarSb/car-sales/data/models"
)

type CityService struct {
	base *BaseService[models.City, dto.CreateCityRequest, dto.UpdateCityRequest, dto.CityResponse]
}

func NewCityService(cfg *config.Config) *CityService {
	return &CityService{
		base: NewBaseService[models.City, dto.CreateCityRequest, dto.UpdateCityRequest, dto.CityResponse](
			cfg,
			[]preload{{string: "Country"}},
		),
	}
}

func (s *CityService) Create(ctx context.Context, req *dto.CreateCityRequest) (*dto.CityResponse, error) {
	return s.base.Create(ctx, req)
}

func (s *CityService) Update(ctx context.Context, id int, req *dto.UpdateCityRequest) (*dto.CityResponse, error) {
	return s.base.Update(ctx, id, req)
}

func (s *CityService) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

func (s *CityService) GetById(id int) (*dto.CityResponse, error) {
	return s.base.GetById(id)
}

func (s *CityService) GetByFilter(req *dto.PaginationInputWithFilter) (*dto.PagedList[dto.CityResponse], error) {
	return s.base.GetByFilter(req)
}
