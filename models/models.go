package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id    primitive.ObjectID `bson:"_id"`
	Name  string             `json:"name"`
	Email string             `json:"email"`
	Age   int                `json:"age"`
}
