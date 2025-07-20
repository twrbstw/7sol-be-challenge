package responses

import (
	"time"
)

type UserResp struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
}

type GetUserByIdResp struct {
	User UserResp `json:"user"`
}

type CreateUserResp struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
}

type ListUserResp struct {
	Users []UserResp `json:"users"`
}
