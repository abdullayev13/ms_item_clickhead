package postgres

import (
	"context"
	"errors"
	"github.com/abdullayev13/ms_item_clickhead/auth/genproto/auth"
	"github.com/jmoiron/sqlx"
)

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) *userRepo {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) Create(ctx context.Context, user *auth.CreateUser) (*auth.UserPrimaryKey, error) {
	query := `
			INSERT INTO users (
			name,
			username,
			password,
			role
			)
			VALUES ($1, $2, $3, $4)
			RETURNING id`

	res, err := r.db.QueryContext(ctx, query, user.Name, user.Username, user.Password, user.Role)
	if err != nil {
		return nil, err
	}

	if !res.Next() {
		return nil, errors.New("sql no rows on creating user")
	}

	pk := &auth.UserPrimaryKey{}

	err = res.Scan(&pk.Id)
	if err != nil {
		return nil, err
	}

	return pk, nil
}

func (r *userRepo) GetByID(ctx context.Context, pk *auth.UserPrimaryKey) (*auth.User, error) {
	query := `SELECT id, name, username, role FROM users WHERE id = $1`
	user := &auth.User{}
	err := r.db.GetContext(ctx, user, query, pk.Id)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepo) GetAll(ctx context.Context, req *auth.GetListUserRequest) (*auth.GetListUserResponse, error) {
	switch req.Order {
	case "name":
	case "username":
	case "role":
	default:
		req.Order = "id"
	}

	query := `SELECT id, name, username, role FROM users ORDER BY $1 LIMIT $2 OFFSET $3`
	users := make([]*auth.User, 0)

	err := r.db.SelectContext(ctx, &users, query, req.Order, req.Limit, req.Offset)
	if err != nil {
		return nil, err
	}

	res := &auth.GetListUserResponse{Users: users}

	err = r.db.GetContext(ctx, &res.Count,
		`SELECT count(*) FROM users`)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *userRepo) Update(ctx context.Context, user *auth.UpdateUser) error {
	query := `
			UPDATE users SET
			name=$1,
			username=$2,
			password=$3,
			role=$4
			WHERE id=$5`

	_, err := r.db.QueryContext(ctx, query, user.Name, user.Username, user.Password, user.Role, user.Id)
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepo) Delete(ctx context.Context, pk *auth.UserPrimaryKey) error {
	query := `DELETE FROM users WHERE id=$1`

	_, err := r.db.ExecContext(ctx, query, pk.Id)
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepo) GetPasswordByID(ctx context.Context, pk *auth.UserPrimaryKey) (string, error) {
	query := `SELECT password FROM users WHERE id = $1`
	var password string
	err := r.db.GetContext(ctx, &password, query, pk.Id)

	if err != nil {
		return "", err
	}

	return password, nil
}

func (r *userRepo) ExistsByUsername(ctx context.Context, username string) (bool, error) {
	query := `SELECT count(id) FROM users WHERE id = $1`
	count := 0
	err := r.db.GetContext(ctx, &count, query, username)

	if err != nil {
		return false, err
	}

	return count > 0, nil
}
