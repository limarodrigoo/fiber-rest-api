package db

import (
	"context"
	"fmt"
	"rest-api/config"
	"rest-api/helper"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoInstance struct {
	Client           *mongo.Client
	DB               *mongo.Database
	CryptoCollection *mongo.Collection
}

var MI MongoInstance
var ctx = context.TODO()

func ConnectDatabase(config *config.Config) *mongo.Database {
	uri := getUri(*config)

	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		helper.ErrorPanic(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		helper.ErrorPanic(err)
	}

	return client.Database(config.DatabaseName)
}

func getUri(config config.Config) string {
	return fmt.Sprintf("mongodb://%s:%s@%s:%s/", config.DatabaseUser, config.DatabasePassword, config.DatabaseHost, config.DatabasePort)
}
