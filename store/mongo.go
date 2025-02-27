package store

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	client *mongo.Client
}

func (db *MongoDB) Connect(uri string) error {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return err
	}
	db.client = client
	return nil
}

func (db *MongoDB) Disconnect() error {
	return db.client.Disconnect(context.TODO())
}
