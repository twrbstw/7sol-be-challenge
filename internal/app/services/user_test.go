package services_test

import (
	"errors"
	"seven-solutions-challenge/internal/adapters/inbound/http/requests"
	"seven-solutions-challenge/internal/adapters/inbound/http/responses"
	mockPorts "seven-solutions-challenge/internal/app/ports/mocks"
	"seven-solutions-challenge/internal/app/services"
	"seven-solutions-challenge/internal/domain"
	e "seven-solutions-challenge/pkg/errors"
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

func TestCreateSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mockPorts.NewMockIUserRepo(ctrl)
	mockBcryptHasher := mockPorts.NewMockIHasher(ctrl)
	userService := services.NewUserService(mockUserRepo, mockBcryptHasher)

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
}

func TestCreateBcryptHasherError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mockPorts.NewMockIUserRepo(ctrl)
	mockBcryptHasher := mockPorts.NewMockIHasher(ctrl)
	userService := services.NewUserService(mockUserRepo, mockBcryptHasher)

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
}

func TestCreateUserRepoError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mockPorts.NewMockIUserRepo(ctrl)
	mockBcryptHasher := mockPorts.NewMockIHasher(ctrl)
	userService := services.NewUserService(mockUserRepo, mockBcryptHasher)

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
}

func TestUpdateSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mockPorts.NewMockIUserRepo(ctrl)
	mockBcryptHasher := mockPorts.NewMockIHasher(ctrl)
	userService := services.NewUserService(mockUserRepo, mockBcryptHasher)

	mockUserRepo.EXPECT().Update(gomock.Any(), gomock.Any()).Return(nil)

	err := userService.Update(ctx, requests.UpdateUserReq{
		Id:    "test_id",
		Name:  "test_name",
		Email: "test_email",
	})

	assert.Nil(t, err)
}

func TestUpdateError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mockPorts.NewMockIUserRepo(ctrl)
	mockBcryptHasher := mockPorts.NewMockIHasher(ctrl)
	userService := services.NewUserService(mockUserRepo, mockBcryptHasher)

	expected := errors.New("test_error")
	mockUserRepo.EXPECT().Update(gomock.Any(), gomock.Any()).Return(errors.New("test_error"))

	err := userService.Update(ctx, requests.UpdateUserReq{
		Id:    "test_id",
		Name:  "test_name",
		Email: "test_email",
	})

	assert.Equal(t, expected, err)
}

func TestDeleteSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mockPorts.NewMockIUserRepo(ctrl)
	mockBcryptHasher := mockPorts.NewMockIHasher(ctrl)
	userService := services.NewUserService(mockUserRepo, mockBcryptHasher)

	mockUserRepo.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(nil)

	err := userService.Delete(ctx, requests.DeleteUserReq{
		Id: "test_id",
	})

	assert.Nil(t, err)
}

func TestDeleteError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mockPorts.NewMockIUserRepo(ctrl)
	mockBcryptHasher := mockPorts.NewMockIHasher(ctrl)
	userService := services.NewUserService(mockUserRepo, mockBcryptHasher)

	expected := errors.New("test_error")
	mockUserRepo.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(errors.New("test_error"))

	err := userService.Delete(ctx, requests.DeleteUserReq{
		Id: "test_id",
	})

	assert.Equal(t, expected, err)
}
