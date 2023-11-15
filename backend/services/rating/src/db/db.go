package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"rating.accommodation.com/config"
)

func DbInit(cfg *config.Config) (*mongo.Client, error) {
	uri := cfg.MongoDBURI
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	CreateIndex(client)
	return client, nil
}

func CreateIndex(client *mongo.Client) {
}
