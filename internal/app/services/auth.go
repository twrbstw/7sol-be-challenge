package services

import (
	"context"
	"errors"
	request "seven-solutions-challenge/internal/adapters/inbound/http/requests"
	response "seven-solutions-challenge/internal/adapters/inbound/http/responses"
	"seven-solutions-challenge/internal/adapters/outbound/db/mongo/repositories"
	"seven-solutions-challenge/internal/adapters/outbound/db/mongo/requests"
	"seven-solutions-challenge/internal/domain"
	e "seven-solutions-challenge/internal/shared/errors"
	"seven-solutions-challenge/pkg"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type IAuthService interface {
	Register(ctx context.Context, req request.AuthRegisterReq) error
	Login(ctx context.Context, req request.AuthLoginReq) (*response.AuthLoginResp, error)
}

type AuthService struct {
	userRepo repositories.IUserRepo
	appCfg   domain.AppConfig
}

func NewAuthService(userRepo repositories.IUserRepo, appCfg domain.AppConfig) IAuthService {
	return &AuthService{
		userRepo: userRepo,
		appCfg:   appCfg,
	}
}

// Login implements IAuthService.
func (a *AuthService) Login(ctx context.Context, req request.AuthLoginReq) (*response.AuthLoginResp, error) {
	user, err := a.userRepo.GetByEmail(ctx, requests.GetByEmailReq{Email: req.Email})
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, errors.New(e.ERR_SERVICE_INCORRECT_EMAIL_OR_PASSWORD)
	}

	token, err := pkg.GenerateJwt(user.Name, user.Email, a.appCfg)
	if err != nil {
		return nil, errors.New(e.ERR_SERVICE_GENERATING_JWT_FAILED)
	}
	return &response.AuthLoginResp{
		Email: user.Email,
		Token: token,
	}, nil
}

// Register implements IAuthService.
func (a *AuthService) Register(ctx context.Context, req request.AuthRegisterReq) error {
	hashedPassword, err := pkg.HashString(req.Password)
	if err != nil {
		return errors.New(e.ERR_SERVICE_HASHING)
	}

	a.userRepo.Create(ctx, requests.CreateReq{
		Id:        primitive.NewObjectID().Hex(),
		Name:      req.Name,
		Email:     req.Email,
		Password:  hashedPassword,
		CreatedAt: time.Now(),
	})

	return nil
}
