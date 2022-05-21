package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id        primitive.ObjectID `json:"id"`
	Username  string             `json:"username" validate:"required"`
	Password  string             `json:"password" validate:"required"`
	LastLogin int32              `json:"lastLogin"`
	CreatedAt int32              `json:"createdAt" validate:"required"`
}

//var UsersCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")
