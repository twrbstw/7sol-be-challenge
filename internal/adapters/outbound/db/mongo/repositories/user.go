package repositories

import (
	"context"
	"errors"
	d "seven-solutions-challenge/internal/adapters/outbound/db/mongo"
	"seven-solutions-challenge/internal/adapters/outbound/db/mongo/requests"
	"seven-solutions-challenge/internal/app/ports"
	"seven-solutions-challenge/internal/domain"
	e "seven-solutions-challenge/internal/shared/errors"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type UserRepo struct {
	userCollection *mongo.Collection
}

func NewUserRepo(db d.DatabaseConnection) ports.IUserRepo {
	userCollection := db.GetCollection(d.COLLECTION_USERS)
	return &UserRepo{
		userCollection: userCollection,
	}
}

func (u *UserRepo) GetById(ctx context.Context, req requests.GetByIdReq) (*domain.User, error) {
	filter := bson.D{{Key: "_id", Value: req.Id}}

	var result domain.User
	err := u.userCollection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New(e.ERR_USER_NOT_FOUND)
		}
		return nil, errors.New(err.Error())
	}

	return &result, nil
}

func (u *UserRepo) Create(ctx context.Context, req requests.CreateReq) (*domain.User, error) {
	newUser := domain.User(req)
	_, err := u.userCollection.InsertOne(ctx, newUser)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return nil, errors.New(e.ERR_USER_EMAIL_DUPLICATED)
		}
		return nil, errors.New(err.Error())
	}

	return &domain.User{
		Id:        req.Id,
		Name:      req.Name,
		Email:     req.Email,
		Password:  req.Password,
		CreatedAt: req.CreatedAt,
	}, nil
}

func (u *UserRepo) List(ctx context.Context) ([]domain.User, error) {
	filter := bson.D{}

	cursor, err := u.userCollection.Find(ctx, filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New(e.ERR_USER_NOT_FOUND)
		}
		return nil, errors.New(err.Error())
	}

	var results []domain.User
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	return results, nil
}

func (u *UserRepo) Update(ctx context.Context, req requests.UpdateReq) error {
	_, err := u.GetById(ctx, requests.GetByIdReq{Id: req.Id})
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

func (u *UserRepo) Delete(ctx context.Context, req requests.DeleteReq) error {
	filter := bson.D{{Key: "_id", Value: req.Id}}

	_, err := u.userCollection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserRepo) GetByEmail(ctx context.Context, req requests.GetByEmailReq) (*domain.User, error) {
	filter := bson.D{{Key: "email", Value: req.Email}}

	var result domain.User
	err := u.userCollection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New(e.ERR_USER_NOT_FOUND)
		}
		return nil, errors.New(err.Error())
	}

	return &result, nil
}
