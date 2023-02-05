package controllers

import (
	"context"
	"github.com/MrMohebi/didi-auto-connect-api.git/common"
	"github.com/MrMohebi/didi-auto-connect-api.git/faces"
	"github.com/MrMohebi/didi-auto-connect-api.git/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"
)

func DidiAccountCreate() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var reqBody faces.DidiAccountCreateReq
		common.ValidBindForm(c, &reqBody)

		now := time.Now()

		user, isOkay := models.UserWithToken(c)
		if !isOkay {
			return
		}

		var didi models.DidiAccount
		err := models.DidiAccountsCollection.FindOne(ctx, bson.M{"username": reqBody.Username, "userID": user.Id}).Decode(&didi)
		if err != nil {
			_, err := models.DidiAccountsCollection.InsertOne(
				ctx,
				bson.D{
					{"userID", user.Id},
					{"username", reqBody.Username},
					{"password", reqBody.Password},
					{"createdAt", now.Unix()},
				},
			)
			common.IsErr(err)

			c.JSON(http.StatusOK, true)
			return
		}

		c.JSON(http.StatusConflict, false)
		return
	}
}

func DidiAccountDelete() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		username, isOkay := c.GetQuery("username")

		user, isOkay := models.UserWithToken(c)
		if !isOkay {
			return
		}

		var didi models.DidiAccount
		err := models.DidiAccountsCollection.FindOne(ctx, bson.M{"username": username, "userID": user.Id}).Decode(&didi)
		if err != nil {
			c.JSON(http.StatusNotFound, false)
			return
		}

		_, err = models.DidiAccountsCollection.DeleteOne(
			ctx,
			bson.D{
				{"userID", user.Id},
				{"username", username},
			},
		)
		common.IsErr(err)

		c.JSON(http.StatusOK, true)
		return
	}
}

func DidiAccountModify() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var reqBody faces.DidiAccountModifyReq
		common.ValidBindForm(c, &reqBody)

		id, err := primitive.ObjectIDFromHex(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, false)
			return
		}

		now := time.Now()

		user, isOkay := models.UserWithToken(c)
		if !isOkay {
			return
		}

		filter := bson.M{"_id": id, "userID": user.Id}

		var didi models.DidiAccount
		err = models.DidiAccountsCollection.FindOne(ctx, filter).Decode(&didi)
		if err != nil {
			c.JSON(http.StatusNotFound, false)
			return
		}

		_, err = models.DidiAccountsCollection.UpdateOne(
			ctx,
			filter,
			bson.M{
				"$set": bson.M{
					"username":  reqBody.Username,
					"password":  reqBody.Password,
					"updatedAt": now.Unix(),
				},
			},
		)
		common.IsErr(err)

		c.JSON(http.StatusOK, true)
		return
	}
}
