package models

import (
	"context"
	"github.com/MrMohebi/didi-auto-connect-api.git/configs"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"time"
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

// UserWithToken check if user has token and its valid
func UserWithToken(c *gin.Context) (user User, isOkay bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	token := c.GetHeader("token")

	userNotFound := UsersCollection.FindOne(ctx, bson.M{"token": token}).Decode(&user)
	if userNotFound != nil {
		c.JSON(http.StatusUnauthorized, 401)
	}

	isOkay = user != (User{})

	return user, isOkay

}
