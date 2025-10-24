package main

import (
	"context"
	"production_service/bootstrap"
	"production_service/infrastructure/grpc_client"
	"production_service/infrastructure/grpc_service"
	maintenance_schedule_service "production_service/infrastructure/grpc_service/maintenance_schedule"

	gc "github.com/anhvanhoa/service-core/domain/grpc_client"
)

func main() {
	StartGRPCServer()
}

func StartGRPCServer() {
	app := bootstrap.App()
	env := app.Env
	log := app.Log

	clientFactory := gc.NewClientFactory(env.GrpcClients...)
	permissionClient := grpc_client.NewPermissionClient(clientFactory.GetClient(env.PermissionServiceAddr))

	maintenanceScheduleService := maintenance_schedule_service.NewMaintenanceScheduleService(app.Repos)
	grpcSrv := grpc_service.NewGRPCServer(
		env, log,
		app.Cache,
		maintenanceScheduleService,
	)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	permissions := app.Helper.ConvertResourcesToPermissions(grpcSrv.GetResources())
	if _, err := permissionClient.PermissionServiceClient.RegisterPermission(ctx, permissions); err != nil {
		log.Fatal("Failed to register permission: " + err.Error())
	}
	if err := grpcSrv.Start(ctx); err != nil {
		log.Fatal("gRPC server error: " + err.Error())
	}
}
