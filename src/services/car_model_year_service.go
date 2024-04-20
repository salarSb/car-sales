package services

import (
	"context"
	"github.com/salarSb/car-sales/api/dto"
	"github.com/salarSb/car-sales/config"
	"github.com/salarSb/car-sales/data/models"
)

type CarModelYearService struct {
	base *BaseService[models.CarModelYear, dto.CreateCarModelYearRequest, dto.UpdateCarModelYearRequest, dto.CarModelYearResponse]
}

func NewCarModelYearService(cfg *config.Config) *CarModelYearService {
	return &CarModelYearService{
		base: NewBaseService[models.CarModelYear, dto.CreateCarModelYearRequest, dto.UpdateCarModelYearRequest, dto.CarModelYearResponse](
			cfg,
			[]preload{{string: "Year"}},
		),
	}
}

func (s *CarModelYearService) Create(ctx context.Context, req *dto.CreateCarModelYearRequest) (*dto.CarModelYearResponse, error) {
	return s.base.Create(ctx, req)
}

func (s *CarModelYearService) Update(ctx context.Context, id int, req *dto.UpdateCarModelYearRequest) (*dto.CarModelYearResponse, error) {
	return s.base.Update(ctx, id, req)
}

func (s *CarModelYearService) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

func (s *CarModelYearService) GetById(id int) (*dto.CarModelYearResponse, error) {
	return s.base.GetById(id)
}

func (s *CarModelYearService) GetByFilter(req *dto.PaginationInputWithFilter) (*dto.PagedList[dto.CarModelYearResponse], error) {
	return s.base.GetByFilter(req)
}
