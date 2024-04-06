package services

import (
	"context"
	"database/sql"
	"github.com/salarSb/car-sales/common"
	"github.com/salarSb/car-sales/config"
	"github.com/salarSb/car-sales/constants"
	"github.com/salarSb/car-sales/data/db"
	"github.com/salarSb/car-sales/pkg/logging"
	"github.com/salarSb/car-sales/pkg/service_errors"
	"gorm.io/gorm"
	"time"
)

type BaseService[T any, Tc any, Tu any, Tr any] struct {
	Database *gorm.DB
	Logger   logging.Logger
}

func NewBseService[T any, Tc any, Tu any, Tr any](cfg *config.Config) *BaseService[T, Tc, Tu, Tr] {
	return &BaseService[T, Tc, Tu, Tr]{
		Database: db.GetDb(),
		Logger:   logging.NewLogger(cfg),
	}
}

func (b *BaseService[T, Tc, Tu, Tr]) Create(ctx context.Context, req *Tc) (*Tr, error) {
	model, err := common.TypeConverter[T](req)
	if err != nil {
		return nil, err
	}
	tx := b.Database.WithContext(ctx).Begin()
	err = tx.Create(model).Error
	if err != nil {
		tx.Rollback()
		b.Logger.Error(logging.Postgres, logging.Insert, err.Error(), nil)
		return nil, err
	}
	tx.Commit()
	return common.TypeConverter[Tr](model)
}

func (b *BaseService[T, Tc, Tu, Tr]) Update(ctx context.Context, id int, req *Tu) (*Tr, error) {
	updateMap, err := common.TypeConverter[map[string]interface{}](req)
	if err != nil {
		return nil, err
	}
	(*updateMap)["modified_by"] = &sql.NullInt64{Int64: int64(ctx.Value(constants.UserIdKey).(float64)), Valid: true}
	(*updateMap)["modified_at"] = sql.NullTime{Valid: true, Time: time.Now().UTC()}
	model := new(T)
	tx := b.Database.WithContext(ctx).Begin()
	if err = tx.
		Model(model).
		Where("id = ? AND deleted_by is null AND deleted_at is null", id).
		Updates(*updateMap).
		Error; err != nil {
		tx.Rollback()
		b.Logger.Error(logging.Postgres, logging.Update, err.Error(), nil)
		return nil, err
	}
	tx.Commit()
	return b.GetById(id)
}

func (b *BaseService[T, Tc, Tu, Tr]) Delete(ctx context.Context, id int) error {
	uid := ctx.Value(constants.UserIdKey)
	if uid == nil {
		return &service_errors.ServiceError{EndUserMessage: service_errors.PermissionDenied}
	}
	tx := b.Database.WithContext(ctx).Begin()
	model := new(T)
	deleteMap := map[string]interface{}{
		"deleted_by": &sql.NullInt64{Int64: int64(uid.(float64)), Valid: true},
		"deleted_at": sql.NullTime{Valid: true, Time: time.Now().UTC()},
	}
	if cnt := tx.
		Model(model).
		Where("id = ? AND deleted_by is null AND deleted_at is null", id).
		Updates(deleteMap).
		RowsAffected; cnt == 0 {
		tx.Rollback()
		b.Logger.Error(logging.Postgres, logging.Delete, service_errors.RecordNotFound, nil)
		return &service_errors.ServiceError{EndUserMessage: service_errors.RecordNotFound}
	}
	tx.Commit()
	return nil
}

func (b *BaseService[T, Tc, Tu, Tr]) GetById(id int) (*Tr, error) {
	model := new(T)
	err := b.Database.Where("id = ? AND deleted_by is null AND deleted_at is null", id).First(model).Error
	if err != nil {
		return nil, err
	}
	return common.TypeConverter[Tr](model)
}
