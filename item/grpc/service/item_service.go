package service

import (
	"context"
	"github.com/abdullayev13/ms_item_clickhead/item/config"
	"github.com/abdullayev13/ms_item_clickhead/item/genproto/item"
	"github.com/abdullayev13/ms_item_clickhead/item/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ItemService struct {
	cfg  config.Config
	strg storage.StorageI
	*item.UnimplementedItemServiceServer
}

func NewItemService(cfg config.Config, strg storage.StorageI) *ItemService {
	return &ItemService{
		cfg:  cfg,
		strg: strg,
	}
}

func (s *ItemService) Create(ctx context.Context, req *item.CreateItem) (*item.Item, error) {
	id, err := s.strg.Item().Create(ctx, req)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	user, err := s.strg.Item().GetByID(ctx, id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return user, nil
}

func (s *ItemService) GetByID(ctx context.Context, pk *item.ItemPrimaryKey) (*item.Item, error) {
	user, err := s.strg.Item().GetByID(ctx, pk)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return user, nil
}

func (s *ItemService) GetList(ctx context.Context, req *item.GetListItemRequest) (*item.GetListItemResponse, error) {
	users, err := s.strg.Item().GetAll(ctx, req)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return users, nil
}

func (s *ItemService) Update(ctx context.Context, req *item.UpdateItem) (*item.Item, error) {
	err := s.strg.Item().Update(ctx, req)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	user, err := s.strg.Item().GetByID(ctx, &item.ItemPrimaryKey{Id: req.Id})
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return user, nil
}

func (s *ItemService) Delete(ctx context.Context, pk *item.ItemPrimaryKey) (*item.MessageString, error) {
	err := s.strg.Item().Delete(ctx, pk)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return &item.MessageString{Message: "successfully deleted"}, nil
}
