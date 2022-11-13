package controllers

import (
	"context"
	"countries_api/helpers"
	"countries_api/models"
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

	UserCollection.FindOne(ctx, bson.M{"username": user.Username}).Decode(&user)

	return helpers.ProvideResponse(c, fiber.StatusOK, "Success", bson.M{
		"id":       newUser.Id,
		"username": newUser.Username,
	})
}

func Login(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	var user models.User
	defer cancel()

	c.BodyParser(&user)

	if user.Username == "" {
		return helpers.ProvideResponse(c, fiber.StatusBadRequest, "Failed, please provide the 'username'", bson.M{})
	}

	if user.Password == "" {
		return helpers.ProvideResponse(c, fiber.StatusBadRequest, "Failed, please provide the 'password'", bson.M{})
	}

	var userFound models.User

	UserCollection.FindOne(ctx, bson.M{"username": user.Username}).Decode(&userFound)

	if (userFound == models.User{} || !helpers.CheckPasswordHash(user.Password, userFound.Password)) {
		return helpers.ProvideResponse(c, fiber.StatusBadRequest, "Failed, wrong username or password", bson.M{})
	} else {
		return helpers.ProvideResponse(c, fiber.StatusOK, "Success", bson.M{
			"id":       userFound.Id,
			"username": userFound.Username,
		})
	}
}
