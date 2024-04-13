package connection

import (
	"context"
	"fmt"
	"time"

	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DatabaseCollection struct {
	UserInDB *mongo.Collection
}

var CollectionsForDB *DatabaseCollection
var Client *mongo.Client = DbConnection()

func DbConnection() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Error().Msg("Error in connection with DB")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Error().Msg("Error while connectinfg to database....")
	}
	fmt.Println("Connected to mongo database successfully...")
	return client

}

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("fiberApi").Collection(collectionName)
	return collection

}
