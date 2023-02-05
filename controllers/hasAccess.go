package controllers

import (
	"context"
	"github.com/MrMohebi/didi-auto-connect-api.git/faces"
	"github.com/MrMohebi/didi-auto-connect-api.git/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"time"
)

func HasAccess() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		now := time.Now()

		deviceHash, isOkay := c.GetQuery("deviceHash")
		if !isOkay {
			c.JSON(http.StatusBadRequest, 400)
			return
		}

		var message = ""
		var hasAccess = true

		user, isOkay := models.UserWithToken(c)
		if !isOkay {
			return
		}

		devicesCursor, _ := models.DevicesCollection.Find(ctx, bson.M{"userID": user.Id})
		if devicesCursor.RemainingBatchLength() > int(user.DeviceLimitation) {
			hasAccess = false
			message = "شما از حداکثر تعداد ممکن دستگاه متصل استفاده کرده اید. برای استفاده مجدد 48 ساعت صبر کنید تا دستگاه های قبلی به صورت خودکار از سیستم حذف شوند"
		}

		if user.ActiveTill < now.Unix() {
			hasAccess = false
			message = "مدت اعتبار حساب شما تمام شده است. لطفا اشتراک خود را تمدید کنید"
		}

		deviceNotfound := models.DevicesCollection.FindOne(ctx, bson.M{"hash": deviceHash, "userID": user.Id})
		if deviceNotfound != nil {
			hasAccess = false
			message = "دستگاهی با این ایدی یافت نشد!"
		}

		c.JSON(http.StatusOK, faces.HasAccessRes{
			HasAccess: hasAccess,
			Message:   message,
		})
	}
}
