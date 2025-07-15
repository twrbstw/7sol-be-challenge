package responses

import (
	"seven-solutions-challenge/src/models"
	"time"
)

type GetUserByIdResp struct {
	models.User
}

type CreateUserResp struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
}
