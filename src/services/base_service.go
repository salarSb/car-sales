package services

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/salarSb/car-sales/api/dto"
	"github.com/salarSb/car-sales/common"
	"github.com/salarSb/car-sales/config"
	"github.com/salarSb/car-sales/constants"
	"github.com/salarSb/car-sales/data/db"
	"github.com/salarSb/car-sales/data/models"
	"github.com/salarSb/car-sales/pkg/logging"
	"github.com/salarSb/car-sales/pkg/service_errors"
	"gorm.io/gorm"
	"math"
	"reflect"
	"strings"
	"time"
)

type preload struct {
	string
}

type BaseService[T any, Tc any, Tu any, Tr any] struct {
	Database *gorm.DB
	Logger   logging.Logger
	Preloads []preload
}

func NewBseService[T any, Tc any, Tu any, Tr any](cfg *config.Config, preloads []preload) *BaseService[T, Tc, Tu, Tr] {
	return &BaseService[T, Tc, Tu, Tr]{
		Database: db.GetDb(),
		Logger:   logging.NewLogger(cfg),
		Preloads: preloads,
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
	bm, err := common.TypeConverter[models.BaseModel](model)
	if err != nil {
		return nil, err
	}
	return b.GetById(bm.Id)
}

func (b *BaseService[T, Tc, Tu, Tr]) Update(ctx context.Context, id int, req *Tu) (*Tr, error) {
	updateMap, err := common.TypeConverter[map[string]interface{}](req)
	snakeMap := map[string]interface{}{}
	for k, v := range *updateMap {
		snakeMap[common.ToSnakeCase(k)] = v
	}
	if err != nil {
		return nil, err
	}
	snakeMap["modified_by"] = &sql.NullInt64{Int64: int64(ctx.Value(constants.UserIdKey).(float64)), Valid: true}
	snakeMap["modified_at"] = sql.NullTime{Valid: true, Time: time.Now().UTC()}
	model := new(T)
	tx := b.Database.WithContext(ctx).Begin()
	if err = tx.
		Model(model).
		Where("id = ? AND deleted_by is null AND deleted_at is null", id).
		Updates(snakeMap).
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
	database := Preload(b.Database, b.Preloads)
	err := database.Where("id = ? AND deleted_by is null AND deleted_at is null", id).First(model).Error
	if err != nil {
		return nil, err
	}
	return common.TypeConverter[Tr](model)
}

func (b *BaseService[T, Tc, Tu, Tr]) GetByFilter(req *dto.PaginationInputWithFilter) (*dto.PagedList[Tr], error) {
	return Paginate[T, Tr](req, b.Preloads, b.Database)
}

func getQuery[T any](filter *dto.DynamicFilter) string {
	t := new(T)
	typeT := reflect.TypeOf(*t)
	query := make([]string, 0)
	query = append(query, "deleted_by is null")
	if filter.Filter != nil {
		for name, filter := range filter.Filter {
			fld, ok := typeT.FieldByName(name)
			if ok {
				fld.Name = common.ToSnakeCase(fld.Name)
				switch filter.Type {
				case "contains":
					query = append(query, fmt.Sprintf("%s ILike '%%%s%%'", fld.Name, filter.From))
				case "notContains":
					query = append(query, fmt.Sprintf("%s not ILike '%%%s%%'", fld.Name, filter.From))
				case "startsWith":
					query = append(query, fmt.Sprintf("%s ILike '%s%%'", fld.Name, filter.From))
				case "endsWith":
					query = append(query, fmt.Sprintf("%s ILike '%%%s'", fld.Name, filter.From))
				case "equals":
					query = append(query, fmt.Sprintf("%s = '%s'", fld.Name, filter.From))
				case "notEqual":
					query = append(query, fmt.Sprintf("%s != '%s'", fld.Name, filter.From))
				case "lessThan":
					query = append(query, fmt.Sprintf("%s < %s", fld.Name, filter.From))
				case "lessThanOrEqual":
					query = append(query, fmt.Sprintf("%s <= %s", fld.Name, filter.From))
				case "greaterThan":
					query = append(query, fmt.Sprintf("%s > %s", fld.Name, filter.From))
				case "greaterThanOrEqual":
					query = append(query, fmt.Sprintf("%s >= %s", fld.Name, filter.From))
				case "inRange":
					if fld.Type.Kind() == reflect.String {
						query = append(query, fmt.Sprintf("%s >= '%s'", fld.Name, filter.From))
						query = append(query, fmt.Sprintf("%s <= '%s'", fld.Name, filter.To))
					} else {
						query = append(query, fmt.Sprintf("%s >= %s", fld.Name, filter.From))
						query = append(query, fmt.Sprintf("%s <= %s", fld.Name, filter.To))
					}
				}
			}
		}
	}
	return strings.Join(query, " AND ")
}

func getSort[T any](filter *dto.DynamicFilter) string {
	t := new(T)
	typeT := reflect.TypeOf(*t)
	sort := make([]string, 0)
	if filter.Sort != nil {
		for _, tp := range *filter.Sort {
			fld, ok := typeT.FieldByName(tp.ColId)
			if ok && (tp.Sort == "asc" || tp.Sort == "desc") {
				fld.Name = common.ToSnakeCase(fld.Name)
				sort = append(sort, fmt.Sprintf("%s %s", fld.Name, tp.Sort))
			}
		}
	}
	return strings.Join(sort, ", ")
}

func Preload(db *gorm.DB, preloads []preload) *gorm.DB {
	for _, item := range preloads {
		db = db.Preload(item.string)
	}
	return db
}

func NewPagedList[T any](items *[]T, count int64, pageNumber int, pageSize int64) *dto.PagedList[T] {
	pl := &dto.PagedList[T]{
		PageNumber: pageNumber,
		TotalRows:  count,
		Items:      items,
	}
	pl.TotalPages = int(math.Ceil(float64(count) / float64(pageSize)))
	pl.HasNextPage = pl.PageNumber < pl.TotalPages
	pl.HasPreviousPage = pl.PageNumber > 1
	return pl
}

func Paginate[T any, Tr any](pagination *dto.PaginationInputWithFilter, preloads []preload, db *gorm.DB) (*dto.PagedList[Tr], error) {
	model := new(T)
	var items *[]T
	var rItems *[]Tr
	db = Preload(db, preloads)
	query := getQuery[T](&pagination.DynamicFilter)
	sort := getSort[T](&pagination.DynamicFilter)
	var totalRows int64 = 0
	db.Model(model).Where(query).Count(&totalRows)
	err := db.Where(query).Offset(pagination.GetOffset()).Limit(pagination.GetPageSize()).Order(sort).Find(&items).Error
	if err != nil {
		return nil, err
	}
	rItems, err = common.TypeConverter[[]Tr](items)
	if err != nil {
		return nil, err
	}
	return NewPagedList(rItems, totalRows, pagination.PageNumber, int64(pagination.PageSize)), nil
}
