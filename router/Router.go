package router

import (
	"github.com/MrMohebi/didi-auto-connect-api.git/controllers"
	"github.com/gin-gonic/gin"
)

func Routs(r *gin.Engine) {
	AssetsRoute(r)
	r.LoadHTMLGlob("templates/**/*.html")

	r.GET("/", controllers.Index())
	r.GET("/docs", controllers.Docs())

	api := r.Group("/api")
	{
		api.POST("/login", controllers.Login())
		api.GET("/hasAccess", controllers.HasAccess())
		api.GET("/isRegistered", controllers.IsRegistered())

		didiAccount := api.Group("/didiAccount")
		{
			didiAccount.POST("/create", controllers.DidiAccountCreate())
			didiAccount.POST("/modify/:id", controllers.DidiAccountModify())
			didiAccount.POST("/delete/:id", controllers.DidiAccountDelete())
			didiAccount.GET("/get", controllers.DidiAccountGet())
		}

	}
}
