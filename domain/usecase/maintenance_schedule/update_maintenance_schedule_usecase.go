package maintenance_schedule

import (
	"context"
	"encoding/json"
	"time"

	"production_service/domain/entity"
	"production_service/domain/repository"
)

type UpdateMaintenanceScheduleRequest struct {
	ID                      string
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
	BeforeImages            []string
	AfterImages             []string
	CreatedBy               string
}

type UpdateMaintenanceScheduleUsecaseI interface {
	Execute(ctx context.Context, req UpdateMaintenanceScheduleRequest) (*entity.MaintenanceSchedule, error)
}

type UpdateMaintenanceScheduleUsecase struct {
	repo repository.MaintenanceScheduleRepository
}

func NewUpdateMaintenanceScheduleUsecase(repo repository.MaintenanceScheduleRepository) UpdateMaintenanceScheduleUsecaseI {
	return &UpdateMaintenanceScheduleUsecase{repo: repo}
}

func (u *UpdateMaintenanceScheduleUsecase) Execute(ctx context.Context, req UpdateMaintenanceScheduleRequest) (*entity.MaintenanceSchedule, error) {
	if err := u.validate(req); err != nil {
		return nil, err
	}
	_, err := u.repo.GetByID(ctx, req.ID)
	if err != nil {
		return nil, ErrNotFound
	}

	existing := &entity.MaintenanceSchedule{
		ID:                      req.ID,
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
		SafetyPrecautions:       req.SafetyPrecautions,
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

	if req.PartsReplaced != "" {
		err = json.Unmarshal([]byte(req.PartsReplaced), &existing.PartsReplaced)
		if err != nil {
			return nil, ErrUnmarshalFailed
		}
	}
	if req.ToolsRequired != "" {
		err = json.Unmarshal([]byte(req.ToolsRequired), &existing.ToolsRequired)
		if err != nil {
			return nil, ErrUnmarshalFailed
		}
	}
	if req.PreMaintenanceReadings != "" {
		err = json.Unmarshal([]byte(req.PreMaintenanceReadings), &existing.PreMaintenanceReadings)
		if err != nil {
			return nil, ErrUnmarshalFailed
		}
	}
	if req.PostMaintenanceReadings != "" {
		err = json.Unmarshal([]byte(req.PostMaintenanceReadings), &existing.PostMaintenanceReadings)
		if err != nil {
			return nil, ErrUnmarshalFailed
		}
	}
	if req.CalibrationValues != "" {
		err = json.Unmarshal([]byte(req.CalibrationValues), &existing.CalibrationValues)
		if err != nil {
			return nil, ErrUnmarshalFailed
		}
	}

	if err := u.repo.Update(ctx, existing); err != nil {
		return nil, ErrUpdateFailed
	}

	return existing, nil
}

func (u *UpdateMaintenanceScheduleUsecase) validate(req UpdateMaintenanceScheduleRequest) error {
	if req.ID == "" {
		return ErrIdIsRequired
	}
	return nil
}
