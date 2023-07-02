package grpc

import (
	"github.com/abdullayev13/ms_item_clickhead/auth/config"
	"github.com/abdullayev13/ms_item_clickhead/auth/genproto/auth"
	"github.com/abdullayev13/ms_item_clickhead/auth/grpc/service"
	"github.com/abdullayev13/ms_item_clickhead/auth/storage"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func SetUpServer(cfg config.Config, strg storage.StorageI) (grpcServer *grpc.Server) {

	grpcServer = grpc.NewServer()

	auth.RegisterAuthServiceServer(grpcServer, service.NewUserService(cfg, strg))

	reflection.Register(grpcServer)
	return
}
