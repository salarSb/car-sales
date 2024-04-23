package services

import (
	"context"
	"github.com/salarSb/car-sales/api/dto"
	"github.com/salarSb/car-sales/config"
	"github.com/salarSb/car-sales/constants"
	"github.com/salarSb/car-sales/data/models"
)

type CarModelCommentService struct {
	base *BaseService[models.CarModelComment, dto.CreateCarModelCommentRequest, dto.UpdateCarModelCommentRequest, dto.CarModelCommentResponse]
}

func NewCarModelCommentService(cfg *config.Config) *CarModelCommentService {
	return &CarModelCommentService{
		base: NewBaseService[models.CarModelComment, dto.CreateCarModelCommentRequest, dto.UpdateCarModelCommentRequest, dto.CarModelCommentResponse](
			cfg,
			[]preload{{string: "User"}},
		),
	}
}

func (s *CarModelCommentService) Create(ctx context.Context, req *dto.CreateCarModelCommentRequest) (*dto.CarModelCommentResponse, error) {
	req.UserId = int(ctx.Value(constants.UserIdKey).(float64))
	return s.base.Create(ctx, req)
}

func (s *CarModelCommentService) Update(ctx context.Context, id int, req *dto.UpdateCarModelCommentRequest) (*dto.CarModelCommentResponse, error) {
	return s.base.Update(ctx, id, req)
}

func (s *CarModelCommentService) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

func (s *CarModelCommentService) GetById(id int) (*dto.CarModelCommentResponse, error) {
	return s.base.GetById(id)
}

func (s *CarModelCommentService) GetByFilter(req *dto.PaginationInputWithFilter) (*dto.PagedList[dto.CarModelCommentResponse], error) {
	return s.base.GetByFilter(req)
}
