package models

import (
	"github.com/MrMohebi/didi-auto-connect-api.git/configs"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type DidiAccount struct {
	Id        primitive.ObjectID `json:"id" bson:"_id"`
	UserID    primitive.ObjectID `json:"userID" validate:"required"`
	Username  string             `json:"username" validate:"required"`
	Password  string             `json:"password" validate:"required"`
	IsDefault bool               `json:"isDefault" `
	LastLogin int64              `json:"lastLogin" `
	CreatedAt int64              `json:"createdAt" validate:"required"`
	UpdatedAt int64              `json:"updatedAt"`
}

var DidiAccountsCollection *mongo.Collection = configs.GetCollection(configs.GetDBClint(), "didiAccounts")
