package middlewares

import (
	"github.com/gofiber/fiber/v2"
)

func HasRole(role string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		//user := c.Locals("user").(*models.User)
		//fmt.Println(user.Role)

		return nil
	}
}
