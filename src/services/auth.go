package services

import (
	"context"
	"errors"
	e "seven-solutions-challenge/src/errors"
	"seven-solutions-challenge/src/models"
	repositories "seven-solutions-challenge/src/repositories/userRep"
	"seven-solutions-challenge/src/requests"
	"seven-solutions-challenge/src/responses"
	"seven-solutions-challenge/src/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type IAuthService interface {
	Register(ctx context.Context, req requests.AuthRegisterReq) error
	Login(ctx context.Context, req requests.AuthLoginReq) (*responses.AuthLoginResp, error)
}

type AuthService struct {
	userRepo repositories.IUserRepo
	appCfg   models.AppConfig
}

func NewAuthService(userRepo repositories.IUserRepo, appCfg models.AppConfig) IAuthService {
	return &AuthService{
		userRepo: userRepo,
		appCfg:   appCfg,
	}
}

// Login implements IAuthService.
func (a *AuthService) Login(ctx context.Context, req requests.AuthLoginReq) (*responses.AuthLoginResp, error) {
	user, err := a.userRepo.GetByEmail(ctx, repositories.GetByEmail{Email: req.Email})
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, errors.New(e.ERR_SERVICE_INCORRECT_EMAIL_OR_PASSWORD)
	}

	token, err := utils.GenerateJwt(user.Name, user.Email, a.appCfg)
	if err != nil {
		return nil, errors.New(e.ERR_SERVICE_GENERATING_JWT_FAILED)
	}
	return &responses.AuthLoginResp{
		Email: user.Email,
		Token: token,
	}, nil
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
