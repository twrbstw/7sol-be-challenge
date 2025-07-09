package services

import (
	"context"
	"errors"
	"seven-solutions-challenge/src/database"
	e "seven-solutions-challenge/src/errors"
	"seven-solutions-challenge/src/models"
	"seven-solutions-challenge/src/requests"
	"seven-solutions-challenge/src/responses"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type IUserService interface {
	GetById(ctx context.Context, req requests.GetUserByIdReq) (*responses.GetUserByIdResp, error)
	List() error
	Create(ctx context.Context) error
	Update() error
	Delete() error
}

type UserService struct {
	db database.DatabaseConnection
}

func NewUserService(db database.DatabaseConnection) IUserService {
	return &UserService{
		db: db,
	}
}

// GetById implements IProfileService.
func (u *UserService) GetById(ctx context.Context, req requests.GetUserByIdReq) (*responses.GetUserByIdResp, error) {

	collection := u.db.GetCollection(database.COLLECTION_USERS)

	filter := bson.D{{Key: "id", Value: "test_id"}}

	var result models.User
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New(e.ERR_USER_NOT_FOUND)
		}
	}
	return &responses.GetUserByIdResp{
		User: result,
	}, nil
}

// List implements IUserService.
func (u *UserService) List() error {
	panic("unimplemented")
}

// Create implements IUserService.
func (u *UserService) Create(ctx context.Context) error {
	panic("unimplemented")
}

// Update implements IUserService.
func (u *UserService) Update() error {
	panic("unimplemented")
}

// Delete implements IUserService.
func (u *UserService) Delete() error {
	panic("unimplemented")
}
