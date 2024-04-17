package services

import (
	"context"
	"github.com/salarSb/car-sales/api/dto"
	"github.com/salarSb/car-sales/config"
	"github.com/salarSb/car-sales/data/models"
)

type CarModelService struct {
	base *BaseService[models.CarModel, dto.CreateCarModelRequest, dto.UpdateCarModelRequest, dto.CarModelResponse]
}

func NewCarModelService(cfg *config.Config) *CarModelService {
	return &CarModelService{
		base: NewBaseService[models.CarModel, dto.CreateCarModelRequest, dto.UpdateCarModelRequest, dto.CarModelResponse](
			cfg,
			[]preload{
				{string: "Gearbox"},
				{string: "CarType"},
				{string: "Company.Country"},
				{string: "CarModelColors.Color"},
				{string: "CarModelYears.Year"},
				{string: "CarModelYears.CarModelPriceHistories"},
				{string: "CarModelProperties.Property.PropertyCategory"},
				{string: "CarModelComments.User"},
				{string: "CarModelFiles.File"},
			},
		),
	}
}

func (s *CarModelService) Create(ctx context.Context, req *dto.CreateCarModelRequest) (*dto.CarModelResponse, error) {
	return s.base.Create(ctx, req)
}

func (s *CarModelService) Update(ctx context.Context, id int, req *dto.UpdateCarModelRequest) (*dto.CarModelResponse, error) {
	return s.base.Update(ctx, id, req)
}

func (s *CarModelService) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

func (s *CarModelService) GetById(id int) (*dto.CarModelResponse, error) {
	return s.base.GetById(id)
}

func (s *CarModelService) GetByFilter(req *dto.PaginationInputWithFilter) (*dto.PagedList[dto.CarModelResponse], error) {
	return s.base.GetByFilter(req)
}
