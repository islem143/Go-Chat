package models

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	BaseModel `bson:",inline"`
	ID        string `json:"id" bson:"_id,omitempty"`
	Name      string `json:"name" bson:"name,omitempty"`
	Email     string `json:"email" bson:"email,omitempty"`
	Picture   string `json:"picture" bson:"picture,omitempty"`
	//Role     string `json:"role" bson:"role,omitempty"`
	Password []byte `json:"-"`
}

func (user *User) CollectionName() string {
	return "users"
}
func FindUserById(value string) (*User, error) {
	objectId, err := primitive.ObjectIDFromHex(value)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": objectId}
	user := &User{}
	err = FindOne(user.CollectionName(), filter, user)
	if err != nil {

		return nil, err
	}
	return user, nil
}
func FindUser(key string, value string) (*User, error) {

	filter := bson.M{key: value}
	user := &User{}
	err := FindOne(user.CollectionName(), filter, user)
	if err != nil {

		return nil, err
	}
	return user, nil
}

func FindAllUsers() (*[]User, error) {

	users := &[]User{}
	err := FindAll("users", users, nil)
	if err != nil {

		return nil, err
	}
	return users, nil
}
func InsertUser(user *User) error {
	err := Insert(user.CollectionName(), user)
	if err != nil {

		return err
	}
	return nil
}

func NewUser(name string, email string, password []byte) *User {
	model := &User{
		Name:     name,
		Email:    email,
		Password: password,
	}
	return model
}
