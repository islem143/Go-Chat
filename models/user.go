package models

import (
	myerrors "github.com/islem143/go-chat/Errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	BaseModel `bson:",inline"`
	ID        string   `json:"id" bson:"_id,omitempty"`
	Name      string   `json:"name" bson:"name,omitempty"`
	Email     string   `json:"email" bson:"email,omitempty"`
	Picture   string   `json:"picture" bson:"picture,omitempty"`
	Contacts  []string `json:"contacts" bson:"contacts,omitempty"`
	//Role     string `json:"role" bson:"role,omitempty"`
	Password []byte `json:"-"`
}

func (user *User) CollectionName() string {
	return "users"
}
func FindUserById(value string) (*User, error) {
	objectId, err := primitive.ObjectIDFromHex(value)
	if err != nil {
		return nil, myerrors.DbErrors(err)
	}
	filter := bson.M{"_id": objectId}
	user := &User{}
	err = FindOne(user.CollectionName(), filter, user)
	if err != nil {

		return nil, myerrors.DbErrors(err)
	}
	return user, nil

}
func FindUser(key string, value string) (*User, error) {

	filter := bson.M{key: value}
	user := &User{}
	err := FindOne(user.CollectionName(), filter, user)
	if err != nil {
		return nil, myerrors.DbErrors(err)

	}
	return user, nil
}

func FindAllUsers(filter bson.D) (*[]User, error) {

	users := &[]User{}

	err := FindAll("users", users, filter, nil)
	if err != nil {
		return nil, myerrors.DbErrors(err)

	}
	return users, nil
}
func InsertUser(user *User) error {
	err := Insert(user.CollectionName(), user)
	if err != nil {

		return myerrors.DbErrors(err)
	}
	return nil
}

func UpdateUser(user *User, filter bson.M, update bson.D) error {

	err := Update(user.CollectionName(), filter, update)
	if err != nil {

		return myerrors.DbErrors(err)
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
