package services

import (
	"context"
	"github.com/salarSb/car-sales/api/dto"
	"github.com/salarSb/car-sales/config"
	"github.com/salarSb/car-sales/data/models"
)

type YearService struct {
	base *BaseService[models.Year, dto.CreateYearRequest, dto.UpdateYearRequest, dto.YearResponse]
}

func NewYearService(cfg *config.Config) *YearService {
	return &YearService{
		base: NewBaseService[models.Year, dto.CreateYearRequest, dto.UpdateYearRequest, dto.YearResponse](
			cfg,
			[]preload{},
		),
	}
}

func (s *YearService) Create(ctx context.Context, req *dto.CreateYearRequest) (*dto.YearResponse, error) {
	return s.base.Create(ctx, req)
}

func (s *YearService) Update(ctx context.Context, id int, req *dto.UpdateYearRequest) (*dto.YearResponse, error) {
	return s.base.Update(ctx, id, req)
}

func (s *YearService) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

func (s *YearService) GetById(id int) (*dto.YearResponse, error) {
	return s.base.GetById(id)
}

func (s *YearService) GetByFilter(req *dto.PaginationInputWithFilter) (*dto.PagedList[dto.YearResponse], error) {
	return s.base.GetByFilter(req)
}
