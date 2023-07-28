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

func ConnectDatabase(config *config.Config) {
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

	transactionCollection := client.Database(config.DatabaseName).Collection(config.DatabaseCollection)
	// blockCollection := client.Database(config.DatabaseName).Collection(config.DatabaseCollection)

	MI = MongoInstance{
		Client:           client,
		DB:               client.Database(config.DatabaseName),
		CryptoCollection: transactionCollection,
	}

}

func getUri(config config.Config) string {
	return fmt.Sprintf("mongodb://%s:%s@%s:%s/", config.DatabaseUser, config.DatabasePassword, config.DatabaseHost, config.DatabasePort)
}
