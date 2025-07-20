package ports

import (
	"context"
	"seven-solutions-challenge/internal/adapters/outbound/db/mongo/requests"
	"seven-solutions-challenge/internal/domain"
)

type IUserRepo interface {
	GetById(ctx context.Context, req requests.GetByIdReq) (*domain.User, error)
	Create(ctx context.Context, req requests.CreateReq) (*domain.User, error)
	List(ctx context.Context) ([]domain.User, error)
	Update(ctx context.Context, req requests.UpdateReq) error
	Delete(ctx context.Context, req requests.DeleteReq) error
	GetByEmail(ctx context.Context, req requests.GetByEmailReq) (*domain.User, error)
}
