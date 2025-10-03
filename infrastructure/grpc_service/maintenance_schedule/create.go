package maintenance_schedule_service

import (
	"context"
	"encoding/json"
	"production_service/domain/entity"
	"production_service/domain/usecase/maintenance_schedule"

	proto_maintenance_schedule "github.com/anhvanhoa/sf-proto/gen/maintenance_schedule/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *MaintenanceScheduleService) CreateMaintenanceSchedule(ctx context.Context, req *proto_maintenance_schedule.CreateMaintenanceScheduleRequest) (*proto_maintenance_schedule.CreateMaintenanceScheduleResponse, error) {
	usecaseReq := s.convertProtoToCreateReq(req)
	ms, err := s.usecase.CreateMaintenanceSchedule(ctx, usecaseReq)
	if err != nil {
		return nil, err
	}
	protoMs, err := s.convertEntityToProto(ms)
	if err != nil {
		return nil, err
	}
	return &proto_maintenance_schedule.CreateMaintenanceScheduleResponse{
		MaintenanceSchedule: protoMs,
	}, nil
}

func (s *MaintenanceScheduleService) convertProtoToCreateReq(req *proto_maintenance_schedule.CreateMaintenanceScheduleRequest) maintenance_schedule.CreateMaintenanceScheduleRequest {
	ms := maintenance_schedule.CreateMaintenanceScheduleRequest{
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

func (s *MaintenanceScheduleService) convertEntityToProto(ms *entity.MaintenanceSchedule) (*proto_maintenance_schedule.MaintenanceSchedule, error) {
	protoMs := &proto_maintenance_schedule.MaintenanceSchedule{
		Id:                      ms.ID,
		DeviceId:                ms.DeviceID,
		MaintenanceType:         string(ms.MaintenanceType),
		MaintenanceCategory:     string(ms.MaintenanceCategory),
		Priority:                string(ms.Priority),
		CompletionRating:        int32(ms.CompletionRating),
		MaintenanceIntervalDays: int32(ms.MaintenanceIntervalDays),
		WarrantyImpact:          ms.WarrantyImpact,
		DowntimeMinutes:         int32(ms.DowntimeMinutes),
		Notes:                   ms.Notes,
		MaintenanceLog:          ms.MaintenanceLog,
		BeforeImages:            ms.BeforeImages,
		AfterImages:             ms.AfterImages,
		CreatedBy:               ms.CreatedBy,
		EstimatedDurationHours:  ms.EstimatedDurationHours,
		ActualDurationHours:     ms.ActualDurationHours,
		Technician:              ms.Technician,
		TechnicianContact:       ms.TechnicianContact,
		Cost:                    ms.Cost,
		SafetyPrecautions:       ms.SafetyPrecautions,
		TestResults:             ms.TestResults,
		Status:                  string(ms.Status),
		CreatedAt:               timestamppb.New(ms.CreatedAt),
	}
	if ms.ScheduledDate != nil {
		protoMs.ScheduledDate = timestamppb.New(*ms.ScheduledDate)
	}
	if ms.CompletedDate != nil {
		protoMs.CompletedDate = timestamppb.New(*ms.CompletedDate)
	}
	if ms.NextMaintenanceDate != nil {
		protoMs.NextMaintenanceDate = timestamppb.New(*ms.NextMaintenanceDate)
	}
	if ms.PartsReplaced != nil {
		partsReplaced, err := json.Marshal(ms.PartsReplaced)
		protoMs.PartsReplaced = string(partsReplaced)
		if err != nil {
			return nil, err
		}
	}
	if ms.ToolsRequired != nil {
		toolsRequired, err := json.Marshal(ms.ToolsRequired)
		protoMs.ToolsRequired = string(toolsRequired)
		if err != nil {
			return nil, err
		}
	}
	if ms.PreMaintenanceReadings != nil {
		preMaintenanceReadings, err := json.Marshal(ms.PreMaintenanceReadings)
		protoMs.PreMaintenanceReadings = string(preMaintenanceReadings)
		if err != nil {
			return nil, err
		}
	}
	if ms.PostMaintenanceReadings != nil {
		postMaintenanceReadings, err := json.Marshal(ms.PostMaintenanceReadings)
		protoMs.PostMaintenanceReadings = string(postMaintenanceReadings)
		if err != nil {
			return nil, err
		}
	}
	if ms.CalibrationValues != nil {
		calibrationValues, err := json.Marshal(ms.CalibrationValues)
		protoMs.CalibrationValues = string(calibrationValues)
		if err != nil {
			return nil, err
		}
	}
	return protoMs, nil
}
