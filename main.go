package main

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	myerrors "github.com/islem143/go-chat/Errors"
	"github.com/islem143/go-chat/database"
	"github.com/islem143/go-chat/routes"
)

func main() {

	database.Connect()

	app := fiber.New(fiber.Config{
		// Global custom error handler
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			message := "internal server error"
			var e *myerrors.ApiError
			if errors.As(err, &e) {
				code = e.Code
				return c.Status(code).JSON(e)
			}
			return c.Status(code).JSON(message)
		},
	})
	// app.Use(cors.New(cors.Config{
	// 	AllowCredentials: true, //Very important while using a HTTPonly Cookie, frontend can easily get and return back the cookie.
	// }))

	routes.Setup(app) // A routes package/folder is created with 'Setup' function created.

	app.Listen(":8000")
}
