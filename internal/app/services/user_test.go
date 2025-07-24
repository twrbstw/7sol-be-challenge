package services_test

import (
	"errors"
	"seven-solutions-challenge/internal/adapters/inbound/http/requests"
	"seven-solutions-challenge/internal/adapters/inbound/http/responses"
	mockPorts "seven-solutions-challenge/internal/app/ports/mocks"
	"seven-solutions-challenge/internal/app/services"
	"seven-solutions-challenge/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetByIdSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mockPorts.NewMockIUserRepo(ctrl)
	mockBcryptHasher := mockPorts.NewMockIHasher(ctrl)
	userService := services.NewUserService(mockUserRepo, mockBcryptHasher)

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
}

func TestGetByIdError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mockPorts.NewMockIUserRepo(ctrl)
	mockBcryptHasher := mockPorts.NewMockIHasher(ctrl)
	userService := services.NewUserService(mockUserRepo, mockBcryptHasher)

	expected := errors.New("test_error")
	mockUserRepo.EXPECT().GetById(gomock.Any(), gomock.Any()).Return(nil, errors.New("test_error"))

	resp, err := userService.GetById(ctx, requests.GetUserByIdReq{
		ID: "test_id",
	})

	assert.Nil(t, resp)
	assert.Equal(t, expected, err)
}

func TestListSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mockPorts.NewMockIUserRepo(ctrl)
	mockBcryptHasher := mockPorts.NewMockIHasher(ctrl)
	userService := services.NewUserService(mockUserRepo, mockBcryptHasher)

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
}

func TestListError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mockPorts.NewMockIUserRepo(ctrl)
	mockBcryptHasher := mockPorts.NewMockIHasher(ctrl)
	userService := services.NewUserService(mockUserRepo, mockBcryptHasher)

	expected := errors.New("test_error")
	mockUserRepo.EXPECT().List(gomock.Any()).Return(nil, errors.New("test_error"))

	resp, err := userService.List(ctx)

	assert.Nil(t, resp)
	assert.Equal(t, expected, err)
}
