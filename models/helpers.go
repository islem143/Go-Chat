package models

import (
	"context"

	"github.com/gofiber/fiber/v2/log"
	myerrors "github.com/islem143/go-chat/Errors"
	"github.com/islem143/go-chat/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Document interface{}

func FindAll(collection string, results Document, filter bson.M) error {

	coll := database.Client.Database(database.Database).Collection(collection)
	cursor, err := coll.Find(context.TODO(), bson.D{})
	if err != nil {
		if err == mongo.ErrNoDocuments {
			results = nil
			return err
		}
		log.Error(err)
		return &myerrors.MyError{Code: 500, Message: "internal server error"}

	}

	if err = cursor.All(context.TODO(), results); err != nil {
		panic(err)
	}

	return nil
}
func FindOne(collection string, filter bson.M, result Document) error {

	coll := database.Client.Database(database.Database).Collection(collection)
	err := coll.FindOne(context.TODO(), filter).Decode(result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			result = nil
			return err
		}
		log.Error(err)
		return &myerrors.MyError{Code: 500, Message: "internal server error"}

	}

	return nil
}

func Insert(collection string, data Document) error {
	coll := database.Client.Database(database.Database).Collection(collection)

	_, err := coll.InsertOne(context.TODO(), data)

	if err != nil {
		log.Error(err)
		return &myerrors.MyError{Code: 500, Message: "internal server error"}
	}
	return nil
}
