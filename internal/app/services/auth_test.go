package services_test

import (
	"context"
	"errors"
	"seven-solutions-challenge/internal/adapters/inbound/http/requests"
	"seven-solutions-challenge/internal/adapters/inbound/http/responses"
	"seven-solutions-challenge/internal/app/ports"
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

func setup(t *testing.T) (ports.IAuthService,
	*mockPorts.MockIUserRepo,
	*mockPorts.MockIHasher,
	*mockPorts.MockIJwtGenerator,
) {
	ctrl := gomock.NewController(t)
	t.Cleanup(ctrl.Finish)

	mockUserRepo := mockPorts.NewMockIUserRepo(ctrl)
	mockHasher := mockPorts.NewMockIHasher(ctrl)
	mockJwtGen := mockPorts.NewMockIJwtGenerator(ctrl)
	authService := services.NewAuthService(mockUserRepo, mockHasher, mockJwtGen)

	return authService, mockUserRepo, mockHasher, mockJwtGen
}

func TestRegister(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		authService, mockUserRepo, mockBcryptHasher, _ := setup(t)

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
	})

	t.Run("create error", func(t *testing.T) {
		authService, mockUserRepo, mockBcryptHasher, _ := setup(t)

		mockBcryptHasher.EXPECT().HashPassword(gomock.Any()).Return("test_hashed_password", nil)
		expected := errors.New("test_error")
		mockUserRepo.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil, errors.New("test_error"))

		err := authService.Register(ctx, requests.AuthRegisterReq{
			Name:     "test_name",
			Email:    "test_email",
			Password: "test_password",
		})

		assert.Equal(t, expected, err)
	})

	t.Run("hash error", func(t *testing.T) {
		authService, _, mockBcryptHasher, _ := setup(t)

		expected := errors.New(e.ERR_SERVICE_HASHING)
		mockBcryptHasher.EXPECT().HashPassword(gomock.Any()).Return("", errors.New(e.ERR_SERVICE_HASHING))

		err := authService.Register(ctx, requests.AuthRegisterReq{
			Name:     "test_name",
			Email:    "test_email",
			Password: "test_password",
		})

		assert.Equal(t, expected, err)
	})
}

func TestLogin(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		authService, mockUserRepo, mockBcryptHasher, mockJwtGenerator := setup(t)

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
	})

	t.Run("get by email error", func(t *testing.T) {
		authService, mockUserRepo, _, _ := setup(t)

		expected := errors.New("test_error")
		mockUserRepo.EXPECT().GetByEmail(gomock.Any(), gomock.Any()).Return(nil, errors.New("test_error"))

		resp, err := authService.Login(ctx, requests.AuthLoginReq{
			Email:    "test_email",
			Password: "test_password",
		})

		assert.Nil(t, resp)
		assert.Equal(t, expected, err)
	})

	t.Run("hash error", func(t *testing.T) {
		authService, mockUserRepo, mockBcryptHasher, _ := setup(t)

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
	})

	t.Run("", func(t *testing.T) {
		authService, mockUserRepo, mockBcryptHasher, mockJwtGenerator := setup(t)

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
	})
}
