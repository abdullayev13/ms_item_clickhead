package service

import (
	"context"
	"github.com/abdullayev13/ms_item_clickhead/auth/config"
	"github.com/abdullayev13/ms_item_clickhead/auth/genproto/auth"
	"github.com/abdullayev13/ms_item_clickhead/auth/pkg/helper"
	"github.com/abdullayev13/ms_item_clickhead/auth/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strconv"
	"strings"
)

type AuthService struct {
	cfg  config.Config
	strg storage.StorageI
	*auth.UnimplementedAuthServiceServer
}

func NewUserService(cfg config.Config, strg storage.StorageI) *AuthService {
	return &AuthService{
		cfg:  cfg,
		strg: strg,
	}
}

func (s *AuthService) CheckUri(ctx context.Context, req *auth.CheckUriRequest) (*auth.CheckUriResponse, error) {
	id, err := helper.ParseAccessToken(req.Token, s.cfg.TokenSecretKey)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	user, err := s.strg.User().GetByID(ctx, &auth.UserPrimaryKey{Id: id})
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	ok := false
	{
		if user.Role == "admin" {
			ok = true
		} else {
			if strings.HasPrefix(req.Uri, "api/auth/user-me") {
				ok = true
			} else if strings.HasPrefix(req.Uri, "api/auth/product/item") {
				suffix := req.Uri[len("api/auth/product/item"):]
				_, err = strconv.Atoi(suffix)
				isNum := err == nil
				if isNum || strings.HasPrefix(suffix, "list") {
					ok = true
				}
			}
		}
	}

	res := &auth.CheckUriResponse{Ok: ok, UserId: id}
	return res, nil
}

func (s *AuthService) Login(ctx context.Context, req *auth.UserLoginRequest) (*auth.UserLoginResponse, error) {
	user, err := s.strg.User().GetByUsername(ctx, req.Username)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}

	if user.Password != req.Password {
		return nil, status.Errorf(codes.InvalidArgument, "username or password wrong")
	}

	token, err := helper.GenerateToken(user.Id, s.cfg.TokenSecretKey, s.cfg.AccessTokenExpiring, s.cfg.RefreshTokenExpiring)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := new(auth.UserLoginResponse)
	{
		res.AccessToken = token.AccessToken
		res.RefreshToken = token.RefreshToken
		res.User = &auth.User{
			Id:       user.Id,
			Name:     user.Name,
			Username: user.Username,
			Role:     user.Role,
		}
	}

	return res, nil
}

func (s *AuthService) Create(ctx context.Context, req *auth.CreateUser) (*auth.UserLoginResponse, error) {
	exists, err := s.strg.User().ExistsByUsername(ctx, req.Username)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if exists {
		return nil, status.Error(codes.AlreadyExists, "username already taken")
	}

	id, err := s.strg.User().Create(ctx, req)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	user, err := s.strg.User().GetByID(ctx, id)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	res := &auth.UserLoginResponse{User: user}

	if req.GenerateToken {
		token, err := helper.GenerateToken(user.Id, s.cfg.TokenSecretKey, s.cfg.AccessTokenExpiring, s.cfg.RefreshTokenExpiring)
		if err != nil {
			return nil, status.Errorf(codes.Internal, err.Error())
		}

		res.AccessToken = token.AccessToken
		res.RefreshToken = token.RefreshToken
	}

	return res, nil
}

func (s *AuthService) GetByID(ctx context.Context, req *auth.UserPrimaryKey) (*auth.User, error) {
	res, err := s.strg.User().GetByID(ctx, req)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return res, nil
}

func (s *AuthService) GetList(ctx context.Context, req *auth.GetListUserRequest) (*auth.GetListUserResponse, error) {
	res, err := s.strg.User().GetAll(ctx, req)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return res, nil
}

func (s *AuthService) Update(ctx context.Context, req *auth.UpdateUser) (*auth.User, error) {
	err := s.strg.User().Update(ctx, req)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return s.GetByID(ctx, &auth.UserPrimaryKey{Id: req.Id})
}

func (s *AuthService) Delete(ctx context.Context, req *auth.UserPrimaryKey) (*auth.MessageString, error) {
	err := s.strg.User().Delete(ctx, req)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	res := &auth.MessageString{Message: "successfully deleted"}

	return res, nil
}
