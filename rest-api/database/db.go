package database

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClinet *mongo.Client

func GetCollection(collectName string) *mongo.Collection {
	return MongoClinet.Database("demo_db").Collection(collectName)
}

func StartMongoDB() error {
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	var err error
	MongoClinet, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	return nil
}

func CloseMongoDB() {
	err := MongoClinet.Disconnect(context.TODO())
	if err != nil {
		panic(err)
	}
}
