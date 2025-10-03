package maintenance_schedule_service

import (
	"context"
	"production_service/domain/usecase/maintenance_schedule"

	proto_maintenance_schedule "github.com/anhvanhoa/sf-proto/gen/maintenance_schedule/v1"
)

func (s *MaintenanceScheduleService) DeleteMaintenanceSchedule(ctx context.Context, req *proto_maintenance_schedule.DeleteMaintenanceScheduleRequest) (*proto_maintenance_schedule.DeleteMaintenanceScheduleResponse, error) {
	usecaseReq := maintenance_schedule.DeleteMaintenanceScheduleRequest{
		ID: req.Id,
	}

	err := s.usecase.DeleteMaintenanceSchedule(ctx, usecaseReq)
	if err != nil {
		return nil, err
	}

	return &proto_maintenance_schedule.DeleteMaintenanceScheduleResponse{
		Message: "Xóa lịch bảo trì thành công",
		Success: true,
	}, nil
}
