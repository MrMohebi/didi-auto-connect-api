package main

import (
	"github.com/MrMohebi/didi-auto-connect-api.git/common"
	"github.com/MrMohebi/didi-auto-connect-api.git/configs"
	"github.com/MrMohebi/didi-auto-connect-api.git/router"
	"github.com/gin-gonic/gin"
)

// nodemon --exec go run main.go --signal SIGTERM

func main() {
	configs.Setup()

	server := gin.Default()

	router.Routs(server)

	err := server.Run("localhost:8005")
	common.IsErr(err, "Err in starting server")
}
