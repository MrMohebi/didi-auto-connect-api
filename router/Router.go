package router

import (
	"github.com/MrMohebi/didi-auto-connect-api.git/contorolers"
	"github.com/gin-gonic/gin"
)

func Routs(r *gin.Engine) {
	AssetsRoute(r)

	r.LoadHTMLGlob("templates/**/*.html")

	r.GET("/test", contorolers.Test)

	r.GET("/", contorolers.Index)
	r.GET("/docs", contorolers.Docs)
}
