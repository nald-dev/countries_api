package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type User struct {
	Id       primitive.ObjectID `json:"_id" bson:"_id"`
	Username string             `json:"username,omitempty" validate:"required"`
	Password string             `json:"password,omitempty" validate:"required"`
}

type Country struct {
	Id          primitive.ObjectID `json:"_id" bson:"_id"`
	Name        string             `json:"name,omitempty" validate:"required"`
	CapitalCity string             `json:"capitalCity,omitempty" validate:"required"`
	Currency    string             `json:"currency,omitempty" validate:"required"`
}
