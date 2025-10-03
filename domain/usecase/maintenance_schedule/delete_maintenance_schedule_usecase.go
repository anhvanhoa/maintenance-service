package maintenance_schedule

import (
	"context"

	"production_service/domain/repository"
)

type DeleteMaintenanceScheduleRequest struct {
	ID string
}

type DeleteMaintenanceScheduleUsecaseI interface {
	Execute(ctx context.Context, req DeleteMaintenanceScheduleRequest) error
}

type DeleteMaintenanceScheduleUsecase struct {
	repo repository.MaintenanceScheduleRepository
}

func NewDeleteMaintenanceScheduleUsecase(repo repository.MaintenanceScheduleRepository) DeleteMaintenanceScheduleUsecaseI {
	return &DeleteMaintenanceScheduleUsecase{repo: repo}
}

func (u *DeleteMaintenanceScheduleUsecase) Execute(ctx context.Context, req DeleteMaintenanceScheduleRequest) error {
	if err := u.validate(req); err != nil {
		return err
	}
	if err := u.repo.Delete(ctx, req.ID); err != nil {
		return ErrDeleteFailed
	}
	return nil
}

func (u *DeleteMaintenanceScheduleUsecase) validate(req DeleteMaintenanceScheduleRequest) error {
	if req.ID == "" {
		return ErrIdIsRequired
	}
	return nil
}
