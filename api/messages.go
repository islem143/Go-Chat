package api

import (
	"github.com/gofiber/fiber/v2"
	myerrors "github.com/islem143/go-chat/Errors"
	"github.com/islem143/go-chat/models"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetMessages(c *fiber.Ctx) error {
	messages, err := models.FindMessages("user_id", "4")
	if err == mongo.ErrNoDocuments {
		return myerrors.NotFoundError("document not Found")

	}
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
	user := c.Locals("user").(*models.User)
	message := models.NewMessage(user.ID, data["receiver_id"], data["message"])

	err := models.InsertMessages(message)
	if err != nil {

		return myerrors.InternalServerError("interal server error")
	}
	return c.JSON(message)
}
