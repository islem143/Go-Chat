package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/islem143/go-chat/models"
	"github.com/islem143/go-chat/validation"
	"go.mongodb.org/mongo-driver/bson"
)

type GetMessageRequsetBody struct {
	ReceiverId string `json:"receiver_id"`
	UserId     string `json:"user_id"` // This could be any type
}

func GetMessages(c *fiber.Ctx) error {
	var requestBody GetMessageRequsetBody
	if err := c.BodyParser(&requestBody); err != nil {

		return err
	}
	filter := bson.D{{Key: "user_id", Value: requestBody.UserId}, {Key: "receiver_id", Value: requestBody.ReceiverId}}
	messages, err := models.FindMessages(filter)

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

	err = models.InsertMessages(message)
	if err != nil {
		return err
	}
	return c.JSON(message)
}
