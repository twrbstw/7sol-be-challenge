package requests

import "seven-solutions-challenge/pkg"

type GetUserByIdReq struct {
	ID string
}

type CreateUserReq struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (c *CreateUserReq) Validate() error {
	return pkg.ValidateJson(c)
}

type UpdateUserReq struct {
	Id    string
	Name  string `json:"name,omitempty" validate:"required"`
	Email string `json:"email,omitempty" validate:"required,email"`
}

func (u *UpdateUserReq) IsEmailAndNameEmpty() bool {
	if u.Email == "" && u.Name == "" {
		return false
	}
	return true
}

func (u *UpdateUserReq) Validate() error {
	return pkg.ValidateJson(u)
}

type DeleteUserReq struct {
	Id string
}
