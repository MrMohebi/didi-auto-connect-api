package controllers

import (
	"context"
	"fmt"
	"github.com/MrMohebi/didi-auto-connect-api.git/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"
)

func IsRegistered() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		username, isOkay := c.GetQuery("username")
		if !isOkay {
			c.JSON(http.StatusBadRequest, 400)
			return
		}

		var user models.User

		userNotFound := models.UsersCollection.FindOne(ctx, bson.M{"username": bson.M{"$regex": primitive.Regex{Pattern: fmt.Sprintf("^%s$", username), Options: "i"}}}).Decode(&user)
		if userNotFound != nil {
			c.JSON(http.StatusNotFound, false)
			return
		}

		c.JSON(http.StatusOK, true)
	}
}
