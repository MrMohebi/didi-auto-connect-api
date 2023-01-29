package controllers

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

		hasAccess := true
		message := ""

		isJoined := true
		userNotFound := models.UsersCollection.FindOne(ctx, bson.M{"username": bson.M{"$regex": primitive.Regex{Pattern: fmt.Sprintf("^%s$", reqBody.Username), Options: "i"}}}).Decode(&user)
		if userNotFound != nil {
			isJoined = false
		}

		if !isJoined {
			singUp(&reqBody)
			_ = models.UsersCollection.FindOne(ctx, bson.M{"username": bson.M{"$regex": primitive.Regex{Pattern: fmt.Sprintf("^%s$", reqBody.Username), Options: "i"}}}).Decode(&user)
		}

		if isJoined && !login(&reqBody, &user, &hasAccess, &message) {
			c.JSON(http.StatusUnauthorized, 401)
			return
		}

		c.JSON(http.StatusOK, faces.LoginRes{
			Token:     user.Token,
			HasAccess: hasAccess,
			Message:   message,
			Link:      "",
		})
	}
}

func login(reqBody *faces.LoginReq, user *models.User, hasAccess *bool, message *string) bool {
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
				{"lastLogin", now.Unix()},
				{"createdAt", now.Unix()},
			},
		)
		common.IsErr(err)
		_ = models.DevicesCollection.FindOne(ctx, bson.M{"_id": resultDevice.InsertedID}).Decode(&device)
	}

	if user.ActiveTill < now.Unix() {
		*hasAccess = false
		*message = "مدت اعتبار حساب شما تمام شده است. لطفا اکانت خود را تمدید کنید"
		return true
	}

	devicesCursor, _ := models.DevicesCollection.Find(ctx, bson.M{"userID": user.Id})

	if devicesCursor.RemainingBatchLength() > int(user.DeviceLimitation) {
		*hasAccess = false
		*message = "شما از حداکثر تعداد ممکن دستگاه متصل استفاده کرده اید. برای استفاده مجدد 48 ساعت صبر کنید تا دستگاه های قبلی به صورت خودکار از سیستم حذف شوند"
	}

	_, _ = models.DevicesCollection.UpdateOne(
		ctx,
		bson.M{"userID": user.Id, "hash": device.Hash},
		bson.D{{"$set", bson.D{
			{"lastLogin", now.Unix()},
		}}},
	)

	return true
}

func singUp(reqBody *faces.LoginReq) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	password, _ := common.HashPassword(reqBody.Password)
	token := common.RandStr(32)

	now := time.Now()
	resultUser, err := models.UsersCollection.InsertOne(
		ctx,
		bson.D{
			{"username", reqBody.Username},
			{"password", password},
			{"token", token},
			{"deviceLimitation", 2},
			{"lastLogin", now.Unix()},
			{"activeTill", now.Unix()},
			{"createdAt", now.Unix()},
		},
	)
	_, err = models.DevicesCollection.InsertOne(
		ctx,
		bson.D{
			{"userID", resultUser.InsertedID},
			{"hash", reqBody.DeviceHash},
			{"lastLogin", now.Unix()},
			{"createdAt", now.Unix()},
		},
	)
	common.IsErr(err)
}
