package repo

import (
	"context"
	"errors"
	"fmt"

	"github.com/anhvanhoa/service-core/common"
	"github.com/anhvanhoa/service-core/utils"
	"github.com/go-pg/pg/v10"

	"production_service/domain/entity"
	"production_service/domain/repository"
)

type maintenanceScheduleRepository struct {
	db     *pg.DB
	helper utils.Helper
}

func NewMaintenanceScheduleRepository(db *pg.DB, helper utils.Helper) repository.MaintenanceScheduleRepository {
	return &maintenanceScheduleRepository{db: db, helper: helper}
}

func (r *maintenanceScheduleRepository) Create(ctx context.Context, model *entity.MaintenanceSchedule) error {
	_, err := r.db.WithContext(ctx).Model(model).Insert()
	return err
}

func (r *maintenanceScheduleRepository) GetByID(ctx context.Context, id string) (*entity.MaintenanceSchedule, error) {
	var m entity.MaintenanceSchedule
	err := r.db.WithContext(ctx).Model(&m).Where("id = ?", id).Limit(1).Select()
	if errors.Is(err, pg.ErrNoRows) {
		return nil, nil
	}
	return &m, err
}

func (r *maintenanceScheduleRepository) Update(ctx context.Context, model *entity.MaintenanceSchedule) error {
	_, err := r.db.WithContext(ctx).Model(model).Where("id = ?", model.ID).UpdateNotZero()
	return err
}

func (r *maintenanceScheduleRepository) Delete(ctx context.Context, id string) error {
	_, err := r.db.WithContext(ctx).Model((*entity.MaintenanceSchedule)(nil)).Where("id = ?", id).Delete()
	return err
}

func (r *maintenanceScheduleRepository) List(ctx context.Context, pagination common.Pagination, filter repository.MaintenanceScheduleFilter) ([]*entity.MaintenanceSchedule, int, error) {
	var results []*entity.MaintenanceSchedule
	q := r.db.WithContext(ctx).Model(&results)

	if filter.DeviceID != nil && *filter.DeviceID != "" {
		q = q.Where("device_id = ?", *filter.DeviceID)
	}
	if len(filter.Statuses) > 0 {
		q = q.Where("status IN (?)", pg.In(filter.Statuses))
	}
	if len(filter.Types) > 0 {
		q = q.Where("maintenance_type IN (?)", pg.In(filter.Types))
	}
	if len(filter.Categories) > 0 {
		q = q.Where("maintenance_category IN (?)", pg.In(filter.Categories))
	}
	if len(filter.Priorities) > 0 {
		q = q.Where("priority IN (?)", pg.In(filter.Priorities))
	}
	if filter.FromDate != nil && *filter.FromDate != "" {
		q = q.Where("scheduled_date >= ?", *filter.FromDate)
	}
	if filter.ToDate != nil && *filter.ToDate != "" {
		q = q.Where("scheduled_date <= ?", *filter.ToDate)
	}

	sortBy := "scheduled_date"
	if pagination.SortBy != "" {
		sortBy = pagination.SortBy
	}
	sortDirection := "ASC"
	if pagination.SortOrder == "DESC" {
		sortDirection = "DESC"
	}
	q = q.Order(fmt.Sprintf("%s %s", sortBy, sortDirection))

	// Pagination
	page := pagination.Page
	limit := pagination.PageSize
	if limit <= 0 || limit > 1000 {
		limit = 50
	}
	if page <= 0 {
		page = 1
	}
	offset := r.helper.CalculateOffset(page, limit)
	q = q.Offset(offset).Limit(limit)

	count, err := q.SelectAndCount()
	if err != nil {
		return nil, 0, err
	}
	return results, count, nil
}
