package routes

import (
	"countries_api/controllers"

	"github.com/gofiber/fiber/v2"
)

func MainRoute(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Countries API")
	})

	app.Get("/countries/:name?", controllers.Countries)
}
