package models

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Database *mongo.Database

func InitDB() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	mongodbURI := os.Getenv("MONGODB_URI")
	if mongodbURI == "" {
		log.Fatalf("MONGODB_URI environment variable is not set")
	}

	clientOptions := options.Client().ApplyURI(mongodbURI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return err
	}

	Database = client.Database("mOx652024")

	log.Println("Connected to MongoDB successfully")

	return nil
}
