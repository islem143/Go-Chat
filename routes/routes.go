package routes

import (
	"fmt"
	"log"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	apis "github.com/islem143/go-chat/api"
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
	app.Get("/ws/:id", websocket.New(func(c *websocket.Conn) {
		//id := c.Params("id")
		//authUser := c.Locals("user").(*models.User)
		pool.Clients["123"] = &Client{id: "123", conn: c}
		// if authUser.ID != id {
		// 	c.WriteJSON(&Message{Message: "unauthorized"})
		// 	return
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

			fmt.Println("Received message:", message.Message)
			fmt.Println("Receiver", message.ReceiverId)
			msg := &Message{

				Message:    message.Message,
				ReceiverId: message.ReceiverId,
			}
			_, ok := pool.Clients[msg.ReceiverId]
			if !ok {
				
			}
			pool.Messages <- msg
		}

	}))

}
