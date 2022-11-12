package main

import "github.com/gofiber/fiber/v2"

type Error struct {
	Message string
}
type Country struct {
	Name string
	CapitalCity string
	CurrencyName string
}

func getAllCountries() []Country {
	countries:= []Country{
		{"Indonesia","Jakarta","Rupiah"},
		{"Malaysia","Kuala Lumpur","Ringgit"},
		{"India","New Delhi","Rupee"},
	}

	return countries
}

func countries(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(getAllCountries())
}

func searchCountry(ctx *fiber.Ctx) error {
	var country Country

	allCountries:= getAllCountries()

	nameQuery:= ctx.Params("name")

	for i := range allCountries {
		if allCountries[i].Name == nameQuery {
			country = allCountries[i]

			break
		}
	}
	
	if (country == Country{}) {
		return ctx.Status(fiber.StatusNotFound).JSON(Error {
			"Country with name " + nameQuery + " not found" ,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(country)
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Countries API")
	})

	app.Get("/countries/:name?", func(c *fiber.Ctx) error {
		if (c.Params(("name")) != "") {
			return searchCountry(c)
		} else {
			return countries(c)
		}
	})

	app.Listen(":3000")
}