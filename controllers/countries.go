package controllers

import (
	"context"
	"countries_api/helpers"
	"countries_api/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func Countries(c *fiber.Ctx) error {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	results, _ := CountryCollection.Find(ctx, bson.M{})

	var countries []models.Country

	defer results.Close(ctx)

	for results.Next(ctx) {
		var country models.Country

		results.Decode(&country)

		countries = append(countries, country)
	}

	return helpers.ProvideResponse(c, fiber.StatusOK, "Success", bson.M{"items": countries})
}
