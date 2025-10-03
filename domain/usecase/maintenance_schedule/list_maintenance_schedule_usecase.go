package maintenance_schedule

import (
	"context"

	"production_service/domain/entity"
	"production_service/domain/repository"

	"github.com/anhvanhoa/service-core/common"
	"github.com/anhvanhoa/service-core/utils"
)

type ListMaintenanceScheduleRequest struct {
	Pagination common.Pagination
	Filter     repository.MaintenanceScheduleFilter
}

type ListMaintenanceScheduleResponse = common.PaginationResult[*entity.MaintenanceSchedule]

type ListMaintenanceScheduleUsecaseI interface {
	Execute(ctx context.Context, req ListMaintenanceScheduleRequest) (*ListMaintenanceScheduleResponse, error)
}

type ListMaintenanceScheduleUsecase struct {
	repo   repository.MaintenanceScheduleRepository
	helper utils.Helper
}

func NewListMaintenanceScheduleUsecase(repo repository.MaintenanceScheduleRepository, helper utils.Helper) ListMaintenanceScheduleUsecaseI {
	return &ListMaintenanceScheduleUsecase{repo: repo, helper: helper}
}

func (u *ListMaintenanceScheduleUsecase) Execute(ctx context.Context, req ListMaintenanceScheduleRequest) (*ListMaintenanceScheduleResponse, error) {
	repoFilter := repository.MaintenanceScheduleFilter{
		DeviceID:   req.Filter.DeviceID,
		Statuses:   req.Filter.Statuses,
		Types:      req.Filter.Types,
		Categories: req.Filter.Categories,
		Priorities: req.Filter.Priorities,
		FromDate:   req.Filter.FromDate,
		ToDate:     req.Filter.ToDate,
	}

	data, total, err := u.repo.List(ctx, req.Pagination, repoFilter)
	if err != nil {
		return nil, err
	}
	totalPages := u.helper.CalculateTotalPages(int64(total), int64(req.Pagination.PageSize))
	return &ListMaintenanceScheduleResponse{
		Data:       data,
		Total:      int64(total),
		TotalPages: totalPages,
		PageSize:   req.Pagination.PageSize,
		Page:       req.Pagination.Page,
	}, nil
}
