package api

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/golang-jwt/jwt/v5"
	myerrors "github.com/islem143/go-chat/Errors"
	"github.com/islem143/go-chat/models"
	"github.com/islem143/go-chat/validation"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {

	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}
	err := validation.Validate(validation.RegisterRequestValidator{
		Name:  data["name"],
		Email: data["email"],
	})
	if err != nil {
		return err
	}

	user, err := models.FindUser("email", data["email"])

	if err != nil && user != nil {

		return err
	}
	if user != nil {
		return myerrors.ClientError("email already exsists")
	}

	password, err := bcrypt.GenerateFromPassword([]byte(data["password"]), 6)
	if err != nil {
		log.Error(err)
		return myerrors.InternalServerError("interal server error")

	}
	user = &models.User{
		Name:     data["name"],
		Email:    data["email"],
		Password: password,
	}
	err = models.Insert(user.CollectionName(), user)
	if err != nil {

		return myerrors.InternalServerError("interal server error")
	}

	return c.JSON(user)

}

func List(c *fiber.Ctx) error {
	//authUser := c.Locals("user").(*models.User)

	users, err := models.FindAllUsers(bson.D{})
	if err != nil {
		return err
	}
	return c.JSON(ApiResponseList{List: users})
}

func One(c *fiber.Ctx) error {
	id := c.Params("id")

	if id == "" {
		return myerrors.NotFoundError("document not Found")
	}
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}
	user, err := models.FindUserById(id)
	if err != nil {
		return err
	}
	return c.JSON(user)
}

const SecretKey = "secret"

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}
	err := validation.Validate(validation.LoginRequestValidator{
		Email:    data["email"],
		Password: string(data["password"]),
	})
	if err != nil {
		return err
	}

	user, err := models.FindUser("email", data["email"])
	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"]))
	if err != nil {
		return myerrors.UnauthorizedError()

	}
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    user.ID,                                            //issuer contains the ID of the user.
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)), //Adds time to the token i.e. 24 hours.
	})
	token, errf := claims.SignedString([]byte(SecretKey))

	if errf != nil {
		return myerrors.ClientError("could not login")

	}
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	} //Creates the cookie to be passed.

	c.Cookie(&cookie)

	// return c.JSON(fiber.Map{
	// 	"message": "success",
	// 	"user":    user,
	// 	"token":   token,
	// })
	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour), //Sets the expiry time an hour ago in the past.
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(ApiResponse{Message: "logged out successfuly."})

}
