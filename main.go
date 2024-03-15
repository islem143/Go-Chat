package main

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/islem143/go-chat/database"
	"github.com/islem143/go-chat/models"
	"github.com/islem143/go-chat/routes"
	"github.com/islem143/go-chat/validation"
)

func main() {

	database.Connect()
	models.SetUp()
	app := fiber.New(fiber.Config{
		// Global custom error handler
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}
			fmt.Print(code)
			return c.Status(fiber.StatusBadRequest).JSON(validation.GlobalErrorHandlerResp{
				Success: false,
				Message: err.Error(),
			})
		},
	})
	// app.Use(cors.New(cors.Config{
	// 	AllowCredentials: true, //Very important while using a HTTPonly Cookie, frontend can easily get and return back the cookie.
	// }))

	routes.Setup(app) // A routes package/folder is created with 'Setup' function created.

	app.Listen(":8000")
}
