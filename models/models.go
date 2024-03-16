package models

import (
	"context"
	"log"

	"github.com/islem143/go-chat/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Model interface {
	Collection() string
}

type User struct {
	ID    string `json:"id" bson:"_id,omitempty"`
	Name  string `json:"name" bson:"name,omitempty"`
	Email string `json:"email" bson:"email,omitempty"`
	//Password []byte `json:"-"`
}

func (user *User) Collection() string {
	return "users"
}

func FindUserByEmail(email string) (*User, error) {
	filter := bson.D{{Key: "email", Value: email}}
	user := &User{}
	err := FindOne(user.Collection(), filter, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func InsertUser(user *User) error {
	err := Insert(user.Collection(), user)
	if err != nil {
		return err
	}
	return nil
}

// func (user *User) FindOne(filter bson.D) []User {
// 	cursor := FindOne(user.Collection(), filter)
// 	var users []User
// 	if err := cursor.All(context.TODO(), &users); err != nil {
// 		panic(err)
// 	}
// 	for _, result := range users {
// 		res, _ := bson.MarshalExtJSON(result, false, false)
// 		fmt.Println(string(res))
// 	}
// 	return users

// }
func SetUp() {
	coll := database.Client.Database(database.Database).Collection("users")
	indexModel := mongo.IndexModel{
		Keys:    bson.M{"email": 2},
		Options: options.Index().SetUnique(true),
	}
	indexModel2 := mongo.IndexModel{
		Keys:    bson.M{"name": 2},
		Options: options.Index().SetUnique(true),
	}

	// Create the index
	_, err := coll.Indexes().CreateOne(context.Background(), indexModel)
	if err != nil {
		log.Fatal(err)
	}
	_, err = coll.Indexes().CreateOne(context.Background(), indexModel2)
	if err != nil {
		log.Fatal(err)
	}
}
