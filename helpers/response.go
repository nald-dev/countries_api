package helpers

import (
	"countries_api/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ProvideResponse(c *fiber.Ctx, status int, message string, data primitive.M) error {
	return c.Status(status).JSON(models.Response{
		Status:  status,
		Message: message,
		Data:    data,
	})
}
