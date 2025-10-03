package maintenance_schedule_service

import (
	"context"
	"production_service/domain/entity"
	"production_service/domain/usecase/maintenance_schedule"

	proto_maintenance_schedule "github.com/anhvanhoa/sf-proto/gen/maintenance_schedule/v1"
)

func (s *MaintenanceScheduleService) UpdateMaintenanceSchedule(ctx context.Context, req *proto_maintenance_schedule.UpdateMaintenanceScheduleRequest) (*proto_maintenance_schedule.UpdateMaintenanceScheduleResponse, error) {
	usecaseReq := s.convertProtoToUpdateReq(req)
	ms, err := s.usecase.UpdateMaintenanceSchedule(ctx, usecaseReq)
	if err != nil {
		return nil, err
	}
	protoMs, err := s.convertEntityToProto(ms)
	if err != nil {
		return nil, err
	}
	return &proto_maintenance_schedule.UpdateMaintenanceScheduleResponse{
		MaintenanceSchedule: protoMs,
	}, nil
}

func (s *MaintenanceScheduleService) convertProtoToUpdateReq(req *proto_maintenance_schedule.UpdateMaintenanceScheduleRequest) maintenance_schedule.UpdateMaintenanceScheduleRequest {
	ms := maintenance_schedule.UpdateMaintenanceScheduleRequest{
		ID:                      req.Id,
		DeviceID:                req.DeviceId,
		MaintenanceType:         entity.MaintenanceType(req.MaintenanceType),
		MaintenanceCategory:     entity.MaintenanceCategory(req.MaintenanceCategory),
		Priority:                entity.Priority(req.Priority),
		EstimatedDurationHours:  req.EstimatedDurationHours,
		ActualDurationHours:     req.ActualDurationHours,
		Technician:              req.Technician,
		TechnicianContact:       req.TechnicianContact,
		Cost:                    req.Cost,
		PartsReplaced:           req.PartsReplaced,
		ToolsRequired:           req.ToolsRequired,
		SafetyPrecautions:       req.SafetyPrecautions,
		PreMaintenanceReadings:  req.PreMaintenanceReadings,
		PostMaintenanceReadings: req.PostMaintenanceReadings,
		CalibrationValues:       req.CalibrationValues,
		TestResults:             req.TestResults,
		Status:                  entity.Status(req.Status),
		CompletionRating:        int(req.CompletionRating),
		MaintenanceIntervalDays: int(req.MaintenanceIntervalDays),
		WarrantyImpact:          req.WarrantyImpact,
		DowntimeMinutes:         int(req.DowntimeMinutes),
		Notes:                   req.Notes,
		MaintenanceLog:          req.MaintenanceLog,
		BeforeImages:            req.BeforeImages,
		AfterImages:             req.AfterImages,
		CreatedBy:               req.CreatedBy,
	}

	if req.ScheduledDate != nil {
		scheduledDate := req.ScheduledDate.AsTime()
		ms.ScheduledDate = &scheduledDate
	}
	if req.CompletedDate != nil {
		completedDate := req.CompletedDate.AsTime()
		ms.CompletedDate = &completedDate
	}
	if req.NextMaintenanceDate != nil {
		nextMaintenanceDate := req.NextMaintenanceDate.AsTime()
		ms.NextMaintenanceDate = &nextMaintenanceDate
	}

	return ms
}
