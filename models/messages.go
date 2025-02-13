package models

import (
	myerrors "github.com/islem143/go-chat/Errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	READ   = "read"
	UNREAD = "unread"
)

type Message struct {
	BaseModel  `bson:",inline"`
	ReceiverId string `json:"receiver_id" bson:"receiver_id,omitempty"`
	Status     string `json:"status" bson:"status,omitempty"`
	UserId     string `json:"user_id" bson:"user_id,omitempty"`
	Text       string `json:"text" bson:"text,omitempty"`
	Type       string `json:"type"`
}

func (user *Message) CollectionName() string {
	return "messages"
}

func FindMessages(filter bson.D, opts *options.FindOptions) (*[]Message, error) {

	messages := &[]Message{}
	err := FindAll("messages", messages, filter, opts)
	if err != nil {

		return nil, myerrors.DbErrors(err)
	}
	return messages, nil
}
func FindMessage(value string) (*Message, error) {
	objectId, err := primitive.ObjectIDFromHex(value)
	if err != nil {
		return nil, myerrors.DbErrors(err)
	}
	filter := bson.M{"_id": objectId}
	message := &Message{}
	err = FindOne(message.CollectionName(), filter, message)
	if err != nil {

		return nil, myerrors.DbErrors(err)
	}
	return message, nil
}

func InsertMessages(message *Message) error {
	err := Insert(message.CollectionName(), message)
	if err != nil {

		return myerrors.DbErrors(err)
	}
	return nil
}

func UpdateMessage(message *Message, feilds primitive.E) error {
	objectId, err := primitive.ObjectIDFromHex(message.ID)
	if err != nil {
		return myerrors.DbErrors(err)
	}

	update := bson.D{{Key: "$set", Value: bson.D{feilds}}}
	filter := bson.M{"_id": objectId}
	err = Update(message.CollectionName(), filter, update)
	if err != nil {

		return myerrors.DbErrors(err)
	}
	return nil
}

func UpdateMessages(feilds primitive.E, filter bson.M) error {

	update := bson.D{{Key: "$set", Value: bson.D{feilds}}}

	err := UpdateMany("messages", filter, update)
	if err != nil {

		return myerrors.DbErrors(err)
	}
	return nil
}

func NewMessage(user_id string, receiver_id string, message string) *Message {
	model := &Message{
		Text:       message,
		UserId:     user_id,
		ReceiverId: receiver_id,
		Status:     UNREAD,
	}
	return model
}
