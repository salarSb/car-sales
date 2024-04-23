package services

import (
	"context"
	"github.com/salarSb/car-sales/api/dto"
	"github.com/salarSb/car-sales/config"
	"github.com/salarSb/car-sales/data/models"
)

type CarModelFileService struct {
	base *BaseService[models.CarModelFile, dto.CreateCarModelFileRequest, dto.UpdateCarModelFileRequest, dto.CarModelFileResponse]
}

func NewCarModelFileService(cfg *config.Config) *CarModelFileService {
	return &CarModelFileService{
		base: NewBaseService[models.CarModelFile, dto.CreateCarModelFileRequest, dto.UpdateCarModelFileRequest, dto.CarModelFileResponse](
			cfg,
			[]preload{{string: "File"}},
		),
	}
}

func (s *CarModelFileService) Create(ctx context.Context, req *dto.CreateCarModelFileRequest) (*dto.CarModelFileResponse, error) {
	return s.base.Create(ctx, req)
}

func (s *CarModelFileService) Update(ctx context.Context, id int, req *dto.UpdateCarModelFileRequest) (*dto.CarModelFileResponse, error) {
	return s.base.Update(ctx, id, req)
}

func (s *CarModelFileService) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

func (s *CarModelFileService) GetById(id int) (*dto.CarModelFileResponse, error) {
	return s.base.GetById(id)
}

func (s *CarModelFileService) GetByFilter(req *dto.PaginationInputWithFilter) (*dto.PagedList[dto.CarModelFileResponse], error) {
	return s.base.GetByFilter(req)
}
