package routes

import (
	"countries_api/controllers"

	"github.com/gofiber/fiber/v2"
)

func MainRoute(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Countries API")
	})

	app.Post("/register", controllers.CreateUser)
	app.Post("/login", controllers.Login)

	app.Get("/countries/:name?", controllers.Countries)
}
