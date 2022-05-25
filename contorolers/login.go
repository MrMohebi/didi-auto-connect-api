package contorolers

import (
	"context"
	"fmt"
	"github.com/MrMohebi/didi-auto-connect-api.git/common"
	"github.com/MrMohebi/didi-auto-connect-api.git/faces"
	"github.com/MrMohebi/didi-auto-connect-api.git/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"
)

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var reqBody faces.LoginReq
		common.ValidBindForm(c, &reqBody)

		var user models.User

		isLimit := false
		token := common.RandStr(32)

		isJoined := true
		userNotFound := models.UsersCollection.FindOne(ctx, bson.M{"username": bson.M{"$regex": primitive.Regex{Pattern: fmt.Sprintf("^%s$", reqBody.Username), Options: "i"}}}).Decode(&user)
		if userNotFound != nil {
			isJoined = false
		}

		if !isJoined {
			singUp(&reqBody, &token)
		}

		if isJoined && !login(&reqBody, &user, &token, &isLimit) {
			c.JSON(http.StatusUnauthorized, 401)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"token":     token,
			"hasAccess": !isLimit,
			"isLimit":   isLimit,
		})
	}
}

func login(reqBody *faces.LoginReq, user *models.User, token *string, isLimit *bool) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if !common.VerifyPasswordHash(reqBody.Password, user.Password) {
		return false
	}

	now := time.Now()
	_, err := models.UsersCollection.UpdateOne(
		ctx,
		bson.D{{"_id", user.Id}},
		bson.D{{"$set", bson.D{
			{"token", token},
			{"lastLogin", now.Unix()},
		}}},
	)
	common.IsErr(err)

	var device models.Device
	isDeviceRegistered := true
	deviceNotfound := models.DevicesCollection.FindOne(ctx, bson.M{"hash": reqBody.DeviceHash, "userID": user.Id}).Decode(&device)
	if deviceNotfound != nil {
		isDeviceRegistered = false
	}

	//create new device
	if !isDeviceRegistered {
		resultDevice, err := models.DevicesCollection.InsertOne(
			ctx,
			bson.D{
				{"userID", user.Id},
				{"hash", reqBody.DeviceHash},
				{"isActive", false},
				{"lastLogin", now.Unix()},
				{"createdAt", now.Unix()},
			},
		)
		common.IsErr(err)
		_ = models.DevicesCollection.FindOne(ctx, bson.M{"_id": resultDevice.InsertedID}).Decode(&device)
	}
	println(device.Hash)
	println(device.IsActive)
	var activeDevice models.Device
	// check if it has limitation
	if !device.IsActive {
		var limitationTime int64 = 5 * 60 * 60
		if err := models.DevicesCollection.FindOne(ctx, bson.M{"userID": user.Id, "isActive": true}).Decode(&activeDevice); err == nil {
			println("active found")
			if activeDevice.LastLogin > (now.Unix() - limitationTime) {
				*isLimit = true
			}
		}

		println(activeDevice.Hash)
		_, _ = models.DevicesCollection.UpdateOne(
			ctx,
			bson.M{"userID": user.Id, "isActive": true},
			bson.D{{"$set", bson.D{
				{"isActive", false},
			}}},
		)
		_, _ = models.DevicesCollection.UpdateOne(
			ctx,
			bson.M{"userID": user.Id, "hash": device.Hash},
			bson.D{{"$set", bson.D{
				{"isActive", true},
				{"lastLogin", now.Unix()},
			}}},
		)
	}
	return true
}

func singUp(reqBody *faces.LoginReq, token *string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	password, _ := common.HashPassword(reqBody.Password)

	now := time.Now()
	resultUser, err := models.UsersCollection.InsertOne(
		ctx,
		bson.D{
			{"username", reqBody.Username},
			{"password", password},
			{"token", token},
			{"lastLogin", now.Unix()},
			{"createdAt", now.Unix()},
		},
	)
	_, err = models.DevicesCollection.InsertOne(
		ctx,
		bson.D{
			{"userID", resultUser.InsertedID},
			{"hash", reqBody.DeviceHash},
			{"isActive", true},
			{"lastLogin", now.Unix()},
			{"createdAt", now.Unix()},
		},
	)
	common.IsErr(err)
}
