package postgres

import (
	"context"
	"github.com/abdullayev13/ms_item_clickhead/item/config"
	"github.com/abdullayev13/ms_item_clickhead/item/storage"
	"github.com/iancoleman/strcase"
	"github.com/jmoiron/sqlx"

	"fmt"
)

type Store struct {
	db   *sqlx.DB
	item storage.ItemRepoI
}

func NewPostgres(ctx context.Context, cfg config.Config) (storage.StorageI, error) {
	dsn := fmt.Sprintf("user=%s dbname=%s password=%s port=%d sslmode=disable",
		cfg.PostgresUser, cfg.PostgresDatabase, cfg.PostgresPassword, cfg.PostgresPort)

	db, err := sqlx.ConnectContext(ctx, "postgres", dsn)
	if err != nil {
		return nil, err
	}

	db.MapperFunc(func(s string) string {
		return strcase.ToSnake(s)
	})

	if err != nil {
		return nil, err
	}

	return &Store{
		db:   db,
		item: NewItemRepo(db),
	}, nil
}

func (s *Store) CloseDB() {
	s.db.Close()
}

func (s *Store) Item() storage.ItemRepoI {
	return s.item
}
