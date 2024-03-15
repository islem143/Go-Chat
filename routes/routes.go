package routes

import (
	"context"
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/islem143/go-chat/database"
	"github.com/islem143/go-chat/models"
	"go.mongodb.org/mongo-driver/mongo"
)

func IsDup(err error) bool {
	var e mongo.WriteException
	if errors.As(err, &e) {
		for _, we := range e.WriteErrors {
			if we.Code == 11000 {
				return true
			}
		}
	}
	return false
}
func Setup(app *fiber.App) {
	api := app.Group("/")

	api.Get("/", func(c *fiber.Ctx) error {
		return fiber.NewError(fiber.StatusServiceUnavailable, "On vacation!")

		user := models.User{Name: "islem", Email: "test@test.com"}
		result, err := database.Client.Database(database.Database).Collection("users").InsertOne(context.TODO(), user)
		if err != nil {
			if IsDup(err) {
				return c.JSON("duplciate email")
			}

		}
		return c.JSON(result)
	})
	// api.Get("/get-user/", controllers.User)
	// api.Post("/register", controllers.Register)

	// api.Post("/login", controllers.Login)

	// api.Post("/logout", controllers.Logout)
}
