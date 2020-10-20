package mongocli

import (
	"context"
	"errors"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	db          *mongo.Client
	defaultHost = "mongodb://localhost:27017"
)

func GetDB() *mongo.Client {
	if db != nil {
		return db
	}

	uri := os.Getenv("DB_HOST")
	if uri == "" {
		uri = defaultHost
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	db = client

	return db
}

func Disconnect() error {
	if db == nil {
		return errors.New("there is no mongodb instance.")
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	db.Disconnect(ctx)
	return nil
}
