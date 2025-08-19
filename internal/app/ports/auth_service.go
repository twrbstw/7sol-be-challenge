package ports

import (
	"context"
	"seven-solutions-challenge/internal/adapters/inbound/http/requests"
	"seven-solutions-challenge/internal/adapters/inbound/http/responses"
)

type IAuthService interface {
	Register(ctx context.Context, req requests.AuthRegisterReq) (*responses.AuthRegisterResp, error)
	Login(ctx context.Context, req requests.AuthLoginReq) (*responses.AuthLoginResp, error)
}
