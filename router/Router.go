package router

import (
	"github.com/MrMohebi/didi-auto-connect-api.git/contorolers"
	"github.com/gin-gonic/gin"
)

func Routs(r *gin.Engine) {
	r.GET("/test", contorolers.Test)
}
