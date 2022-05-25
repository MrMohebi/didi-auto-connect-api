package models

import (
	"github.com/MrMohebi/didi-auto-connect-api.git/configs"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Device struct {
	Id        primitive.ObjectID `json:"id" bson:"_id"`
	UserID    primitive.ObjectID `json:"userID" validate:"required"`
	Hash      string             `json:"hash" validate:"required"`
	IsActive  bool               `json:"isActive" validate:"required"`
	LastLogin int64              `json:"lastLogin" `
	CreatedAt int64              `json:"createdAt" validate:"required"`
}

var DevicesCollection *mongo.Collection = configs.GetCollection(configs.GetDBClint(), "devices")
