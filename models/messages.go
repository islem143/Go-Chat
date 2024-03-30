package models

type Message struct {
	BaseModel
	ReceiverId string `json:"receiver_id" bson:"receiver_id,omitempty"`
	UserId     string `json:"user_id" bson:"user_id,omitempty"`
	Text       string `json:"text" bson:"text,omitempty"`
}

func (user *Message) CollectionName() string {
	return "messages"
}

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
	err := Insert(message.CollectionName(), message)
	if err != nil {

		return err
	}
	return nil
}

func NewMessage(user_id string, receiver_id string, message string) *Message {
	model := &Message{
		Text:       message,
		UserId:     user_id,
		ReceiverId: receiver_id,
	}
	return model
}
