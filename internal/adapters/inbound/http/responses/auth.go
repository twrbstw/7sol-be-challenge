package responses

type AuthLoginResp struct {
	Email string `json:"email"`
	Token string `json:"token"`
}
