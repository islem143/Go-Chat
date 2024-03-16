package models

import (
	"context"
	"fmt"

	myerrors "github.com/islem143/go-chat/Errors"
	"github.com/islem143/go-chat/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Document interface{}

func FindOne(collection string, filter bson.D, result Document) error {

	coll := database.Client.Database(database.Database).Collection(collection)
	err := coll.FindOne(context.TODO(), filter).Decode(result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return myerrors.NotFoundError("document not Found")

		}
		return myerrors.ErrInternalServerError

	}

	return nil
}

func Insert(collection string, data Document) error {
	coll := database.Client.Database(database.Database).Collection(collection)

	_, err := coll.InsertOne(context.TODO(), data)
	if err != nil {
		fmt.Println(err)
		return myerrors.InternalServerError("internal server error")
	}
	return nil
}
