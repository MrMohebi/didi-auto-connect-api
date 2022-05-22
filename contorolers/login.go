package contorolers

import (
	"context"
	"github.com/MrMohebi/didi-auto-connect-api.git/common"
	"github.com/MrMohebi/didi-auto-connect-api.git/faces"
	"github.com/MrMohebi/didi-auto-connect-api.git/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
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

		isJoined := false
		err := models.UsersCollection.FindOne(ctx, bson.M{"username": bson.M{"$regex": reqBody.Username, "$options": "i"}}).Decode(&user)
		if err != nil {
			isJoined = true
		}

		if !isJoined {
			singUp()
			return
		}

		login()

		c.JSON(http.StatusOK, gin.H{
			"token":     token,
			"hasAccess": !isLimit,
			"isLimit":   isLimit,
		})
	}
}

func singUp() {

}

func login() {
}
