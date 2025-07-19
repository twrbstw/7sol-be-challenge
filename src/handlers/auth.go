package handlers

import (
	e "seven-solutions-challenge/src/errors"
	"seven-solutions-challenge/src/requests"
	"seven-solutions-challenge/src/services"
	"seven-solutions-challenge/src/utils"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	authService services.IAuthService
}

func NewAuthHandler(authService services.IAuthService) *AuthHandler {
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

	if err := utils.ValidateBody(ctx, &req); err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	err = a.authService.Register(ctx.Context(), req)
	if err != nil {
		respCode := e.HandleErrResp(err)
		return ctx.Status(respCode).SendString(err.Error())
	}

	return ctx.SendStatus(fiber.StatusCreated)
}

func (a AuthHandler) Login(ctx *fiber.Ctx) error {
	var req requests.AuthLoginReq
	err := ctx.BodyParser(&req)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(e.ERR_HANDLER_PARSING_REQ)
	}

	if err := utils.ValidateBody(ctx, &req); err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	resp, err := a.authService.Login(ctx.Context(), req)
	if err != nil {
		respCode := e.HandleErrResp(err)
		return ctx.Status(respCode).SendString(err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(resp)
}
