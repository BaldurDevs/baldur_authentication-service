package config

import (
	"context"
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetConnection() (*mongo.Client, error) {
	var mongoUrl string
	godotenv.Load(".env")

	mongoUrl = os.Getenv("ATLAS_URL")
	if len(mongoUrl) == 0 {
		return nil, errors.New("no ATLAS_URL present")
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(mongoUrl))
	if err != nil {
		log.Println("Error creating new mongo client...")
		return nil, err
	}

	ctx, _ := context.WithTimeout(context.Background(), TimeOutInterval)
	err = client.Connect(ctx)
	if err != nil {
		log.Println("Error connecting mongo...")
		return nil, err
	}

	return client, nil
}

func CloseConnection(ctx context.Context, client *mongo.Client) {
	if err := client.Disconnect(ctx); err != nil {
		panic(err)
	}
}
