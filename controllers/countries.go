package controllers

import (
	"context"
	"countries_api/models"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

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

	return c.Status(fiber.StatusOK).JSON(models.Response{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    bson.M{"items": countries},
	})
}
