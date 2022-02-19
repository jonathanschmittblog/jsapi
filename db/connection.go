package db

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var connection *mongo.Client

func GetDatabase() *mongo.Database {
	if connection == nil {
		connect()
	}
	return connection.Database(os.Getenv("JSAPI_MONGODB_DB_NAME"))
}

func connect() error {
	clientOptions := options.Client().ApplyURI(os.Getenv("JSAPI_MONGODB_URI"))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return err
	}
	connection = client
	return nil
}