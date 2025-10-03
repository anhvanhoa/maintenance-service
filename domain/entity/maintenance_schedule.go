package entity

import (
	"time"
)

type MaintenanceSchedule struct {
	tableName struct{} `pg:"maintenance_schedules"`

	ID                      string
	DeviceID                string
	MaintenanceType         MaintenanceType
	MaintenanceCategory     MaintenanceCategory
	Priority                Priority
	ScheduledDate           *time.Time
	EstimatedDurationHours  float64
	CompletedDate           *time.Time
	ActualDurationHours     float64
	Technician              string
	TechnicianContact       string
	Cost                    float64
	PartsReplaced           []any
	ToolsRequired           []any
	SafetyPrecautions       string
	PreMaintenanceReadings  map[string]any
	PostMaintenanceReadings map[string]any
	CalibrationValues       map[string]any
	TestResults             string
	Status                  Status
	CompletionRating        int
	NextMaintenanceDate     *time.Time
	MaintenanceIntervalDays int
	WarrantyImpact          bool
	DowntimeMinutes         int
	Notes                   string
	MaintenanceLog          string
	BeforeImages            []string `pg:",array"`
	AfterImages             []string `pg:",array"`
	CreatedBy               string
	CreatedAt               time.Time
	UpdatedAt               *time.Time
}

func (m *MaintenanceSchedule) TableName() any {
	return m.tableName
}

type MaintenanceType string

const (
	MaintenanceTypeCleaning       MaintenanceType = "cleaning"
	MaintenanceTypeCalibration    MaintenanceType = "calibration"
	MaintenanceTypeReplacement    MaintenanceType = "replacement"
	MaintenanceTypeRepair         MaintenanceType = "repair"
	MaintenanceTypeInspection     MaintenanceType = "inspection"
	MaintenanceTypeSoftwareUpdate MaintenanceType = "software_update"
)

type MaintenanceCategory string

const (
	MaintenanceCategoryPreventive MaintenanceCategory = "preventive"
	MaintenanceCategoryCorrective MaintenanceCategory = "corrective"
	MaintenanceCategoryEmergency  MaintenanceCategory = "emergency"
	MaintenanceCategoryRoutine    MaintenanceCategory = "routine"
)

type Priority string

const (
	PriorityLow      Priority = "low"
	PriorityMedium   Priority = "medium"
	PriorityHigh     Priority = "high"
	PriorityCritical Priority = "critical"
)

type Status string

const (
	StatusScheduled  Status = "scheduled"
	StatusInProgress Status = "in_progress"
	StatusCompleted  Status = "completed"
	StatusCancelled  Status = "cancelled"
	StatusPostponed  Status = "postponed"
)
