package grpc_service

import (
	"production_service/bootstrap"

	grpc_server "github.com/anhvanhoa/service-core/bootstrap/grpc"
	"github.com/anhvanhoa/service-core/domain/log"
	"google.golang.org/grpc"
)

func NewGRPCServer(
	env *bootstrap.Env,
	log *log.LogGRPCImpl,
) *grpc_server.GRPCServer {
	config := &grpc_server.GRPCServerConfig{
		IsProduction: env.IsProduction(),
		PortGRPC:     env.PortGrpc,
		NameService:  env.NameService,
	}
	return grpc_server.NewGRPCServer(
		config,
		log,
		func(server *grpc.Server) {
		},
	)
}
