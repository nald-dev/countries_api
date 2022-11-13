package configs

import (
	"context"
	"countries_api/preferences"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func EnvMongoURI() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("MONGO_URI")
}

func ConnectDB() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(EnvMongoURI()))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	//ping the database
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB")

	// Uncomment here if you want to generate indexes
	// GenerateIndexes(client, ctx)

	return client
}

// Client instance
var DB = ConnectDB()

// getting database collections
func GetCollection(collectionName string) *mongo.Collection {
	collection := DB.Database(preferences.DB_NAME).Collection(collectionName)

	return collection
}

func GenerateIndexes(client *mongo.Client, ctx context.Context) {
	// Users should have unique username

	usersOpt := options.Index()

	usersOpt.SetName("Users")
	usersOpt.SetUnique(true)

	usersIndex := mongo.IndexModel{Keys: bson.M{"username": 1}, Options: usersOpt}
	if _, err := client.Database(preferences.DB_NAME).Collection("users").Indexes().CreateOne(ctx, usersIndex); err != nil {
		log.Println(err)
	}

	// Countries should have unique name

	countriesOpt := options.Index()

	countriesOpt.SetName("Countries")
	countriesOpt.SetUnique(true)

	countriesIndex := mongo.IndexModel{Keys: bson.M{"name": 1}, Options: countriesOpt}
	if _, err := client.Database(preferences.DB_NAME).Collection("countries").Indexes().CreateOne(ctx, countriesIndex); err != nil {
		log.Println(err)
	}
}
