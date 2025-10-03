package maintenance_schedule_service

import (
	"context"
	"production_service/domain/usecase/maintenance_schedule"

	proto_maintenance_schedule "github.com/anhvanhoa/sf-proto/gen/maintenance_schedule/v1"
)

func (s *MaintenanceScheduleService) GetMaintenanceSchedule(ctx context.Context, req *proto_maintenance_schedule.GetMaintenanceScheduleRequest) (*proto_maintenance_schedule.GetMaintenanceScheduleResponse, error) {
	usecaseReq := maintenance_schedule.GetMaintenanceScheduleRequest{
		ID: req.Id,
	}

	ms, err := s.usecase.GetMaintenanceSchedule(ctx, usecaseReq)
	if err != nil {
		return nil, err
	}

	protoMs, err := s.convertEntityToProto(ms)
	if err != nil {
		return nil, err
	}
	return &proto_maintenance_schedule.GetMaintenanceScheduleResponse{
		MaintenanceSchedule: protoMs,
	}, nil
}
