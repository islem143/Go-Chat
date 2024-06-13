package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt/v5"
	myerrors "github.com/islem143/go-chat/Errors"
	"github.com/islem143/go-chat/models"
)

const SecretKey = "secret"

func IsAuth(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	//auth := c.Get("Authorization")
	if cookie != "" {
		//res := strings.Split(auth, " ")[1]

		token, err := jwt.ParseWithClaims(cookie, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(SecretKey), nil //using the SecretKey which was generated in th Login function
		})

		if err != nil {
			log.Error(err)

			return myerrors.UnauthorizedError()
		}

		claims, ok := token.Claims.(*jwt.RegisteredClaims)

		if ok {
			user, err := models.FindUserById(claims.Issuer)
			if err != nil {
				log.Error(err)
			}

			if myerrors.DocumentNotFoundError(err) {
				return myerrors.NotFoundError("user not found")
			} else if err != nil {

				return myerrors.InternalServerError("internal server error")
			}

			c.Locals("user", user)
		}
	}

	return c.Next()
}
