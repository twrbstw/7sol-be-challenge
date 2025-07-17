package requests

type AuthRegisterReq struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthLoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
