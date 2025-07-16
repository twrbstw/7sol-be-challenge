package services

import (
	"context"
	"errors"
	e "seven-solutions-challenge/src/errors"
	repositories "seven-solutions-challenge/src/repositories/userRep"
	"seven-solutions-challenge/src/requests"
	"seven-solutions-challenge/src/responses"
	"seven-solutions-challenge/src/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IUserService interface {
	GetById(ctx context.Context, req requests.GetUserByIdReq) (*responses.GetUserByIdResp, error)
	List(ctx context.Context) (*responses.ListUserResp, error)
	Create(ctx context.Context, req requests.CreateUserReq) (*responses.CreateUserResp, error)
	Update(ctx context.Context, req requests.UpdateUserReq) error
	Delete(ctx context.Context, req requests.DeleteUserReq) error
}

type UserService struct {
	userRepo repositories.IUserRepo
}

func NewUserService(userRepo repositories.IUserRepo) IUserService {
	return &UserService{
		userRepo: userRepo,
	}
}

// GetById implements IProfileService.
func (u *UserService) GetById(ctx context.Context, req requests.GetUserByIdReq) (*responses.GetUserByIdResp, error) {
	result, err := u.userRepo.GetById(ctx, repositories.GetByIdReq{Id: req.ID})
	if err != nil {
		return nil, err
	}

	return &responses.GetUserByIdResp{
		User: responses.UserResp{
			Id:        result.Id,
			Name:      result.Name,
			Email:     result.Email,
			CreatedAt: result.CreatedAt,
		},
	}, nil
}

// List implements IUserService.
func (u *UserService) List(ctx context.Context) (*responses.ListUserResp, error) {
	results, err := u.userRepo.List(ctx)
	if err != nil {
		return nil, err
	}

	var resp []responses.UserResp
	for _, v := range results {
		resp = append(resp, responses.UserResp{
			Id:        v.Id,
			Name:      v.Name,
			Email:     v.Email,
			CreatedAt: v.CreatedAt,
		})
	}

	return &responses.ListUserResp{
		Users: resp,
	}, nil
}

// Create implements IUserService.
func (u *UserService) Create(ctx context.Context, req requests.CreateUserReq) (*responses.CreateUserResp, error) {
	hashedPassword, err := utils.HashString(req.Password)
	if err != nil {
		return nil, errors.New(e.ERR_SERVICE_HASHING)
	}

	result, err := u.userRepo.Create(ctx, repositories.CreateReq{
		Id:        primitive.NewObjectID().Hex(),
		Name:      req.Name,
		Email:     req.Email,
		Password:  hashedPassword,
		CreatedAt: time.Now(),
	})
	if err != nil {
		return nil, err
	}

	return &responses.CreateUserResp{
		Id:        result.Id,
		Name:      result.Name,
		Email:     result.Email,
		CreatedAt: result.CreatedAt,
	}, nil
}

// Update implements IUserService.
func (u *UserService) Update(ctx context.Context, req requests.UpdateUserReq) error {
	err := u.userRepo.Update(ctx, repositories.UpdateReq{
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
func (u *UserService) Delete(ctx context.Context, req requests.DeleteUserReq) error {
	err := u.userRepo.Delete(ctx, repositories.DeleteReq{
		Id: req.Id,
	})
	if err != nil {
		return err
	}

	return nil
}
