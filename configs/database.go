package configs

import (
	"context"
	"fmt"
	"github.com/MrMohebi/didi-auto-connect-api.git/common"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func ConnectDB() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(EnvMongoURI()))
	common.IsErr(err)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	common.IsErr(err)

	//ping the database
	err = client.Ping(ctx, nil)
	common.IsErr(err)

	fmt.Println("Connected to MongoDB")
	return client
}

// DB Client instance
var DB *mongo.Client = ConnectDB()

// GetCollection getting database collections
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database(EvnMongoDB()).Collection(collectionName)
	return collection
}
