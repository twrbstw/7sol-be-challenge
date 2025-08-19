package responses

type AuthRegisterResp struct {
	Id string `json:"id"`
}

type AuthLoginResp struct {
	Email string `json:"email"`
	Token string `json:"token"`
}
