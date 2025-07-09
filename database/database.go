package database

import (
	"context"
	"seven-solutions-assignment/models"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

var _ DatabaseConnection = (*DatabaseClient)(nil)

type DatabaseConnection interface {
	GetCollection(c DatabaseCollection, opts ...*options.CollectionOptions) *mongo.Collection
	Disconnect(ctx context.Context) error
	Ping(ctx context.Context) error
}

type DatabaseClient struct {
	Client   *mongo.Client
	Database *mongo.Database
}

// Disconnect implements DatabaseConnection.
func (d *DatabaseClient) Disconnect(ctx context.Context) error {
	return d.Client.Disconnect(ctx)
}

// GetCollection implements DatabaseConnection.
func (d *DatabaseClient) GetCollection(c DatabaseCollection, opts ...*options.CollectionOptions) *mongo.Collection {
	return d.Database.Collection(c.String())
}

// Ping implements DatabaseConnection.
func (d *DatabaseClient) Ping(ctx context.Context) error {
	return d.Client.Ping(ctx, readpref.Primary())
}

func NewDatabaseClient(ctx context.Context, dbCfg models.DbConfig) *DatabaseClient {
	opts := options.Client().ApplyURI(dbCfg.Uri)

	client, err := mongo.Connect(opts)
	if err != nil {
		panic("failed to connect to mongodb")
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic("failed to ping mongodb")
	}

	return &DatabaseClient{
		Client:   client,
		Database: client.Database(dbCfg.Name),
	}
}
