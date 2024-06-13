package models

import (
	"context"

	"github.com/gofiber/fiber/v2/log"
	myerrors "github.com/islem143/go-chat/Errors"
	"github.com/islem143/go-chat/database"
	"github.com/islem143/go-chat/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func FindAll(collection string, results types.Document, filter bson.D, opts *options.FindOptions) error {

	coll := database.Client.Database(database.Database).Collection(collection)
	cursor, err := coll.Find(context.TODO(), filter, opts)
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
func FindOne(collection string, filter bson.M, result Model) error {

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

func Insert(collection string, model Model) error {
	coll := database.Client.Database(database.Database).Collection(collection)
	model.SetCreatedAt()
	model.SetUpdatedAt()
	data := any(model)
	_, err := coll.InsertOne(context.TODO(), data)

	if err != nil {
		log.Error(err)
		return &myerrors.MyError{Code: 500, Message: "internal server error"}
	}
	return nil
}
func Update(collection string, filter bson.M, update primitive.D) error {
	opts := options.Update().SetUpsert(true)
	coll := database.Client.Database(database.Database).Collection(collection)

	_, err := coll.UpdateOne(context.TODO(), filter, update, opts)
	if err != nil {
		return err
	}
	return nil
}
