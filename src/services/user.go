package services

import (
	"context"
	"errors"
	"seven-solutions-challenge/src/database"
	e "seven-solutions-challenge/src/errors"
	"seven-solutions-challenge/src/models"
	"seven-solutions-challenge/src/requests"
	"seven-solutions-challenge/src/responses"
	"seven-solutions-challenge/src/utils"
	"time"

	"github.com/gofiber/fiber/v2/log"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type IUserService interface {
	GetById(ctx context.Context, req requests.GetUserByIdReq) (*responses.GetUserByIdResp, error)
	List() error
	Create(ctx context.Context, req requests.CreateUserReq) (*responses.CreateUserResp, error)
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
		return nil, errors.New(err.Error())
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
func (u *UserService) Create(ctx context.Context, req requests.CreateUserReq) (*responses.CreateUserResp, error) {
	hashedPassword, err := utils.HashString(req.Password)
	if err != nil {
		log.Info("Error hashing:", err)
	}

	newUser := models.User{
		ID:        primitive.NewObjectID().Hex(),
		Name:      req.Name,
		Email:     req.Email,
		Password:  hashedPassword,
		CreatedAt: time.Now(),
	}

	result, err := u.db.GetCollection(database.COLLECTION_USERS).InsertOne(ctx, newUser)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return nil, errors.New(e.ERR_USER_EMAIL_DUPLICATED)
		}
		return nil, errors.New(err.Error())
	}

	id, ok := result.InsertedID.(string)
	if !ok {
		log.Fatal("InsertedID is not an ObjectID")
	}

	return &responses.CreateUserResp{
		ID:        id,
		Name:      newUser.Name,
		Email:     newUser.Email,
		CreatedAt: newUser.CreatedAt,
	}, nil
}

// Update implements IUserService.
func (u *UserService) Update() error {
	panic("unimplemented")
}

// Delete implements IUserService.
func (u *UserService) Delete() error {
	panic("unimplemented")
}
