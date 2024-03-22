package routes

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	apis "github.com/islem143/go-chat/api"
	"go.mongodb.org/mongo-driver/mongo"
)

func extractKeyFromErrorMessage(message string) string {
	parts := strings.Split(message, "{ : \"")
	if len(parts) < 2 {
		return "" // unable to extract
	}
	key := strings.Split(parts[1], "\"")[0]
	return key
}
func IsDup(err error) bool {
	var e mongo.WriteException
	if errors.As(err, &e) {
		for _, we := range e.WriteErrors {
			if we.Code == 11000 {
				fmt.Println(extractKeyFromErrorMessage(we.Details.String()))
				return true
			}
		}
	}
	return false
}

func Setup(app *fiber.App) {
	api := app.Group("/")
	api.Get("users/", apis.List)
	api.Get("users/:id", apis.One)
	api.Post("users/login", apis.Login)
	api.Post("users/register", apis.Register)

}
