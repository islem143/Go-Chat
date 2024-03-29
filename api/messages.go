package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	myerrors "github.com/islem143/go-chat/Errors"
	"github.com/islem143/go-chat/models"
	"github.com/islem143/go-chat/validation"
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
	valError := validation.ValidateRequset(validation.MessageRequestValidator{
		Message:    data["message"],
		ReceiverId: data["receiver_id"],
	})

	if valError != nil {
		log.Error(valError)
		return myerrors.ClientBodyError(valError.Message)
	}
	_, err := models.FindUserById(data["receiver_id"])

	if myerrors.DocumentNotFoundError(err) {
		return myerrors.NotFoundError("user not found")
	} else if err != nil {
		log.Error(err)
		return myerrors.InternalServerError("internal server error")
	}

	authUser := c.Locals("user").(*models.User)
	message := models.NewMessage(authUser.ID, data["receiver_id"], data["message"])

	err = models.InsertMessages(message)
	if err != nil {

		return myerrors.InternalServerError("interal server error")
	}
	return c.JSON(message)
}
