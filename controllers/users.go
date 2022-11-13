package controllers

import (
	"context"
	"countries_api/helpers"
	"countries_api/models"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
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

	return c.Status(fiber.StatusOK).JSON(models.Response{
		Status:  http.StatusOK,
		Message: "Success",
		Data: bson.M{
			"id":       newUser.Id,
			"username": newUser.Username,
		},
	})
}

func Login(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	var user models.User
	defer cancel()

	c.BodyParser(&user)

	var userFound models.User

	UserCollection.FindOne(ctx, bson.M{"username": user.Username}).Decode(&userFound)

	if helpers.CheckPasswordHash(user.Password, userFound.Password) {
		return c.Status(fiber.StatusOK).JSON(models.Response{
			Status:  http.StatusOK,
			Message: "Success",
			Data: bson.M{
				"id":       userFound.Id,
				"username": userFound.Username,
			},
		})
	} else {
		return c.Status(fiber.StatusNotFound).JSON(models.Response{
			Status:  http.StatusNotFound,
			Message: "Failure",
			Data:    bson.M{},
		})
	}
}
