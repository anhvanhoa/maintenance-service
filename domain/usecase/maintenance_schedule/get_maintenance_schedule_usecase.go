package maintenance_schedule

import (
	"context"

	"production_service/domain/entity"
	"production_service/domain/repository"
)

type GetMaintenanceScheduleRequest struct {
	ID string
}

type GetMaintenanceScheduleUsecaseI interface {
	Execute(ctx context.Context, req GetMaintenanceScheduleRequest) (*entity.MaintenanceSchedule, error)
}

type GetMaintenanceScheduleUsecase struct {
	repo repository.MaintenanceScheduleRepository
}

func NewGetMaintenanceScheduleUsecase(repo repository.MaintenanceScheduleRepository) GetMaintenanceScheduleUsecaseI {
	return &GetMaintenanceScheduleUsecase{repo: repo}
}

func (u *GetMaintenanceScheduleUsecase) Execute(ctx context.Context, req GetMaintenanceScheduleRequest) (*entity.MaintenanceSchedule, error) {
	if err := u.validate(req); err != nil {
		return nil, err
	}

	m, err := u.repo.GetByID(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (u *GetMaintenanceScheduleUsecase) validate(req GetMaintenanceScheduleRequest) error {
	if req.ID == "" {
		return ErrIdIsRequired
	}
	return nil
}
