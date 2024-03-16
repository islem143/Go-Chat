package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/islem143/go-chat/database"
	"github.com/islem143/go-chat/models"
	"github.com/islem143/go-chat/routes"
)

func main() {

	database.Connect()
	models.SetUp()
	app := fiber.New(fiber.Config{
		// Global custom error handler
		ErrorHandler: func(c *fiber.Ctx, err error) error {

			return c.Status(fiber.StatusBadRequest).JSON(err)
		},
	})
	// app.Use(cors.New(cors.Config{
	// 	AllowCredentials: true, //Very important while using a HTTPonly Cookie, frontend can easily get and return back the cookie.
	// }))

	routes.Setup(app) // A routes package/folder is created with 'Setup' function created.

	app.Listen(":8000")
}
