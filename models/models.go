package models

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Model interface {
	Collection() string
}

type User struct {
	ID       string `json:"id" bson:"_id,omitempty"`
	Name     string `json:"name" bson:"name,omitempty"`
	Email    string `json:"email" bson:"email,omitempty"`
	Password []byte `json:"-"`
}

func (user *User) Collection() string {
	return "users"
}

func FindUserById(value string) (*User, error) {
	objectId, err := primitive.ObjectIDFromHex(value)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": objectId}
	user := &User{}
	err = FindOne(user.Collection(), filter, user)
	if err != nil {

		return nil, err
	}
	return user, nil
}
func FindUser(key string, value string) (*User, error) {
	fmt.Println(key, value)
	filter := bson.M{key: value}
	user := &User{}
	err := FindOne(user.Collection(), filter, user)
	if err != nil {

		return nil, err
	}
	return user, nil
}

func FindAllUsers() (*[]User, error) {

	users := &[]User{}
	err := FindAll("users", users)
	if err != nil {

		return nil, err
	}
	return users, nil
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
