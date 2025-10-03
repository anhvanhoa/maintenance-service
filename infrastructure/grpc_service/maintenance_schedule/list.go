package maintenance_schedule_service

import (
	"context"
	"production_service/domain/entity"
	"production_service/domain/repository"
	"production_service/domain/usecase/maintenance_schedule"

	"github.com/anhvanhoa/service-core/common"
	proto_common "github.com/anhvanhoa/sf-proto/gen/common/v1"
	proto_maintenance_schedule "github.com/anhvanhoa/sf-proto/gen/maintenance_schedule/v1"
)

func (s *MaintenanceScheduleService) ListMaintenanceSchedule(ctx context.Context, req *proto_maintenance_schedule.ListMaintenanceScheduleRequest) (*proto_maintenance_schedule.ListMaintenanceScheduleResponse, error) {
	pagination := common.Pagination{
		Page:     int(req.Pagination.Page),
		PageSize: int(req.Pagination.PageSize),
	}

	filter := s.convertProtoToRepositoryFilter(req.Filter)

	usecaseReq := maintenance_schedule.ListMaintenanceScheduleRequest{
		Pagination: pagination,
		Filter:     filter,
	}

	result, err := s.usecase.ListMaintenanceSchedule(ctx, usecaseReq)
	if err != nil {
		return nil, err
	}

	var protoSchedules []*proto_maintenance_schedule.MaintenanceSchedule
	for _, schedule := range result.Data {
		protoSchedule, err := s.convertEntityToProto(schedule)
		if err != nil {
			return nil, err
		}
		protoSchedules = append(protoSchedules, protoSchedule)
	}

	return &proto_maintenance_schedule.ListMaintenanceScheduleResponse{
		MaintenanceSchedules: protoSchedules,
		Pagination: &proto_common.PaginationResponse{
			Total:      int32(result.Total),
			TotalPages: int32(result.TotalPages),
			PageSize:   int32(result.PageSize),
			Page:       int32(result.Page),
		},
	}, nil
}

func (s *MaintenanceScheduleService) convertProtoToRepositoryFilter(filter *proto_maintenance_schedule.MaintenanceScheduleFilter) repository.MaintenanceScheduleFilter {
	if filter == nil {
		return repository.MaintenanceScheduleFilter{}
	}

	repoFilter := repository.MaintenanceScheduleFilter{}

	if len(filter.Statuses) > 0 {
		var statuses []entity.Status
		for _, status := range filter.Statuses {
			statuses = append(statuses, entity.Status(status))
		}
		repoFilter.Statuses = statuses
	}

	if len(filter.Types) > 0 {
		var types []entity.MaintenanceType
		for _, t := range filter.Types {
			types = append(types, entity.MaintenanceType(t))
		}
		repoFilter.Types = types
	}

	if len(filter.Categories) > 0 {
		var categories []entity.MaintenanceCategory
		for _, category := range filter.Categories {
			categories = append(categories, entity.MaintenanceCategory(category))
		}
		repoFilter.Categories = categories
	}

	if len(filter.Priorities) > 0 {
		var priorities []entity.Priority
		for _, priority := range filter.Priorities {
			priorities = append(priorities, entity.Priority(priority))
		}
		repoFilter.Priorities = priorities
	}

	if filter.FromDate != nil {
		fromDate := filter.FromDate.AsTime()
		repoFilter.FromDate = &fromDate
	}
	if filter.ToDate != nil {
		toDate := filter.ToDate.AsTime()
		repoFilter.ToDate = &toDate
	}

	return repoFilter
}
