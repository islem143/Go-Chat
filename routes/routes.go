package routes

import (
	"fmt"
	"log"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	apis "github.com/islem143/go-chat/api"
	"github.com/islem143/go-chat/models"
)

type Message struct {
	Message    string `json:"message"`
	ReceiverId string `json:"receiver"`
}

func Setup(app *fiber.App) {
	pool := NewPool()

	go pool.Start()
	api := app.Group("/")
	api.Post("users/login", apis.Login)
	api.Post("users/register", apis.Register)
	//api.Use(middlewares.IsAuth)
	api.Get("users/", apis.List)
	api.Get("contacts/", apis.ContactList)
	api.Post("contacts/", apis.AddContact)
	api.Get("users/:id", apis.One)
	api.Post("users/logout", apis.Logout)
	//api.Use(middlewares.HasRole(types.ADMIN))

	messages := app.Group("/messages")
	messages.Get("/", apis.GetMessages)
	messages.Post("/", apis.InsertMessage)

	messages.Use("/ws", func(c *fiber.Ctx) error {
		// IsWebSocketUpgrade returns true if the client
		// requested upgrade to the WebSocket protocol.
		fmt.Println("ttt")
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})
	app.Get("/ws/private-chat/:id", websocket.New(func(c *websocket.Conn) {
		id := c.Params("id")
		if id == "" {
			return
		}
		authUser := c.Locals("user").(*models.User)
		if authUser.ID != id {
			return
		}

		// _, ok := pool.Clients[authUser.ID]
		// if !ok {
		// 	pool.Clients[authUser.ID] = &Client{conn: c}
		// }
		pool.Clients["123"] = &Client{id: "123", conn: c}

		var (
			err error
		)
		message := new(Message)
		for {
			if err = c.ReadJSON(message); err != nil {
				log.Println("read:", err)
				break
			}

			fmt.Println("Received message:", message.Message)
			fmt.Println("Receiver", message.ReceiverId)
			msg := models.Message{
				ReceiverId: message.ReceiverId,
				Text:       message.Message,
				UserId:     authUser.ID,
				Read:       false,
			}
			err := models.InsertMessages(&msg)
			if err != nil {
				return
			}

			_, ok := pool.Clients[message.ReceiverId]

			if !ok {
				// store the message in notification.
				continue
			} else {

				pool.Messages <- &msg
			}
		}

	}))

}
