package handlers

import (
	"seven-solutions-challenge/internal/adapters/inbound/http/requests"
	"seven-solutions-challenge/internal/app/ports"
	"seven-solutions-challenge/pkg"
	e "seven-solutions-challenge/pkg/errors"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	authService ports.IAuthService
}

func NewAuthHandler(authService ports.IAuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (a AuthHandler) RegisterRoutes(r fiber.Router) {
	r.Post("/register", a.Register)
	r.Post("/login", a.Login)
}

func (a AuthHandler) Register(ctx *fiber.Ctx) error {
	var req requests.AuthRegisterReq
	err := ctx.BodyParser(&req)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(e.ERR_HANDLER_PARSING_REQ)
	}

	if err := pkg.ValidateBody(ctx, &req); err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	resp, err := a.authService.Register(ctx.Context(), req)
	if err != nil {
		respCode := e.HandleErrResp(err)
		return ctx.Status(respCode).SendString(err.Error())
	}

	return ctx.Status(fiber.StatusCreated).JSON(resp)
}

func (a AuthHandler) Login(ctx *fiber.Ctx) error {
	var req requests.AuthLoginReq
	err := ctx.BodyParser(&req)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(e.ERR_HANDLER_PARSING_REQ)
	}

	if err := pkg.ValidateBody(ctx, &req); err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	resp, err := a.authService.Login(ctx.Context(), req)
	if err != nil {
		respCode := e.HandleErrResp(err)
		return ctx.Status(respCode).SendString(err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(resp)
}
