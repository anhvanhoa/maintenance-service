package maintenance_schedule

import (
	"context"
	"production_service/domain/entity"
	"production_service/domain/repository"

	"github.com/anhvanhoa/service-core/utils"
)

type MaintenanceScheduleUsecase interface {
	CreateMaintenanceSchedule(ctx context.Context, req CreateMaintenanceScheduleRequest) (*entity.MaintenanceSchedule, error)
	GetMaintenanceSchedule(ctx context.Context, req GetMaintenanceScheduleRequest) (*entity.MaintenanceSchedule, error)
	UpdateMaintenanceSchedule(ctx context.Context, req UpdateMaintenanceScheduleRequest) (*entity.MaintenanceSchedule, error)
	DeleteMaintenanceSchedule(ctx context.Context, req DeleteMaintenanceScheduleRequest) error
	ListMaintenanceSchedule(ctx context.Context, req ListMaintenanceScheduleRequest) (*ListMaintenanceScheduleResponse, error)
}

type MaintenanceScheduleUsecaseImpl struct {
	createMaintenanceScheduleUsecase CreateMaintenanceScheduleUsecaseI
	getMaintenanceScheduleUsecase    GetMaintenanceScheduleUsecaseI
	updateMaintenanceScheduleUsecase UpdateMaintenanceScheduleUsecaseI
	deleteMaintenanceScheduleUsecase DeleteMaintenanceScheduleUsecaseI
	listMaintenanceScheduleUsecase   ListMaintenanceScheduleUsecaseI
}

func NewMaintenanceScheduleUsecase(repo repository.MaintenanceScheduleRepository) MaintenanceScheduleUsecase {
	helper := utils.NewHelper()
	createMaintenanceScheduleUsecase := NewCreateMaintenanceScheduleUsecase(repo)
	getMaintenanceScheduleUsecase := NewGetMaintenanceScheduleUsecase(repo)
	updateMaintenanceScheduleUsecase := NewUpdateMaintenanceScheduleUsecase(repo)
	deleteMaintenanceScheduleUsecase := NewDeleteMaintenanceScheduleUsecase(repo)
	listMaintenanceScheduleUsecase := NewListMaintenanceScheduleUsecase(repo, helper)
	return &MaintenanceScheduleUsecaseImpl{
		createMaintenanceScheduleUsecase: createMaintenanceScheduleUsecase,
		getMaintenanceScheduleUsecase:    getMaintenanceScheduleUsecase,
		updateMaintenanceScheduleUsecase: updateMaintenanceScheduleUsecase,
		deleteMaintenanceScheduleUsecase: deleteMaintenanceScheduleUsecase,
		listMaintenanceScheduleUsecase:   listMaintenanceScheduleUsecase,
	}
}

func (u *MaintenanceScheduleUsecaseImpl) CreateMaintenanceSchedule(ctx context.Context, req CreateMaintenanceScheduleRequest) (*entity.MaintenanceSchedule, error) {
	return u.createMaintenanceScheduleUsecase.Execute(ctx, req)
}

func (u *MaintenanceScheduleUsecaseImpl) GetMaintenanceSchedule(ctx context.Context, req GetMaintenanceScheduleRequest) (*entity.MaintenanceSchedule, error) {
	return u.getMaintenanceScheduleUsecase.Execute(ctx, req)
}

func (u *MaintenanceScheduleUsecaseImpl) UpdateMaintenanceSchedule(ctx context.Context, req UpdateMaintenanceScheduleRequest) (*entity.MaintenanceSchedule, error) {
	return u.updateMaintenanceScheduleUsecase.Execute(ctx, req)
}

func (u *MaintenanceScheduleUsecaseImpl) DeleteMaintenanceSchedule(ctx context.Context, req DeleteMaintenanceScheduleRequest) error {
	return u.deleteMaintenanceScheduleUsecase.Execute(ctx, req)
}

func (u *MaintenanceScheduleUsecaseImpl) ListMaintenanceSchedule(ctx context.Context, req ListMaintenanceScheduleRequest) (*ListMaintenanceScheduleResponse, error) {
	return u.listMaintenanceScheduleUsecase.Execute(ctx, req)
}
