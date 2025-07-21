package services_test

import (
	"context"
	"seven-solutions-challenge/internal/adapters/inbound/http/requests"
	mock_ports "seven-solutions-challenge/internal/app/ports/mocks"
	"seven-solutions-challenge/internal/app/services"
	"seven-solutions-challenge/internal/domain"
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

	mockUserRepo := mock_ports.NewMockIUserRepo(ctrl)
	authService := services.NewAuthService(mockUserRepo, domain.AppConfig{})

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
