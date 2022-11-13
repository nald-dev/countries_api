package controllers

import (
	"context"
	"countries_api/helpers"
	"countries_api/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateUser(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	var user models.User
	defer cancel()

	c.BodyParser(&user)

	hashedPassword, _ := helpers.HashPassword(user.Password)

	newUser := models.User{
		Id:       primitive.NewObjectID(),
		Username: user.Username,
		Password: hashedPassword,
	}

	UserCollection.InsertOne(ctx, newUser)

	return c.Status(fiber.StatusOK).JSON(newUser)
}
