package api

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"

	"github.com/golang-jwt/jwt/v5"
	myerrors "github.com/islem143/go-chat/Errors"
	"github.com/islem143/go-chat/models"
	"github.com/islem143/go-chat/validation"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {

	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	valError := validation.ValidateRequset(validation.RegisterRequestValidator{
		Name:  data["name"],
		Email: data["email"],
	})
	if valError != nil {
		log.Error(valError)
		return valError
	}
	dbuser, err := models.FindUser("email", data["email"])
	if err != nil && err != mongo.ErrNoDocuments {
		log.Error(err)
		return myerrors.InternalServerError("interal server error")
	}
	if dbuser != nil {

		return myerrors.RecordExistsError("Email")
	}

	password, err := bcrypt.GenerateFromPassword([]byte(data["password"]), 6)
	if err != nil {
		log.Error(err)
		return myerrors.InternalServerError("interal server error")

	}
	user := &models.User{
		Name:     data["name"],
		Email:    data["email"],
		Password: password,
	}
	err = models.InsertUser(user)
	if err != nil {

		return myerrors.InternalServerError("interal server error")
	}

	return c.JSON(user)

}
func List(c *fiber.Ctx) error {
	users, err := models.FindAllUsers()
	if err == mongo.ErrNoDocuments {
		return myerrors.NotFoundError("document not Found")

	}
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
	if err == mongo.ErrNoDocuments {
		return myerrors.NotFoundError("document not Found")

	}
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

	valError := validation.ValidateRequset(validation.LoginRequestValidator{
		Email:    data["email"],
		Password: string(data["password"]),
	})

	if valError != nil {
		log.Error(valError)
		return valError
	}

	user, err := models.FindUser("email", data["email"])
	if err != nil && err != mongo.ErrNoDocuments {
		log.Error(err)
		return myerrors.InternalServerError("interal server error")
	}
	if user == nil {

		return myerrors.NotFoundError("user not found")
	}

	err = bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"]))
	if err != nil {
		log.Error(err)
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "incorrect password",
		})
	}
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    user.ID,                                            //issuer contains the ID of the user.
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)), //Adds time to the token i.e. 24 hours.
	})
	token, errf := claims.SignedString([]byte(SecretKey))

	if errf != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "could not login",
		})
	}
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	} //Creates the cookie to be passed.

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})

}

// func User(c *fiber.Ctx) error {

// 	cookie := c.Cookies("jwt")

// 	token, err := jwt.ParseWithClaims(cookie, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
// 		return []byte(SecretKey), nil //using the SecretKey which was generated in th Login function
// 	})

// 	if err != nil {
// 		c.Status(fiber.StatusUnauthorized)
// 		return c.JSON(fiber.Map{
// 			"message": "unauthenticated",
// 		})
// 	}

// 	claims := token.Claims.(*jwt.RegisteredClaims)

// 	var user models.User

// 	database.DB.Where("id = ?", claims.Issuer).First(&user)

// 	return c.JSON(user)

// }

// func Logout(c *fiber.Ctx) error {
// 	cookie := fiber.Cookie{
// 		Name:     "jwt",
// 		Value:    "",
// 		Expires:  time.Now().Add(-time.Hour), //Sets the expiry time an hour ago in the past.
// 		HTTPOnly: true,
// 	}

// 	c.Cookie(&cookie)

// 	return c.JSON(fiber.Map{
// 		"message": "success",
// 	})

// }
