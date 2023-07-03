package main

import (
	"fmt"
	"github.com/abdullayev13/ms_item_clickhead/api_gateway/api"
	"github.com/abdullayev13/ms_item_clickhead/api_gateway/api/handlers"
	"github.com/abdullayev13/ms_item_clickhead/api_gateway/config"
	"github.com/abdullayev13/ms_item_clickhead/api_gateway/grpc/client"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	fmt.Printf("config: %+v\n", cfg)

	grpcSvcs, err := client.NewGrpcClients(cfg)
	if err != nil {
		panic(err)
	}

	r := gin.New()

	r.Use(gin.Logger(), gin.Recovery())

	h := handlers.NewHandler(cfg, grpcSvcs)

	api.SetUpAPI(r.Group("api"), h, cfg)

	fmt.Println("Start api gateway....")

	if err := r.Run(cfg.ServicePort); err != nil {
		return
	}
}
