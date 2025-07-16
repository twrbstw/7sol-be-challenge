package requests

type GetUserByIdReq struct {
	ID string
}

type CreateUserReq struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUserReq struct {
	Id    string
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

func (uur UpdateUserReq) Validate() bool {
	if uur.Email == "" && uur.Name == "" {
		return false
	}
	return true
}

type DeleteUserReq struct {
	Id string
}
