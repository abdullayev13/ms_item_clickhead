package grpc

import (
	"github.com/abdullayev13/ms_item_clickhead/item/config"
	"github.com/abdullayev13/ms_item_clickhead/item/genproto/item"
	"github.com/abdullayev13/ms_item_clickhead/item/grpc/service"
	"github.com/abdullayev13/ms_item_clickhead/item/storage"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func SetUpServer(cfg config.Config, strg storage.StorageI) (grpcServer *grpc.Server) {

	grpcServer = grpc.NewServer()

	item.RegisterItemServiceServer(grpcServer, service.NewItemService(cfg, strg))

	reflection.Register(grpcServer)
	return
}
