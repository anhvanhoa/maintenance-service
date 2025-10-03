package main

import (
	"context"
	"production_service/bootstrap"
	"production_service/infrastructure/grpc_service"
	maintenance_schedule_service "production_service/infrastructure/grpc_service/maintenance_schedule"

	"github.com/anhvanhoa/service-core/domain/discovery"
)

func main() {
	StartGRPCServer()
}

func StartGRPCServer() {
	app := bootstrap.App()
	env := app.Env
	log := app.Log

	discoveryConfig := &discovery.DiscoveryConfig{
		ServiceName:   env.NameService,
		ServicePort:   env.PortGrpc,
		ServiceHost:   env.HostGprc,
		IntervalCheck: env.IntervalCheck,
		TimeoutCheck:  env.TimeoutCheck,
	}

	discovery, err := discovery.NewDiscovery(discoveryConfig)
	if err != nil {
		log.Fatal("Failed to create discovery: " + err.Error())
	}
	discovery.Register()

	maintenanceScheduleService := maintenance_schedule_service.NewMaintenanceScheduleService(app.Repos)
	grpcSrv := grpc_service.NewGRPCServer(
		env, log, maintenanceScheduleService,
	)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	if err := grpcSrv.Start(ctx); err != nil {
		log.Fatal("gRPC server error: " + err.Error())
	}
}
