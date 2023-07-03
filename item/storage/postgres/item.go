package postgres

import (
	"context"
	"errors"
	"github.com/abdullayev13/ms_item_clickhead/item/genproto/item"
	"github.com/jmoiron/sqlx"
)

type itemRepo struct {
	db *sqlx.DB
}

func NewItemRepo(db *sqlx.DB) *itemRepo {
	return &itemRepo{
		db: db,
	}
}

func (r *itemRepo) Create(ctx context.Context, req *item.CreateItem) (*item.ItemPrimaryKey, error) {
	query := `
				INSERT INTO items (
				    name, 
				    info,
					price
				)
				VALUES ($1, $2, $3)
				RETURNING id`

	res, err := r.db.QueryContext(ctx, query, req.Name, req.Info, req.Price)
	if err != nil {
		return nil, err
	}

	if !res.Next() {
		return nil, errors.New("sql no rows on creating item")
	}

	pk := &item.ItemPrimaryKey{}

	err = res.Scan(&pk.Id)
	if err != nil {
		return nil, err
	}

	return pk, nil
}

func (r *itemRepo) GetByID(ctx context.Context, pk *item.ItemPrimaryKey) (*item.Item, error) {
	query := `SELECT id, name, info, price FROM items WHERE id = $1`
	res := &item.Item{}
	err := r.db.GetContext(ctx, res, query, pk.Id)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *itemRepo) GetAll(ctx context.Context, req *item.GetListItemRequest) (*item.GetListItemResponse, error) {
	switch req.Order {
	case "name":
	case "info":
	case "price":
	default:
		req.Order = "id"
	}

	query := `SELECT id, name, info, price FROM items ORDER BY $1 LIMIT $2 OFFSET $3`
	items := make([]*item.Item, 0)

	err := r.db.SelectContext(ctx, &items, query, req.Order, req.Limit, req.Offset)
	if err != nil {
		return nil, err
	}

	res := &item.GetListItemResponse{Items: items}

	err = r.db.GetContext(ctx, &res.Count,
		`SELECT count(*) FROM items`)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *itemRepo) Update(ctx context.Context, req *item.UpdateItem) error {
	query := `
				UPDATE items SET
				name=$1,
				info=$2,
				price=$3
				WHERE id=$4`

	_, err := r.db.QueryContext(ctx, query, req.Name, req.Info, req.Price, req.Id)
	if err != nil {
		return err
	}

	return nil
}

func (r *itemRepo) Delete(ctx context.Context, pk *item.ItemPrimaryKey) error {
	query := `DELETE FROM items WHERE id=$1`

	_, err := r.db.ExecContext(ctx, query, pk.Id)
	if err != nil {
		return err
	}

	return nil
}
