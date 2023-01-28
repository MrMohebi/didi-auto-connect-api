package models

import (
	"github.com/MrMohebi/didi-auto-connect-api.git/configs"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	Id               primitive.ObjectID `json:"id" bson:"_id"`
	Username         string             `json:"username" validate:"required"`
	Password         string             `json:"password" validate:"required"`
	Token            string             `json:"token" validate:"required"`
	DeviceLimitation int8               `json:"deviceLimitation"`
	ActiveTill       int64              `json:"activeTill"`
	LastLogin        int64              `json:"lastLogin"`
	CreatedAt        int64              `json:"createdAt" validate:"required"`
}

var UsersCollection *mongo.Collection = configs.GetCollection(configs.GetDBClint(), "users")
