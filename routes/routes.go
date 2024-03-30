package routes

import (
	"github.com/gofiber/fiber/v2"
	apis "github.com/islem143/go-chat/api"
	"github.com/islem143/go-chat/middlewares"
)

func Setup(app *fiber.App) {
	api := app.Group("/")
	api.Post("users/login", apis.Login)
	api.Post("users/register", apis.Register)
	api.Use(middlewares.IsAuth)
	api.Get("users/", apis.List)
	api.Get("users/:id", apis.One)
	api.Post("users/logout", apis.Logout)
	//api.Use(middlewares.HasRole(types.ADMIN))

	messages := app.Group("/messages")
	messages.Get("/", apis.GetMessages)
	messages.Post("/", apis.InsertMessage)
}
