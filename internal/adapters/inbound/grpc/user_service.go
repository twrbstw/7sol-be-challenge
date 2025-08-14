package grpc

import (
	"context"
	"seven-solutions-challenge/internal/adapters/inbound/http/requests"
	"seven-solutions-challenge/internal/app/ports"
	"seven-solutions-challenge/proto/userpb"
)

type UserServiceServer struct {
	userpb.UnimplementedUserServiceServer
	UserService ports.IUserService
}

func (s *UserServiceServer) CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	res, err := s.UserService.Create(ctx, requests.CreateUserReq{
		Name:     req.GetName(),
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	})
	if err != nil {
		return nil, err
	}

	return &userpb.CreateUserResponse{Id: res.Id}, nil
}

func (s *UserServiceServer) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	resp, err := s.UserService.GetById(ctx, requests.GetUserByIdReq{
		ID: req.GetId(),
	})
	if err != nil {
		return nil, err
	}

	return &userpb.GetUserResponse{
		Id:    resp.User.Id,
		Name:  resp.User.Name,
		Email: resp.User.Email,
	}, nil
}
