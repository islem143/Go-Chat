package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Message struct {
	ID         string             `json:"id" bson:"_id,omitempty"`
	ReceiverId string             `json:"receiver_id" bson:"receiver_id,omitempty"`
	UserId     string             `json:"user_id" bson:"user_id,omitempty"`
	Text       string             `json:"text" bson:"text,omitempty"`
	CreatedAt  primitive.DateTime `json:"created_at" bson:"created_at"`
	UpdatedAt  primitive.DateTime `json:"updated_at" bson:"updated_at"`
}

func (user *Message) Collection() string {
	return "messages"
}

//

func FindMessages(key string, value string) (*[]Message, error) {
	//	filter := bson.M{key: value}
	messages := &[]Message{}
	err := FindAll("messages", messages, nil)
	if err != nil {

		return nil, err
	}
	return messages, nil
}
func InsertMessages(message *Message) error {
	err := Insert(message.Collection(), message)
	if err != nil {

		return err
	}
	return nil
}

func NewMessage(user_id string, receiver_id string, message string) *Message {
	return &Message{
		Text:       message,
		UserId:     user_id,
		ReceiverId: receiver_id,
		CreatedAt:  primitive.NewDateTimeFromTime(time.Now()),
		UpdatedAt:  primitive.NewDateTimeFromTime(time.Now()),
	}
}
