package services

import (
	"context"
	"errors"
	"seven-solutions-challenge/internal/adapters/inbound/http/requests"
	"seven-solutions-challenge/internal/adapters/inbound/http/responses"
	mongoreq "seven-solutions-challenge/internal/adapters/outbound/db/mongo/requests"
	"seven-solutions-challenge/internal/app/ports"
	e "seven-solutions-challenge/pkg/errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthService struct {
	userRepo     ports.IUserRepo
	bcryptHasher ports.IHasher
	jwtGenerator ports.IJwtGenerator
}

func NewAuthService(userRepo ports.IUserRepo, bcryptHasher ports.IHasher, jwtGenerator ports.IJwtGenerator) ports.IAuthService {
	return &AuthService{
		userRepo:     userRepo,
		bcryptHasher: bcryptHasher,
		jwtGenerator: jwtGenerator,
	}
}

// Login implements IAuthService.
func (a *AuthService) Login(ctx context.Context, req requests.AuthLoginReq) (*responses.AuthLoginResp, error) {
	user, err := a.userRepo.GetByEmail(ctx, mongoreq.GetByEmailReq{Email: req.Email})
	if err != nil {
		return nil, err
	}

	err = a.bcryptHasher.ComparePassword(user.Password, req.Password)
	if err != nil {
		return nil, errors.New(e.ERR_SERVICE_INCORRECT_EMAIL_OR_PASSWORD)
	}

	token, err := a.jwtGenerator.GenerateJwt(user.Name, user.Email)
	if err != nil {
		return nil, errors.New(e.ERR_SERVICE_GENERATING_JWT_FAILED)
	}
	return &responses.AuthLoginResp{
		Email: user.Email,
		Token: token,
	}, nil
}

// Register implements IAuthService.
func (a *AuthService) Register(ctx context.Context, req requests.AuthRegisterReq) (*responses.AuthRegisterResp, error) {
	hashedPassword, err := a.bcryptHasher.HashPassword(req.Password)
	if err != nil {
		return nil, errors.New(e.ERR_SERVICE_HASHING)
	}

	user, err := a.userRepo.Create(ctx, mongoreq.CreateReq{
		Id:        primitive.NewObjectID().Hex(),
		Name:      req.Name,
		Email:     req.Email,
		Password:  hashedPassword,
		CreatedAt: time.Now(),
	})
	if err != nil {
		return nil, err
	}

	return &responses.AuthRegisterResp{Id: user.Id}, nil
}
