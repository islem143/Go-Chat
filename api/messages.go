package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/islem143/go-chat/models"
	"github.com/islem143/go-chat/validation"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type GetMessageRequsetBody struct {
	ReceiverId string `json:"receiver_id"`
	UserId     string `json:"user_id"` // This could be any type
}

type UpdateMessageRequestBody struct {
	MessageId     string `json:"messageId"`
	ReceiverId    string `json:"receiverId"`
	ReadAllLatest bool   `json:"readAllLatest"`
}

func GetMessages(c *fiber.Ctx) error {

	m := c.Queries()

	opts := options.Find().SetSort(bson.D{{Key: "created_Value: at", Value: -1}})
	authUser := c.Locals("user").(*models.User)

	//      bson.D{{Key: "user_id", Value: authUser.ID}},
	// 		bson.D{{Key: "user_id", Value: m["receiver_id"]}}
	filter := bson.D{{
		Key: "$or", Value: bson.A{
			bson.D{{Key: "$and", Value: bson.A{bson.D{{Key: "user_id", Value: authUser.ID}}, bson.D{{Key: "receiver_id", Value: m["receiver_id"]}}}}},
			bson.D{{Key: "$and", Value: bson.A{bson.D{{Key: "user_id", Value: m["receiver_id"]}}, bson.D{{Key: "receiver_id", Value: authUser.ID}}}}},
		},
	},
	}
	messages, err := models.FindMessages(filter, opts)

	if err != nil {
		return err

	}

	return c.JSON(ApiResponseList{List: messages})
}

func InsertMessage(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	valError := validation.Validate(validation.MessageRequestValidator{
		Message:    data["message"],
		ReceiverId: data["receiver_id"],
	})

	if valError != nil {

		return valError
	}
	_, err := models.FindUserById(data["receiver_id"])

	if err != nil {

		return err
	}

	authUser := c.Locals("user").(*models.User)
	message := models.NewMessage(authUser.ID, data["receiver_id"], data["message"])
	fmt.Println(message)
	err = models.InsertMessages(message)
	if err != nil {
		return err
	}
	return c.JSON(message)
}
func MarkAsRead(c *fiber.Ctx) error {
	var data UpdateMessageRequestBody

	if err := c.BodyParser(&data); err != nil {
		return err

	}
	authUser := c.Locals("user").(*models.User)

	f := primitive.E{Key: "status", Value: models.READ}
	models.UpdateMessages(f, bson.M{"user_id": data.ReceiverId, "receiver_id": authUser.ID})
	return c.JSON("ok")

	// message, err := models.FindMessage(data.MessageId)

	// if message.UserId != authUser.ID {
	// 	return myerrors.UnauthorizedError()
	// }

	// if err != nil {
	// 	return err

	// }
	// message.Status = models.READ
	// f := primitive.E{Key: "status", Value: models.READ}
	// models.UpdateMessage(message, f)
	//return c.JSON("ok")
}
