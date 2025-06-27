package utils

import (
	"context"
	"log"
	"time"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client
var MongoDatabase *mongo.Database
var (
	StudentCollection *mongo.Collection
	UserCollection    *mongo.Collection
)

func ConnectMongoDB(uri string) *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Ping to test connection
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatalf("MongoDB ping failed: %v", err)
	}

	log.Println("MongoDB connected successfully")
	initDatabase(client)
	return client
}

func initDatabase(client *mongo.Client) {
	MongoClient = client
	MongoDatabase = client.Database(viper.GetString("database.name"))

	StudentCollection = MongoDatabase.Collection(viper.GetString("database.mongodb.collections.students"))
	UserCollection = MongoDatabase.Collection(viper.GetString("database.mongodb.collections.users"))
}
