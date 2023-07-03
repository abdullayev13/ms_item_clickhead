package storage

import (
	"context"
	"github.com/abdullayev13/ms_item_clickhead/item/genproto/item"
)

type StorageI interface {
	CloseDB()
	Item() ItemRepoI
}

type ItemRepoI interface {
	Create(ctx context.Context, req *item.CreateItem) (*item.ItemPrimaryKey, error)
	GetByID(ctx context.Context, pk *item.ItemPrimaryKey) (*item.Item, error)
	GetAll(ctx context.Context, req *item.GetListItemRequest) (*item.GetListItemResponse, error)
	Update(ctx context.Context, req *item.UpdateItem) error
	Delete(ctx context.Context, pk *item.ItemPrimaryKey) error
}
