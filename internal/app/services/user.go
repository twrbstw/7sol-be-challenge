package services

import (
	"context"
	"errors"
	request "seven-solutions-challenge/internal/adapters/inbound/http/requests"
	response "seven-solutions-challenge/internal/adapters/inbound/http/responses"
	"seven-solutions-challenge/internal/adapters/outbound/db/mongo/requests"
	"seven-solutions-challenge/internal/app/ports"
	e "seven-solutions-challenge/pkg/errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService struct {
	userRepo     ports.IUserRepo
	bcryptHasher ports.IHasher
}

func NewUserService(userRepo ports.IUserRepo, bcryptHasher ports.IHasher) ports.IUserService {
	return &UserService{
		userRepo:     userRepo,
		bcryptHasher: bcryptHasher,
	}
}

// GetById implements IProfileService.
func (u *UserService) GetById(ctx context.Context, req request.GetUserByIdReq) (*response.GetUserByIdResp, error) {
	result, err := u.userRepo.GetById(ctx, requests.GetByIdReq{Id: req.ID})
	if err != nil {
		return nil, err
	}

	return &response.GetUserByIdResp{
		User: response.UserResp{
			Id:        result.Id,
			Name:      result.Name,
			Email:     result.Email,
			CreatedAt: result.CreatedAt,
		},
	}, nil
}

// List implements IUserService.
func (u *UserService) List(ctx context.Context) (*response.ListUserResp, error) {
	results, err := u.userRepo.List(ctx)
	if err != nil {
		return nil, err
	}

	var resp []response.UserResp
	for _, v := range results {
		resp = append(resp, response.UserResp{
			Id:        v.Id,
			Name:      v.Name,
			Email:     v.Email,
			CreatedAt: v.CreatedAt,
		})
	}

	return &response.ListUserResp{
		Users: resp,
	}, nil
}

// Create implements IUserService.
func (u *UserService) Create(ctx context.Context, req request.CreateUserReq) (*response.CreateUserResp, error) {
	hashedPassword, err := u.bcryptHasher.HashPassword(req.Password)
	if err != nil {
		return nil, errors.New(e.ERR_SERVICE_HASHING)
	}

	result, err := u.userRepo.Create(ctx, requests.CreateReq{
		Id:        primitive.NewObjectID().Hex(),
		Name:      req.Name,
		Email:     req.Email,
		Password:  hashedPassword,
		CreatedAt: time.Now(),
	})
	if err != nil {
		return nil, err
	}

	return &response.CreateUserResp{
		Id:        result.Id,
		Name:      result.Name,
		Email:     result.Email,
		CreatedAt: result.CreatedAt,
	}, nil
}

// Update implements IUserService.
func (u *UserService) Update(ctx context.Context, req request.UpdateUserReq) error {
	err := u.userRepo.Update(ctx, requests.UpdateReq{
		Id:    req.Id,
		Name:  req.Name,
		Email: req.Email,
	})
	if err != nil {
		return err
	}

	return nil
}

// Delete implements IUserService.
func (u *UserService) Delete(ctx context.Context, req request.DeleteUserReq) error {
	err := u.userRepo.Delete(ctx, requests.DeleteReq{
		Id: req.Id,
	})
	if err != nil {
		return err
	}

	return nil
}
