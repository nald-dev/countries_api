package main

import (
	"countries_api/configs"
	"countries_api/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	configs.ConnectDB()

	routes.MainRoute(app)

	app.Listen(":3000")
}
