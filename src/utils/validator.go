package utils

import (
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

type Validator interface {
	Validate() error
}

func ValidateBody[T Validator](ctx *fiber.Ctx, body T) error {
	if err := body.Validate(); err != nil {
		return err
	}
	return nil
}

var v = validator.New()

func ValidateJson(s interface{}) error {
	// var result error
	err := v.Struct(s)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		return errors
	}

	return nil
}
