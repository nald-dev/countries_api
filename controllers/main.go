package controllers

import (
	"countries_api/configs"

	"go.mongodb.org/mongo-driver/mongo"
)

var CountryCollection *mongo.Collection = configs.GetCollection("countries")
var UserCollection *mongo.Collection = configs.GetCollection("users")
