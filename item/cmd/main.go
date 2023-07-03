package main

import (
	"context"
	"github.com/abdullayev13/ms_item_clickhead/item/config"
	"github.com/abdullayev13/ms_item_clickhead/item/grpc"
	"github.com/abdullayev13/ms_item_clickhead/item/storage/postgres"
	_ "github.com/lib/pq"
	"log"
	"net"
)

func main() {
	cfg := config.Load()

	strg, err := postgres.NewPostgres(context.Background(), cfg)
	if err != nil {
		log.Panic("postgres.NewPostgres" + err.Error())
	}
	defer strg.CloseDB()

	grpcServer := grpc.SetUpServer(cfg, strg)

	lis, err := net.Listen("tcp", cfg.AuthGRPCPort)
	if err != nil {
		log.Panic("net.Listen: " + err.Error())
	}

	log.Printf("GRPC: Server being started port %s\n", cfg.AuthGRPCPort)

	if err = grpcServer.Serve(lis); err != nil {
		log.Panic("grpcServer.Serve: ", err.Error())
	}
}
