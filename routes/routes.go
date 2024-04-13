package routes

import (
	"context"
	"fiberMongo/connection"
	"fiberMongo/models"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = connection.OpenCollection(connection.Client, "user")

func AddUser(c *fiber.Ctx) error {
	user := new(models.User)
	c.BodyParser(user)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	user.Id = primitive.NewObjectID()
	insertResult, err := userCollection.InsertOne(ctx, user)
	if err != nil {
		log.Fatal("Insert Result error .....", err)
	}
	return c.Status(200).JSON(insertResult)
}
