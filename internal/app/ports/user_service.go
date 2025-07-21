package ports

import (
	"context"
	"seven-solutions-challenge/internal/adapters/inbound/http/requests"
	"seven-solutions-challenge/internal/adapters/inbound/http/responses"
)

type IUserService interface {
	GetById(ctx context.Context, req requests.GetUserByIdReq) (*responses.GetUserByIdResp, error)
	List(ctx context.Context) (*responses.ListUserResp, error)
	Create(ctx context.Context, req requests.CreateUserReq) (*responses.CreateUserResp, error)
	Update(ctx context.Context, req requests.UpdateUserReq) error
	Delete(ctx context.Context, req requests.DeleteUserReq) error
}
