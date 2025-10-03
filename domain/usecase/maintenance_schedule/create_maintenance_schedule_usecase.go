package maintenance_schedule

import (
	"context"
	"time"

	"production_service/domain/entity"
	"production_service/domain/repository"
)

// CreateMaintenanceScheduleRequest defines the input to create a maintenance schedule.
type CreateMaintenanceScheduleRequest struct {
	DeviceID                string
	MaintenanceType         entity.MaintenanceType
	MaintenanceCategory     entity.MaintenanceCategory
	Priority                entity.Priority
	ScheduledDate           *time.Time
	EstimatedDurationHours  float64
	CompletedDate           *time.Time
	ActualDurationHours     float64
	Technician              string
	TechnicianContact       string
	Cost                    float64
	PartsReplaced           string
	ToolsRequired           string
	SafetyPrecautions       string
	PreMaintenanceReadings  string
	PostMaintenanceReadings string
	CalibrationValues       string
	TestResults             string
	Status                  entity.Status
	CompletionRating        int
	NextMaintenanceDate     *time.Time
	MaintenanceIntervalDays int
	WarrantyImpact          bool
	DowntimeMinutes         int
	Notes                   string
	MaintenanceLog          string
	BeforeImages            string
	AfterImages             string
	CreatedBy               string
}

// CreateMaintenanceScheduleUsecase handles creation logic
type CreateMaintenanceScheduleUsecase struct {
	repo repository.MaintenanceScheduleRepository
}

func NewCreateMaintenanceScheduleUsecase(repo repository.MaintenanceScheduleRepository) *CreateMaintenanceScheduleUsecase {
	return &CreateMaintenanceScheduleUsecase{repo: repo}
}

func (u *CreateMaintenanceScheduleUsecase) Execute(ctx context.Context, req *CreateMaintenanceScheduleRequest) (*entity.MaintenanceSchedule, error) {
	if req == nil {
		return nil, nil
	}
	if req.DeviceID == "" {
		return nil, nil
	}

	m := &entity.MaintenanceSchedule{
		DeviceID:                req.DeviceID,
		MaintenanceType:         req.MaintenanceType,
		MaintenanceCategory:     req.MaintenanceCategory,
		Priority:                req.Priority,
		ScheduledDate:           req.ScheduledDate,
		EstimatedDurationHours:  req.EstimatedDurationHours,
		CompletedDate:           req.CompletedDate,
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
		Status:                  req.Status,
		CompletionRating:        req.CompletionRating,
		NextMaintenanceDate:     req.NextMaintenanceDate,
		MaintenanceIntervalDays: req.MaintenanceIntervalDays,
		WarrantyImpact:          req.WarrantyImpact,
		DowntimeMinutes:         req.DowntimeMinutes,
		Notes:                   req.Notes,
		MaintenanceLog:          req.MaintenanceLog,
		BeforeImages:            req.BeforeImages,
		AfterImages:             req.AfterImages,
		CreatedBy:               req.CreatedBy,
	}

	// Default status if not provided
	if m.Status == "" {
		m.Status = entity.StatusScheduled
	}

	if err := u.repo.Create(ctx, m); err != nil {
		return nil, err
	}
	return m, nil
}
