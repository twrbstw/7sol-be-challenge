package handlers

import (
	e "seven-solutions-challenge/src/errors"
	"seven-solutions-challenge/src/requests"
	"seven-solutions-challenge/src/services"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userService services.IUserService
}

func NewUserHandler(userService services.IUserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (u UserHandler) RegisterRoutes(r fiber.Router) {
	r.Get("/:uid", u.Get)
	r.Get("/list", u.List)
	r.Post("/", u.Create)
	r.Put("/:uid", u.Update)
	r.Delete("/:uid", u.Delete)
}

func (u UserHandler) Get(ctx *fiber.Ctx) error {
	uid := ctx.Params("uid")

	user, err := u.userService.GetById(ctx.Context(), requests.GetUserByIdReq{ID: uid})
	if err != nil {
		respCode := handleErrResp(err)
		return ctx.Status(respCode).SendString(err.Error())
	}
	return ctx.Status(fiber.StatusOK).JSON(user)
}

func (u UserHandler) List(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).SendString("")
}

func (u UserHandler) Create(ctx *fiber.Ctx) error {
	var req requests.CreateUserReq
	err := ctx.BodyParser(&req)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(e.ERR_PARSING_REQ)
	}

	res, err := u.userService.Create(ctx.Context(), req)
	if err != nil {
		respCode := handleErrResp(err)
		return ctx.Status(respCode).SendString(err.Error())
	}
	return ctx.Status(fiber.StatusOK).JSON(res)
}

func (u UserHandler) Update(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).SendString("")
}

func (u UserHandler) Delete(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).SendString("")
}

func handleErrResp(err error) int {
	switch err.Error() {
	case e.ERR_USER_NOT_FOUND:
		return fiber.StatusNotFound
	case e.ERR_USER_EMAIL_DUPLICATED:
		return fiber.StatusBadRequest
	default:
		return fiber.StatusInternalServerError
	}
}
