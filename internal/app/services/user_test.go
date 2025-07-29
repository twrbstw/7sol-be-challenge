package services_test

import (
	"errors"
	"seven-solutions-challenge/internal/adapters/inbound/http/requests"
	"seven-solutions-challenge/internal/adapters/inbound/http/responses"
	"seven-solutions-challenge/internal/app/ports"
	mockPorts "seven-solutions-challenge/internal/app/ports/mocks"
	"seven-solutions-challenge/internal/app/services"
	"seven-solutions-challenge/internal/domain"
	e "seven-solutions-challenge/pkg/errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func setUpUserServiceTest(t *testing.T) (ports.IUserService,
	*mockPorts.MockIUserRepo,
	*mockPorts.MockIHasher,
) {
	ctrl := gomock.NewController(t)
	t.Cleanup(ctrl.Finish)

	mockUserRepo := mockPorts.NewMockIUserRepo(ctrl)
	mockHasher := mockPorts.NewMockIHasher(ctrl)
	userService := services.NewUserService(mockUserRepo, mockHasher)

	return userService, mockUserRepo, mockHasher
}

func TestGetById(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		userService, mockUserRepo, _ := setUpUserServiceTest(t)

		expected := "test_id"
		mockUserRepo.EXPECT().GetById(gomock.Any(), gomock.Any()).Return(&domain.User{
			Id:        "test_id",
			Name:      "test_name",
			Email:     "test_email",
			Password:  "test_password",
			CreatedAt: timeNow,
		}, nil)

		resp, err := userService.GetById(ctx, requests.GetUserByIdReq{
			ID: "test_id",
		})

		assert.NoError(t, err)
		assert.Equal(t, expected, resp.User.Id)
	})

	t.Run("error", func(t *testing.T) {
		userService, mockUserRepo, _ := setUpUserServiceTest(t)

		expected := errors.New("test_error")
		mockUserRepo.EXPECT().GetById(gomock.Any(), gomock.Any()).Return(nil, errors.New("test_error"))

		resp, err := userService.GetById(ctx, requests.GetUserByIdReq{
			ID: "test_id",
		})

		assert.Nil(t, resp)
		assert.Equal(t, expected, err)
	})
}

func TestListUser(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		userService, mockUserRepo, _ := setUpUserServiceTest(t)
		expected := responses.ListUserResp{
			Users: []responses.UserResp{
				{
					Id:        "test_id_1",
					Name:      "test_name_1",
					Email:     "test_email_1",
					CreatedAt: timeNow,
				},
				{
					Id:        "test_id_2",
					Name:      "test_name_2",
					Email:     "test_email_2",
					CreatedAt: timeNow,
				},
			},
		}
		mockUserRepo.EXPECT().List(gomock.Any()).Return([]domain.User{
			{
				Id:        "test_id_1",
				Name:      "test_name_1",
				Email:     "test_email_1",
				CreatedAt: timeNow,
			},
			{
				Id:        "test_id_2",
				Name:      "test_name_2",
				Email:     "test_email_2",
				CreatedAt: timeNow,
			},
		}, nil)

		resp, err := userService.List(ctx)

		assert.NoError(t, err)
		assert.Equal(t, &expected, resp)
	})

	t.Run("error", func(t *testing.T) {
		userService, mockUserRepo, _ := setUpUserServiceTest(t)
		expected := errors.New("test_error")
		mockUserRepo.EXPECT().List(gomock.Any()).Return(nil, errors.New("test_error"))

		resp, err := userService.List(ctx)

		assert.Nil(t, resp)
		assert.Equal(t, expected, err)
	})
}

func TestCreateUser(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		userService, mockUserRepo, mockBcryptHasher := setUpUserServiceTest(t)

		req := requests.CreateUserReq{
			Name:     "test_name",
			Email:    "test_email",
			Password: "test_password",
		}

		hashedPassword := "test_hashed_password"
		expected := responses.CreateUserResp{
			Id:        "test_id",
			Name:      req.Name,
			Email:     req.Email,
			CreatedAt: timeNow,
		}

		mockBcryptHasher.EXPECT().HashPassword(gomock.Any()).Return(hashedPassword, nil)
		mockUserRepo.EXPECT().Create(gomock.Any(), gomock.Any()).Return(&domain.User{
			Id:        "test_id",
			Name:      req.Name,
			Email:     req.Email,
			Password:  hashedPassword,
			CreatedAt: timeNow,
		}, nil)

		resp, err := userService.Create(ctx, req)

		assert.Nil(t, err)
		assert.Equal(t, &expected, resp)
	})

	t.Run("hash error", func(t *testing.T) {
		userService, _, mockBcryptHasher := setUpUserServiceTest(t)
		req := requests.CreateUserReq{
			Name:     "test_name",
			Email:    "test_email",
			Password: "test_password",
		}

		expected := errors.New(e.ERR_SERVICE_HASHING)

		mockBcryptHasher.EXPECT().HashPassword(gomock.Any()).Return("", errors.New(e.ERR_SERVICE_HASHING))

		resp, err := userService.Create(ctx, req)

		assert.Nil(t, resp)
		assert.Equal(t, expected, err)
	})

	t.Run("create user repo error", func(t *testing.T) {
		userService, mockUserRepo, mockBcryptHasher := setUpUserServiceTest(t)
		req := requests.CreateUserReq{
			Name:     "test_name",
			Email:    "test_email",
			Password: "test_password",
		}

		hashedPassword := "test_hashed_password"
		expected := errors.New("test_error")

		mockBcryptHasher.EXPECT().HashPassword(gomock.Any()).Return(hashedPassword, nil)
		mockUserRepo.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil, errors.New("test_error"))

		resp, err := userService.Create(ctx, req)

		assert.Nil(t, resp)
		assert.Equal(t, expected, err)
	})
}

func TestUpdateUser(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		userService, mockUserRepo, _ := setUpUserServiceTest(t)

		mockUserRepo.EXPECT().Update(gomock.Any(), gomock.Any()).Return(nil)

		err := userService.Update(ctx, requests.UpdateUserReq{
			Id:    "test_id",
			Name:  "test_name",
			Email: "test_email",
		})

		assert.Nil(t, err)
	})

	t.Run("error", func(t *testing.T) {
		userService, mockUserRepo, _ := setUpUserServiceTest(t)

		expected := errors.New("test_error")
		mockUserRepo.EXPECT().Update(gomock.Any(), gomock.Any()).Return(errors.New("test_error"))

		err := userService.Update(ctx, requests.UpdateUserReq{
			Id:    "test_id",
			Name:  "test_name",
			Email: "test_email",
		})

		assert.Equal(t, expected, err)
	})
}

func TestDeleteUser(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		userService, mockUserRepo, _ := setUpUserServiceTest(t)

		mockUserRepo.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(nil)

		err := userService.Delete(ctx, requests.DeleteUserReq{
			Id: "test_id",
		})

		assert.Nil(t, err)
	})

	t.Run("error", func(t *testing.T) {
		userService, mockUserRepo, _ := setUpUserServiceTest(t)

		expected := errors.New("test_error")
		mockUserRepo.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(errors.New("test_error"))

		err := userService.Delete(ctx, requests.DeleteUserReq{
			Id: "test_id",
		})

		assert.Equal(t, expected, err)
	})
}
