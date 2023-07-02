package service

import (
	"context"
	"github.com/abdullayev13/ms_item_clickhead/auth/config"
	"github.com/abdullayev13/ms_item_clickhead/auth/genproto/auth"
	"github.com/abdullayev13/ms_item_clickhead/auth/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func (s *AuthService) CheckUrl(ctx context.Context, req *auth.CheckUrlRequest) (*auth.CheckUrlResponse, error) {
	res := &auth.CheckUrlResponse{Ok: true}
	return res, nil
}

func (s *AuthService) Login(ctx context.Context, req *auth.UserLoginRequest) (*auth.UserLoginResponse, error) {

	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
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
