package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Country struct {
	Id          primitive.ObjectID `json:"id,omitempty"`
	Name        string             `json:"name,omitempty" validate:"required"`
	CapitalCity string             `json:"capitalCity,omitempty" validate:"required"`
	Currency    string             `json:"currency,omitempty" validate:"required"`
}
