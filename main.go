package main

import (
	"countries_api/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	routes.MainRoute(app)

	app.Listen(":3000")
}
