package services

import (
	"context"
	"errors"
	"seven-solutions-challenge/internal/adapters/inbound/http/requests"
	"seven-solutions-challenge/internal/adapters/inbound/http/responses"
	mongoreq "seven-solutions-challenge/internal/adapters/outbound/db/mongo/requests"
	"seven-solutions-challenge/internal/app/ports"
	"seven-solutions-challenge/internal/domain"
	e "seven-solutions-challenge/internal/shared/errors"
	"seven-solutions-challenge/pkg"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo ports.IUserRepo
	appCfg   domain.AppConfig
}

func NewAuthService(userRepo ports.IUserRepo, appCfg domain.AppConfig) ports.IAuthService {
	return &AuthService{
		userRepo: userRepo,
		appCfg:   appCfg,
	}
}

// Login implements IAuthService.
func (a *AuthService) Login(ctx context.Context, req requests.AuthLoginReq) (*responses.AuthLoginResp, error) {
	user, err := a.userRepo.GetByEmail(ctx, mongoreq.GetByEmailReq{Email: req.Email})
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
	return &responses.AuthLoginResp{
		Email: user.Email,
		Token: token,
	}, nil
}

// Register implements IAuthService.
func (a *AuthService) Register(ctx context.Context, req requests.AuthRegisterReq) error {
	hashedPassword, err := pkg.HashString(req.Password)
	if err != nil {
		return errors.New(e.ERR_SERVICE_HASHING)
	}

	a.userRepo.Create(ctx, mongoreq.CreateReq{
		Id:        primitive.NewObjectID().Hex(),
		Name:      req.Name,
		Email:     req.Email,
		Password:  hashedPassword,
		CreatedAt: time.Now(),
	})

	return nil
}
