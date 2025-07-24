package services_test

import (
	"context"
	"errors"
	"seven-solutions-challenge/internal/adapters/inbound/http/requests"
	"seven-solutions-challenge/internal/adapters/inbound/http/responses"
	mockPorts "seven-solutions-challenge/internal/app/ports/mocks"
	"seven-solutions-challenge/internal/app/services"
	"seven-solutions-challenge/internal/domain"
	e "seven-solutions-challenge/pkg/errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

var ctx = context.Background()
var timeNow = time.Now()

func TestRegisterSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mockPorts.NewMockIUserRepo(ctrl)
	mockBcryptHasher := mockPorts.NewMockIHasher(ctrl)
	mockJwtGenerator := mockPorts.NewMockIJwtGenerator(ctrl)
	authService := services.NewAuthService(mockUserRepo, mockBcryptHasher, mockJwtGenerator)

	mockBcryptHasher.EXPECT().HashPassword(gomock.Any()).Return("test_hashed_password", nil)
	output := domain.User{
		Id:        "test_id",
		Name:      "test_name",
		Email:     "test_email",
		Password:  "test_password",
		CreatedAt: timeNow,
	}
	mockUserRepo.EXPECT().Create(gomock.Any(), gomock.Any()).Return(&output, nil)

	err := authService.Register(ctx, requests.AuthRegisterReq{
		Name:     "test_name",
		Email:    "test_email",
		Password: "test_password",
	})

	assert.NoError(t, err)
}

func TestRegisterCreateError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mockPorts.NewMockIUserRepo(ctrl)
	mockBcryptHasher := mockPorts.NewMockIHasher(ctrl)
	mockJwtGenerator := mockPorts.NewMockIJwtGenerator(ctrl)
	authService := services.NewAuthService(mockUserRepo, mockBcryptHasher, mockJwtGenerator)

	mockBcryptHasher.EXPECT().HashPassword(gomock.Any()).Return("test_hashed_password", nil)
	expected := errors.New("test_error")
	mockUserRepo.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil, errors.New("test_error"))

	err := authService.Register(ctx, requests.AuthRegisterReq{
		Name:     "test_name",
		Email:    "test_email",
		Password: "test_password",
	})

	assert.Equal(t, expected, err)
}

func TestRegisterHasherError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mockPorts.NewMockIUserRepo(ctrl)
	mockBcryptHasher := mockPorts.NewMockIHasher(ctrl)
	mockJwtGenerator := mockPorts.NewMockIJwtGenerator(ctrl)
	authService := services.NewAuthService(mockUserRepo, mockBcryptHasher, mockJwtGenerator)

	expected := errors.New(e.ERR_SERVICE_HASHING)
	mockBcryptHasher.EXPECT().HashPassword(gomock.Any()).Return("", errors.New(e.ERR_SERVICE_HASHING))

	err := authService.Register(ctx, requests.AuthRegisterReq{
		Name:     "test_name",
		Email:    "test_email",
		Password: "test_password",
	})

	assert.Equal(t, expected, err)
}

func TestLoginSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mockPorts.NewMockIUserRepo(ctrl)
	mockBcryptHasher := mockPorts.NewMockIHasher(ctrl)
	mockJwtGenerator := mockPorts.NewMockIJwtGenerator(ctrl)
	authService := services.NewAuthService(mockUserRepo, mockBcryptHasher, mockJwtGenerator)

	output := responses.AuthLoginResp{
		Email: "test_email",
		Token: "test_token",
	}

	mockUserRepo.EXPECT().GetByEmail(gomock.Any(), gomock.Any()).Return(&domain.User{
		Id:       "test_id",
		Name:     "test_name",
		Email:    "test_email",
		Password: "test_hashed_password",
	}, nil)
	mockBcryptHasher.EXPECT().ComparePassword(gomock.Any(), gomock.Any()).Return(nil)
	mockJwtGenerator.EXPECT().GenerateJwt(gomock.Any(), gomock.Any()).Return(output.Token, nil)

	resp, err := authService.Login(ctx, requests.AuthLoginReq{
		Email:    "test_email",
		Password: "test_password",
	})

	assert.NoError(t, err)
	assert.Equal(t, resp.Email, output.Email)
	assert.Equal(t, resp.Token, output.Token)
}

func TestLoginGetByEmailError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mockPorts.NewMockIUserRepo(ctrl)
	mockBcryptHasher := mockPorts.NewMockIHasher(ctrl)
	mockJwtGenerator := mockPorts.NewMockIJwtGenerator(ctrl)
	authService := services.NewAuthService(mockUserRepo, mockBcryptHasher, mockJwtGenerator)

	expected := errors.New("test_error")
	mockUserRepo.EXPECT().GetByEmail(gomock.Any(), gomock.Any()).Return(nil, errors.New("test_error"))

	resp, err := authService.Login(ctx, requests.AuthLoginReq{
		Email:    "test_email",
		Password: "test_password",
	})

	assert.Nil(t, resp)
	assert.Equal(t, expected, err)
}

func TestLoginBcryptHasherError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mockPorts.NewMockIUserRepo(ctrl)
	mockBcryptHasher := mockPorts.NewMockIHasher(ctrl)
	mockJwtGenerator := mockPorts.NewMockIJwtGenerator(ctrl)
	authService := services.NewAuthService(mockUserRepo, mockBcryptHasher, mockJwtGenerator)

	mockUserRepo.EXPECT().GetByEmail(gomock.Any(), gomock.Any()).Return(&domain.User{
		Id:       "test_id",
		Name:     "test_name",
		Email:    "test_email",
		Password: "test_hashed_password",
	}, nil)
	expected := errors.New(e.ERR_SERVICE_INCORRECT_EMAIL_OR_PASSWORD)
	mockBcryptHasher.EXPECT().ComparePassword(gomock.Any(), gomock.Any()).Return(errors.New(e.ERR_SERVICE_INCORRECT_EMAIL_OR_PASSWORD))

	resp, err := authService.Login(ctx, requests.AuthLoginReq{
		Email:    "test_email",
		Password: "test_password",
	})

	assert.Nil(t, resp)
	assert.Equal(t, expected, err)
}

func TestLoginGenerateJwtError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mockPorts.NewMockIUserRepo(ctrl)
	mockBcryptHasher := mockPorts.NewMockIHasher(ctrl)
	mockJwtGenerator := mockPorts.NewMockIJwtGenerator(ctrl)
	authService := services.NewAuthService(mockUserRepo, mockBcryptHasher, mockJwtGenerator)

	mockUserRepo.EXPECT().GetByEmail(gomock.Any(), gomock.Any()).Return(&domain.User{
		Id:       "test_id",
		Name:     "test_name",
		Email:    "test_email",
		Password: "test_hashed_password",
	}, nil)
	mockBcryptHasher.EXPECT().ComparePassword(gomock.Any(), gomock.Any()).Return(nil)
	expected := errors.New(e.ERR_SERVICE_GENERATING_JWT_FAILED)
	mockJwtGenerator.EXPECT().GenerateJwt(gomock.Any(), gomock.Any()).Return("", errors.New(e.ERR_SERVICE_GENERATING_JWT_FAILED))

	resp, err := authService.Login(ctx, requests.AuthLoginReq{
		Email:    "test_email",
		Password: "test_password",
	})

	assert.Nil(t, resp)
	assert.Equal(t, expected, err)
}
