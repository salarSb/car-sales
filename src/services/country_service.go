package services

import (
	"context"
	"github.com/salarSb/car-sales/api/dto"
	"github.com/salarSb/car-sales/config"
	"github.com/salarSb/car-sales/data/models"
)

type CountryService struct {
	base *BaseService[models.Country, dto.CountryRequest, dto.CountryRequest, dto.CountryResponse]
}

func NewCountryService(cfg *config.Config) *CountryService {
	return &CountryService{
		base: NewBaseService[models.Country, dto.CountryRequest, dto.CountryRequest, dto.CountryResponse](
			cfg,
			[]preload{{string: "Cities"}, {string: "Cities.Country"}, {string: "Companies"}, {string: "Companies.Country"}},
		),
	}
}

func (s *CountryService) Create(ctx context.Context, req *dto.CountryRequest) (*dto.CountryResponse, error) {
	return s.base.Create(ctx, req)
}

func (s *CountryService) Update(ctx context.Context, id int, req *dto.CountryRequest) (*dto.CountryResponse, error) {
	return s.base.Update(ctx, id, req)
}

func (s *CountryService) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

func (s *CountryService) GetById(id int) (*dto.CountryResponse, error) {
	return s.base.GetById(id)
}

func (s *CountryService) GetByFilter(req *dto.PaginationInputWithFilter) (*dto.PagedList[dto.CountryResponse], error) {
	return s.base.GetByFilter(req)
}
