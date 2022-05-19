package contorolers

import (
	"github.com/gin-gonic/gin"
)

//var validate = validator.New()

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		println(c.PostFormMap("username"))
		//var reqBody faces.LoginReq
		//
		//err := c.BindJSON(&reqBody)
		//common.IsErr(err)

		//ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		//defer cancel()

		//ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		//var user models.User
		//defer cancel()
		//
		////validate the request body
		//if err := c.BindJSON(&user); err != nil {
		//	c.JSON(http.StatusBadRequest, faces.LoginRes{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		//	return
		//}
	}
}
