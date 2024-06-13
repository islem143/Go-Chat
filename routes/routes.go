package routes

import (
	"fmt"
	"log"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	apis "github.com/islem143/go-chat/api"
	"github.com/islem143/go-chat/middlewares"
	"github.com/islem143/go-chat/models"
)

type Message struct {
	Message    string `json:"message"`
	ReceiverId string `json:"receiver"`
	UserId     string `json:"user_id"`
}

type MessageLogin struct {
	Login string `json:"login"`
}

func Setup(app *fiber.App) {
	pool := NewPool()

	go pool.Start()
	api := app.Group("/")
	api.Post("users/login", apis.Login)
	api.Post("users/register", apis.Register)
	api.Use(middlewares.IsAuth)
	api.Get("users/", apis.List)
	api.Get("contacts/", apis.ContactList)
	api.Post("contacts/", apis.AddContact)
	api.Get("users/:id", apis.One)
	api.Post("users/logout", apis.Logout)
	//api.Use(middlewares.HasRole(types.ADMIN))
	app.Use("/ws", func(c *fiber.Ctx) error {
		// IsWebocketUpgrade returns true if the client
		// requested upgrade to the WebSocket protocol.

		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})
	// app.Get("/ws/connect", websocket.New(func(c *websocket.Conn) {

	// 	msg := new(MessageLogin)
	// 	for {
	// 		if err := c.ReadJSON(msg); err != nil {
	// 			log.Println("read:", err)
	// 			c.Close()
	// 		}
	// 		fmt.Printf(msg.Login)

	// 	}

	// }))
	app.Get("/ws/private-chat/:id", websocket.New(func(c *websocket.Conn) {
		id := c.Params("id")
		if id == "" {
			return
		}
		authUser := c.Locals("user").(*models.User)
		fmt.Println(authUser, "auth")
		// if authUser.ID != id {
		// 	return
		// }
		_, ok := pool.Clients[authUser.ID]
		if !ok {

			pool.Clients[authUser.ID] = &Client{id: authUser.ID, conn: c}
		}

		defer func() {
			delete(pool.Clients, authUser.ID)

			c.Close()
		}()

		// _, ok := pool.Clients[authUser.ID]
		// if !ok {
		// 	pool.Clients[authUser.ID] = &Client{conn: c}
		// }

		var (
			err error
		)
		message := new(Message)
		for {
			if err = c.ReadJSON(message); err != nil {
				log.Println("read:", err)
				break
			}
			if message.Message == "" {

				continue
			}
			msg := models.Message{
				ReceiverId: message.ReceiverId,
				Text:       message.Message,
				UserId:     message.UserId,
				Read:       false,
			}
			err := models.InsertMessages(&msg)
			if err != nil {
				return
			}

			_, ok = pool.Clients[message.ReceiverId]

			if !ok {
				// store the message in notification.
				continue
			} else {

				pool.Messages <- &msg
			}
		}

	}))
	messages := app.Group("/messages")
	messages.Get("/", apis.GetMessages)
	messages.Post("/", apis.InsertMessage)

}
