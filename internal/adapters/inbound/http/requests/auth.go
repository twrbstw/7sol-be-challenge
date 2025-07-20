package requests

import "seven-solutions-challenge/pkg"

type AuthRegisterReq struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (a *AuthRegisterReq) Validate() error {
	return pkg.ValidateJson(a)
}

type AuthLoginReq struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (a *AuthLoginReq) Validate() error {
	return pkg.ValidateJson(a)
}
