package client

import (
	"github.com/abdullayev13/ms_item_clickhead/api_gateway/config"
	"github.com/abdullayev13/ms_item_clickhead/api_gateway/genproto/auth"
	"github.com/abdullayev13/ms_item_clickhead/api_gateway/genproto/item"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceManagerI interface {
	AuthService() auth.AuthServiceClient
	ItemService() item.ItemServiceClient
}

type grpcClients struct {
	authService auth.AuthServiceClient
	itemService item.ItemServiceClient
}

func NewGrpcClients(cfg config.Config) (ServiceManagerI, error) {
	//connecting to user service
	connUserService, err := grpc.Dial(
		cfg.AuthServiceHost+cfg.AuthServicePort,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	//connecting to article service
	connArticleService, err := grpc.Dial(
		cfg.ArticleServiceHost+cfg.ArticleServicePort,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	return &grpcClients{
		authService: auth.NewAuthServiceClient(connUserService),
		itemService: item.NewItemServiceClient(connArticleService),
	}, nil
}

func (g *grpcClients) AuthService() auth.AuthServiceClient {
	return g.authService
}
func (g *grpcClients) ItemService() item.ItemServiceClient {
	return g.itemService
}
