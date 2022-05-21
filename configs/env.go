package configs

import (
	"github.com/MrMohebi/didi-auto-connect-api.git/common"
	"github.com/joho/godotenv"
	"os"
)

func EnvSetup() {
	err := godotenv.Load()
	common.IsErr(err, "Error loading .env file")
}

func EnvMongoURI() string {
	return os.Getenv("MONGO_FINAL_URI")
}

func EvnMongoDB() string {
	return os.Getenv("MONGO_DB")
}
