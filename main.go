package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	return client
}

// Client instance
var DB *mongo.Client = ConnectDB()

// getting database collections
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("golangAPI").Collection(collectionName)
	return collection
}

type Country struct {
	Id          primitive.ObjectID `json:"id,omitempty"`
	Name        string             `json:"name,omitempty" validate:"required"`
	CapitalCity string             `json:"capitalCity,omitempty" validate:"required"`
	Currency    string             `json:"currency,omitempty" validate:"required"`
}

var countryCollection *mongo.Collection = GetCollection(DB, "countries")

func countries(c *fiber.Ctx) error {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	results, _ := countryCollection.Find(ctx, bson.M{})

	var countries []Country

	// reading from the db in an optimal way
	defer results.Close(ctx)

	for results.Next(ctx) {
		var country Country

		results.Decode(&country)

		countries = append(countries, country)
	}

	return c.Status(fiber.StatusOK).JSON(countries)
}

func main() {
	app := fiber.New()

	ConnectDB()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Countries API")
	})

	app.Get("/countries/:name?", countries)

	app.Listen(":3000")
}
