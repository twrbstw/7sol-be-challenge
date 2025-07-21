package services_test

import (
	"context"
	"seven-solutions-challenge/internal/adapters/inbound/http/requests"
	"seven-solutions-challenge/internal/app/ports"
	mock_ports "seven-solutions-challenge/internal/app/ports/mocks"
	"seven-solutions-challenge/internal/app/services"
	"seven-solutions-challenge/internal/domain"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/mock/gomock"
)

var _ = Describe("AuthService", func() {
	var (
		ctx          context.Context
		ctrl         *gomock.Controller
		mockUserRepo *mock_ports.MockIUserRepo
		authService  ports.IAuthService
		outputUser   domain.User
	)

	BeforeEach(func() {
		ctx = context.Background()
		ctrl = gomock.NewController(GinkgoT())
		mockUserRepo = mock_ports.NewMockIUserRepo(ctrl)

		authService = services.NewAuthService(mockUserRepo, domain.AppConfig{})
		outputUser = domain.User{
			Id:        "test_id",
			Name:      "test_name",
			Email:     "test_email",
			Password:  "test_password",
			CreatedAt: time.Now(),
		}
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Context("Register", func() {
		It("should return no error on successful registration", func() {
			mockUserRepo.EXPECT().
				Create(gomock.Any(), gomock.Any()).
				Return(&outputUser, nil)

			err := authService.Register(ctx, requests.AuthRegisterReq{
				Name:     "test_name",
				Email:    "test_email",
				Password: "test_password",
			})

			Expect(err).To(BeNil())
		})
	})
})
