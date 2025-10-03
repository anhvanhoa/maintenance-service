package repository

import (
	"context"
	"time"

	"production_service/domain/entity"

	"github.com/anhvanhoa/service-core/common"
)

type MaintenanceScheduleRepository interface {
	Create(ctx context.Context, model *entity.MaintenanceSchedule) error
	GetByID(ctx context.Context, id string) (*entity.MaintenanceSchedule, error)
	Update(ctx context.Context, model *entity.MaintenanceSchedule) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, pagination common.Pagination, filter MaintenanceScheduleFilter) ([]*entity.MaintenanceSchedule, int, error)
}

type MaintenanceScheduleFilter struct {
	DeviceID   *string
	Statuses   []entity.Status
	Types      []entity.MaintenanceType
	Categories []entity.MaintenanceCategory
	Priorities []entity.Priority
	FromDate   *time.Time
	ToDate     *time.Time
}
