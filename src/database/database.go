package database

import (
	"context"
	"seven-solutions-challenge/src/models"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

type DatabaseConnection interface {
	GetCollection(c DatabaseCollection) *mongo.Collection
	Disconnect(ctx context.Context) error
}

type DatabaseClient struct {
	Client   *mongo.Client
	Database *mongo.Database
}

// GetCollection implements DatabaseConnection.
func (d *DatabaseClient) GetCollection(c DatabaseCollection) *mongo.Collection {
	return d.Database.Collection(c.String())
}

// Disconnect implements DatabaseConnection.
func (d *DatabaseClient) Disconnect(ctx context.Context) error {
	return d.Client.Disconnect(ctx)
}

func NewDatabaseClient(ctx context.Context, dbCfg models.DbConfig) DatabaseConnection {
	serverApi := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(dbCfg.Uri).SetServerAPIOptions(serverApi)

	client, err := mongo.Connect(opts)
	if err != nil {
		panic(err)
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	err = createUserIndex(ctx, client, dbCfg.Name)
	if err != nil {
		panic(err)
	}

	return &DatabaseClient{
		Client:   client,
		Database: client.Database(dbCfg.Name),
	}
}

func createUserIndex(ctx context.Context, c *mongo.Client, dbName string) error {
	collection := c.Database(dbName).Collection(COLLECTION_USERS.String())
	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "email", Value: 1}},
		Options: options.Index().SetUnique(true),
	}

	_, err := collection.Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		return err
	}
	return nil
}
