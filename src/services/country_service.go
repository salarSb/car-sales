package services

import (
	"context"
	"database/sql"
	"github.com/salarSb/car-sales/api/dto"
	"github.com/salarSb/car-sales/config"
	"github.com/salarSb/car-sales/constants"
	"github.com/salarSb/car-sales/data/db"
	"github.com/salarSb/car-sales/data/models"
	"github.com/salarSb/car-sales/pkg/logging"
	"gorm.io/gorm"
	"time"
)

type CountryService struct {
	database *gorm.DB
	logger   logging.Logger
}

func NewCountryService(cfg *config.Config) *CountryService {
	return &CountryService{
		database: db.GetDb(),
		logger:   logging.NewLogger(cfg),
	}
}

func (s *CountryService) Create(ctx context.Context, req *dto.CountryRequest) (*dto.CountryResponse, error) {
	country := models.Country{Name: req.Name}
	country.CreatedBy = int(ctx.Value(constants.UserIdKey).(float64))
	country.CreatedAt = time.Now().UTC()
	tx := s.database.WithContext(ctx).Begin()
	err := tx.Create(&country).Error
	if err != nil {
		tx.Rollback()
		s.logger.Error(logging.Postgres, logging.Insert, err.Error(), nil)
		return nil, err
	}
	tx.Commit()
	response := &dto.CountryResponse{Id: country.Id, Name: country.Name}
	return response, nil
}

func (s *CountryService) Update(ctx context.Context, id int, req *dto.CountryRequest) (*dto.CountryResponse, error) {
	updateMap := map[string]interface{}{
		"Name":        req.Name,
		"modified_by": &sql.NullInt64{Int64: int64(ctx.Value(constants.UserIdKey).(float64)), Valid: true},
		"modified_at": sql.NullTime{Valid: true, Time: time.Now().UTC()},
	}
	tx := s.database.WithContext(ctx).Begin()
	err := tx.Model(&models.Country{}).Where("id = ?", id).Updates(updateMap).Error
	if err != nil {
		tx.Rollback()
		s.logger.Error(logging.Postgres, logging.Update, err.Error(), nil)
		return nil, err
	}
	country := &models.Country{}
	err = tx.
		Model(&models.Country{}).
		Where("id = ? AND deleted_by is null AND deleted_at is null", id).
		First(&country).
		Error
	if err != nil {
		tx.Rollback()
		s.logger.Error(logging.Postgres, logging.Select, err.Error(), nil)
		return nil, err
	}
	tx.Commit()
	response := &dto.CountryResponse{Id: country.Id, Name: country.Name}
	return response, nil
}

func (s *CountryService) Delete(ctx context.Context, id int) error {
	deleteMap := map[string]interface{}{
		"deleted_by": &sql.NullInt64{Int64: int64(ctx.Value(constants.UserIdKey).(float64)), Valid: true},
		"deleted_at": sql.NullTime{Valid: true, Time: time.Now().UTC()},
	}
	tx := s.database.WithContext(ctx).Begin()
	if err := tx.Model(&models.Country{}).Where("id = ?", id).Updates(deleteMap).Error; err != nil {
		tx.Rollback()
		s.logger.Error(logging.Postgres, logging.Delete, err.Error(), nil)
		return err
	}
	tx.Commit()
	return nil
}

func (s *CountryService) GetById(ctx context.Context, id int) (*dto.CountryResponse, error) {
	country := &models.Country{}
	err := s.database.Where("id = ? AND deleted_by is null AND deleted_at is null", id).First(&country).Error
	if err != nil {
		s.logger.Error(logging.Postgres, logging.Select, err.Error(), nil)
		return nil, err
	}
	response := &dto.CountryResponse{Id: country.Id, Name: country.Name}
	return response, nil
}
