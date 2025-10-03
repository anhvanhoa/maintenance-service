package maintenance_schedule_service

import (
	"production_service/domain/usecase/maintenance_schedule"
	"production_service/infrastructure/repo"

	proto_maintenance_schedule "github.com/anhvanhoa/sf-proto/gen/maintenance_schedule/v1"
)

type MaintenanceScheduleService struct {
	usecase maintenance_schedule.MaintenanceScheduleUsecase
	proto_maintenance_schedule.UnsafeMaintenanceScheduleServiceServer
}

func NewMaintenanceScheduleService(repos repo.Repositories) proto_maintenance_schedule.MaintenanceScheduleServiceServer {
	usecase := maintenance_schedule.NewMaintenanceScheduleUsecase(repos.MaintenanceScheduleRepo())
	return &MaintenanceScheduleService{
		usecase: usecase,
	}
}
