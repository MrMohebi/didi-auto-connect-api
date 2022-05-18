package main

import (
	"github.com/MrMohebi/didi-auto-connect-api.git/router"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	router.Routs(server)

	err := server.Run(":8005")
	if err != nil {
		return
	}
}
