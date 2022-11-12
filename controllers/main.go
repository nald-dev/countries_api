package controllers

import (
	"context"
	"countries_api/configs"
	"countries_api/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var CountryCollection *mongo.Collection = configs.GetCollection(configs.DB, "countries")

func Countries(c *fiber.Ctx) error {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	results, _ := CountryCollection.Find(ctx, bson.M{})

	var countries []models.Country

	// reading from the db in an optimal way
	defer results.Close(ctx)

	for results.Next(ctx) {
		var country models.Country

		results.Decode(&country)

		countries = append(countries, country)
	}

	return c.Status(fiber.StatusOK).JSON(countries)
}