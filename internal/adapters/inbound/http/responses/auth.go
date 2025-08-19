package responses

import "time"

type AuthRegisterResp struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
}

type AuthLoginResp struct {
	Email string `json:"email"`
	Token string `json:"token"`
}
