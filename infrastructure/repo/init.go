package repo

import (
	"production_service/domain/repository"

	"github.com/anhvanhoa/service-core/utils"
	"github.com/go-pg/pg/v10"
)

type Repositories interface {
	MaintenanceScheduleRepo() repository.MaintenanceScheduleRepository
}

type RepositoriesImpl struct {
	MaintenanceScheduleRepository repository.MaintenanceScheduleRepository
}

func InitRepositories(db *pg.DB, helper utils.Helper) Repositories {
	return &RepositoriesImpl{
		MaintenanceScheduleRepository: NewMaintenanceScheduleRepository(db, helper),
	}
}

func (r *RepositoriesImpl) MaintenanceScheduleRepo() repository.MaintenanceScheduleRepository {
	return r.MaintenanceScheduleRepository
}
