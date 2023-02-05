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

func DidiAccountGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		user, isOkay := models.UserWithToken(c)
		if !isOkay {
			return
		}

		var didis []models.DidiAccount
		cur, err := models.DidiAccountsCollection.Find(ctx, bson.M{"userID": user.Id})
		defer cur.Close(ctx)
		if err != nil {
			c.JSON(http.StatusOK, nil)
		}

		for cur.Next(ctx) {
			var elem models.DidiAccount
			common.IsErr(cur.Decode(&elem))
			didis = append(didis, elem)
		}

		var result []faces.DidiAccountGetRes
		for _, didi := range didis {
			result = append(result, faces.DidiAccountGetRes{
				Id:        didi.Id,
				Username:  didi.Username,
				Password:  didi.Password,
				UpdatedAt: didi.UpdatedAt,
				CreatedAt: didi.CreatedAt,
			})
		}

		c.JSON(http.StatusOK, result)

	}
}

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

		id, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, false)
			return
		}

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

		_, err = models.DidiAccountsCollection.DeleteOne(ctx, filter)
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
