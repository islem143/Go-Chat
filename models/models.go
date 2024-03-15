package models

import (
	"context"
	"log"

	"github.com/islem143/go-chat/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	Name  string `bson:"name,omitempty"`
	Email string `bson:"email,omitempty"`
	//Password []byte `json:"-"`
}

func SetUp() {
	coll := database.Client.Database(database.Database).Collection("users")
	indexModel := mongo.IndexModel{
		Keys:    bson.M{"email": 1},
		Options: options.Index().SetUnique(true),
	}

	// Create the index
	_, err := coll.Indexes().CreateOne(context.Background(), indexModel)
	if err != nil {
		log.Fatal(err)
	}
}
