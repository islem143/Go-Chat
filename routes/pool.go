package routes

import (
	"log"

	"github.com/gofiber/contrib/websocket"
	"github.com/islem143/go-chat/models"
)

type Client struct {
	id   string
	conn *websocket.Conn
	//pool Pool
}

type Pool struct {
	Messages chan *models.Message
	Clients  map[string]*Client
}

func NewPool() *Pool {
	return &Pool{
		Messages: make(chan *models.Message),
		Clients:  make(map[string]*Client),
	}
}

func (p *Pool) Start() {

	for {

		msg := <-p.Messages
		recevierId := msg.ReceiverId
		log.Println(msg.Type)
		val, ok := p.Clients[recevierId]
		if ok {

			//f := primitive.E{Key: "read", Value: true}
			//models.UpdateMessages(msg, f)
			val.conn.WriteJSON((msg))

		}

	}

}
