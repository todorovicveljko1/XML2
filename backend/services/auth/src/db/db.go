package db

import (
	"context"
	"log"

	"auth.accommodation.com/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	// craete email index
	indexEmailModel := mongo.IndexModel{
		Keys:    bson.D{{"email", 1}},
		Options: options.Index().SetUnique(true),
	}
	// create usename index
	indexUsernameModel := mongo.IndexModel{
		Keys:    bson.D{{"username", 1}},
		Options: options.Index().SetUnique(true),
	}
	coll := client.Database("accommodation_auth").Collection("users")
	// create index
	_, err := coll.Indexes().CreateMany(context.Background(), []mongo.IndexModel{indexEmailModel, indexUsernameModel})
	if err != nil {
		panic(err)
	}
}
