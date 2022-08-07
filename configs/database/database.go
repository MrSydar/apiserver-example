package database

import (
	"context"
	"fmt"
	"log"
	"mrsydar/apiserver/configs/constants/envnames"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Client
var Collections struct {
	Accounts *mongo.Collection
}

func connectDB() (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	mongoURI := os.Getenv(envnames.MongoUri)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		return nil, fmt.Errorf("failed to create a new client: %v", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to ping the database: %v", err)
	}

	return client, nil
}

func Init() error {
	log.Print("Initializing database connection")

	var err error
	DB, err = connectDB()
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}

	databaseName := os.Getenv(envnames.DatabaseName)
	Collections.Accounts = DB.Database(databaseName).Collection(os.Getenv(envnames.AccountCollection))

	return nil
}
