package requests

import "time"

type GetByIdReq struct {
	Id string
}

type CreateReq struct {
	Id        string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

type UpdateReq struct {
	Id    string
	Name  string
	Email string
}

type DeleteReq struct {
	Id string
}

type GetByEmailReq struct {
	Email string
}
