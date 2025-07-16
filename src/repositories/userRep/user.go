package repositories

import (
	"context"
	"errors"
	"seven-solutions-challenge/src/database"
	e "seven-solutions-challenge/src/errors"
	"seven-solutions-challenge/src/models"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type IUserRepo interface {
	GetById(ctx context.Context, req GetByIdReq) (*models.User, error)
	Create(ctx context.Context, req CreateReq) (*models.User, error)
	List(ctx context.Context) ([]models.User, error)
	Update(ctx context.Context, req UpdateReq) error
	Delete(ctx context.Context, req DeleteReq) error
}

type UserRepo struct {
	userCollection *mongo.Collection
}

func NewUserRepo(db database.DatabaseConnection) IUserRepo {
	userCollection := db.GetCollection(database.COLLECTION_USERS)
	return &UserRepo{
		userCollection: userCollection,
	}
}

func (u *UserRepo) GetById(ctx context.Context, req GetByIdReq) (*models.User, error) {
	filter := bson.D{{Key: "_id", Value: req.Id}}

	var result models.User
	err := u.userCollection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New(e.ERR_USER_NOT_FOUND)
		}
		return nil, errors.New(err.Error())
	}

	return &models.User{
		Id:        result.Id,
		Name:      result.Name,
		Email:     result.Email,
		Password:  result.Password,
		CreatedAt: result.CreatedAt,
	}, nil
}

func (u *UserRepo) Create(ctx context.Context, req CreateReq) (*models.User, error) {
	newUser := models.User(req)
	_, err := u.userCollection.InsertOne(ctx, newUser)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return nil, errors.New(e.ERR_USER_EMAIL_DUPLICATED)
		}
		return nil, errors.New(err.Error())
	}

	return &models.User{
		Id:        req.Id,
		Name:      req.Name,
		Email:     req.Email,
		Password:  req.Password,
		CreatedAt: req.CreatedAt,
	}, nil
}

func (u *UserRepo) List(ctx context.Context) ([]models.User, error) {
	filter := bson.D{}

	cursor, err := u.userCollection.Find(ctx, filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New(e.ERR_USER_NOT_FOUND)
		}
		return nil, errors.New(err.Error())
	}

	var results []models.User
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	return results, nil
}

func (u *UserRepo) Update(ctx context.Context, req UpdateReq) error {
	_, err := u.GetById(ctx, GetByIdReq{Id: req.Id})
	if err != nil {
		return err
	}

	filter := bson.D{{Key: "_id", Value: req.Id}}
	var update bson.D
	if req.Name != "" {
		update = append(update, bson.E{Key: "name", Value: req.Name})
	}
	if req.Email != "" {
		update = append(update, bson.E{Key: "email", Value: req.Email})
	}
	setMap := bson.D{
		{Key: "$set", Value: update},
	}

	_, err = u.userCollection.UpdateOne(ctx, filter, setMap)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return errors.New(e.ERR_USER_EMAIL_DUPLICATED)
		}
		return err
	}

	return nil
}

func (u *UserRepo) Delete(ctx context.Context, req DeleteReq) error {
	filter := bson.D{{Key: "_id", Value: req.Id}}

	_, err := u.userCollection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}
