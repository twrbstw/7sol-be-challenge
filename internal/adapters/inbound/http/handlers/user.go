package handlers

import (
	"seven-solutions-challenge/internal/adapters/inbound/http/requests"
	"seven-solutions-challenge/internal/app/ports"
	"seven-solutions-challenge/pkg"
	e "seven-solutions-challenge/pkg/errors"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userService ports.IUserService
}

func NewUserHandler(userService ports.IUserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (u UserHandler) RegisterRoutes(r fiber.Router) {
	r.Get("/list", u.List)
	r.Get("/:uid", u.Get)
	r.Post("/", u.Create)
	r.Put("/:uid", u.Update)
	r.Delete("/:uid", u.Delete)
}

func (u UserHandler) Get(ctx *fiber.Ctx) error {
	uid := ctx.Params("uid")

	user, err := u.userService.GetById(ctx.Context(), requests.GetUserByIdReq{ID: uid})
	if err != nil {
		respCode := e.HandleErrResp(err)
		return ctx.Status(respCode).SendString(err.Error())
	}
	return ctx.Status(fiber.StatusOK).JSON(user)
}

func (u UserHandler) List(ctx *fiber.Ctx) error {
	users, err := u.userService.List(ctx.Context())
	if err != nil {
		respCode := e.HandleErrResp(err)
		return ctx.Status(respCode).SendString(err.Error())
	}
	return ctx.Status(fiber.StatusOK).JSON(users)
}

func (u UserHandler) Create(ctx *fiber.Ctx) error {
	var req requests.CreateUserReq
	err := ctx.BodyParser(&req)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(e.ERR_HANDLER_PARSING_REQ)
	}

	res, err := u.userService.Create(ctx.Context(), req)
	if err != nil {
		respCode := e.HandleErrResp(err)
		return ctx.Status(respCode).SendString(err.Error())
	}
	return ctx.Status(fiber.StatusCreated).JSON(res)
}

func (u UserHandler) Update(ctx *fiber.Ctx) error {
	uid := ctx.Params("uid")

	var req requests.UpdateUserReq
	err := ctx.BodyParser(&req)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(e.ERR_HANDLER_PARSING_REQ)
	}

	req.Id = uid

	if err := pkg.ValidateBody(ctx, &req); err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	if !req.IsEmailAndNameEmpty() {
		return ctx.Status(fiber.StatusBadRequest).SendString(e.ERR_HANDLER_NAME_OR_EMAIL_EMPTY)
	}

	err = u.userService.Update(ctx.Context(), req)
	if err != nil {
		respCode := e.HandleErrResp(err)
		return ctx.Status(respCode).SendString(err.Error())
	}

	return ctx.SendStatus(fiber.StatusOK)
}

func (u UserHandler) Delete(ctx *fiber.Ctx) error {
	uid := ctx.Params("uid")

	err := u.userService.Delete(ctx.Context(), requests.DeleteUserReq{Id: uid})
	if err != nil {
		respCode := e.HandleErrResp(err)
		return ctx.Status(respCode).SendString(err.Error())
	}

	return ctx.SendStatus(fiber.StatusOK)
}
