package contorolers

import (
	"github.com/MrMohebi/didi-auto-connect-api.git/common"
	"github.com/MrMohebi/didi-auto-connect-api.git/faces"
	"github.com/MrMohebi/didi-auto-connect-api.git/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		//ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		//defer cancel()

		var reqBody faces.LoginReq
		common.ValidBindForm(c, &reqBody)

		var user models.User

		isLimit := false
		token := common.RandStr(16)

		//err := models.UsersCollection.FindOne(ctx, bson.M{"$regex": reqBody.Username, "$options": "i"}).Decode(&user)
		//common.IsErr(err)

		c.JSON(http.StatusOK, gin.H{
			"asdasd":  isLimit,
			"aaaaa":   token,
			"usename": user.Username,
		})

	}
}

func singUp() {

}

func login() {
}
