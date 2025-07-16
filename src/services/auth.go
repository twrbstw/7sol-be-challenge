package services

import (
	"context"
	"errors"
	e "seven-solutions-challenge/src/errors"
	repositories "seven-solutions-challenge/src/repositories/userRep"
	"seven-solutions-challenge/src/requests"
	"seven-solutions-challenge/src/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IAuthService interface {
	Register(ctx context.Context, req requests.AuthRegisterReq) error
	Login(ctx context.Context, req requests.AuthLoginReq) error
}

type AuthService struct {
	userRepo repositories.IUserRepo
}

func NewAuthService(userRepo repositories.IUserRepo) IAuthService {
	return &AuthService{
		userRepo: userRepo,
	}
}

// Login implements IAuthService.
func (a *AuthService) Login(ctx context.Context, req requests.AuthLoginReq) error {
	panic("unimplemented")
}

// Register implements IAuthService.
func (a *AuthService) Register(ctx context.Context, req requests.AuthRegisterReq) error {
	hashedPassword, err := utils.HashString(req.Password)
	if err != nil {
		return errors.New(e.ERR_SERVICE_HASHING)
	}

	a.userRepo.Create(ctx, repositories.CreateReq{
		Id:        primitive.NewObjectID().Hex(),
		Name:      req.Name,
		Email:     req.Email,
		Password:  hashedPassword,
		CreatedAt: time.Now(),
	})

	return nil
}
