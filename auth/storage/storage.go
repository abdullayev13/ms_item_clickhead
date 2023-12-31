package storage

import (
	"context"
	"github.com/abdullayev13/ms_item_clickhead/auth/genproto/auth"
	"github.com/abdullayev13/ms_item_clickhead/auth/models"
)

type StorageI interface {
	CloseDB()
	User() UserRepoI
}

type UserRepoI interface {
	Create(context.Context, *auth.CreateUser) (*auth.UserPrimaryKey, error)
	GetByID(context.Context, *auth.UserPrimaryKey) (*auth.User, error)
	GetAll(context.Context, *auth.GetListUserRequest) (*auth.GetListUserResponse, error)
	Update(context.Context, *auth.UpdateUser) error
	Delete(context.Context, *auth.UserPrimaryKey) error
	GetByUsername(ctx context.Context, username string) (*models.User, error)
	ExistsByUsername(ctx context.Context, username string) (bool, error)
}
