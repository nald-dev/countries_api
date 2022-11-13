package helpers

import (
	"countries_api/models"

	"github.com/gofiber/fiber/v2"
)

func ProvideResponse(c *fiber.Ctx, status int, message string, data interface{}) error {
	return c.Status(status).JSON(models.Response{
		Status:  status,
		Message: message,
		Data:    data,
	})
}
