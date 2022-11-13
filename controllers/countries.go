package controllers

import (
	"context"
	"countries_api/helpers"
	"countries_api/models"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func Countries(c *fiber.Ctx) error {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	countryNameToSearch := c.Query("name")

	if countryNameToSearch == "" {
		results, _ := CountryCollection.Find(ctx, bson.M{})

		var countries []models.Country

		defer results.Close(ctx)

		for results.Next(ctx) {
			var country models.Country

			results.Decode(&country)

			countries = append(countries, country)
		}

		return helpers.ProvideResponse(c, fiber.StatusOK, "Success", countries)
	} else {
		var countryFound models.Country

		results, _ := CountryCollection.Find(ctx, bson.M{})

		for results.Next(ctx) {
			var country models.Country

			results.Decode(&country)

			if strings.ToLower(strings.Trim(countryNameToSearch, " ")) == strings.ToLower(country.Name) {
				countryFound = country

				break
			}
		}

		if (countryFound == models.Country{}) {
			return helpers.ProvideResponse(c, fiber.StatusNotFound, "Failed, country with name '"+countryNameToSearch+"' doesn't exist", bson.M{})
		} else {
			return helpers.ProvideResponse(c, fiber.StatusOK, "Success", countryFound)
		}
	}
}
