package config

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func ConnectDatabase() {
	_ = godotenv.Load()

	mongoURI := os.Getenv("MONGO_URI")
	mongoDB := os.Getenv("MONGO_DB")

	if mongoURI == "" || mongoDB == "" {
		log.Fatal("Missing MongoDB URI or Database name in environment variables")
	}

	clientOptions := options.Client().ApplyURI(mongoURI)

	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal("Failed to create MongoDB client")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err = client.Connect(ctx); err != nil {
		log.Fatal("Unable to connect to MongoDB")
	}

	if err = client.Ping(ctx, nil); err != nil {
		log.Fatal("Unable to ping MongoDB")
	}

	DB = client.Database(mongoDB)

	log.Println("Successfully connected to MongoDB")
}