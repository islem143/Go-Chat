package api

import (
	"github.com/gofiber/fiber/v2"
	myerrors "github.com/islem143/go-chat/Errors"
	"github.com/islem143/go-chat/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RequestBody struct {
	Id      string `json:"id"`
	Contact string `json:"contact"` // This could be any type
}

func AddContact(c *fiber.Ctx) error {
	var requestBody RequestBody
	if err := c.BodyParser(&requestBody); err != nil {

		return err
	}

	user, err := models.FindUserById(requestBody.Id)

	if err != nil {

		return err
	}

	objectId, err := primitive.ObjectIDFromHex(user.ID)
	if err != nil {
		return myerrors.DbErrors(err)
	}
	filter := bson.M{"_id": objectId}
	update := bson.D{
		{Key: "$push", Value: bson.D{
			{Key: "contacts", Value: requestBody.Contact},
		}},
	}
	err = models.UpdateUser(user, filter, update)
	if err != nil {

		return err
	}
	return c.JSON(ApiResponse{Message: "Contact Added"})
}

func ContactList(c *fiber.Ctx) error {
	//authUser := c.Locals("user").(*models.User)
	res, err := primitive.ObjectIDFromHex("6647d0bfa347f75b4f7d3ee2")
	if err != nil {
		return myerrors.DbErrors(err)
	}
	filter := bson.D{{Key: "_id", Value: bson.D{{Key: "$ne", Value: res}}}}

	users, err := models.FindAllUsers(filter)
	if err != nil {
		return err
	}
	return c.JSON(ApiResponseList{List: users})
}
