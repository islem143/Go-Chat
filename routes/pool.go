package routes

import (
	"github.com/gofiber/contrib/websocket"
)

type Client struct {
	id   string
	conn *websocket.Conn
	//pool Pool
}

type Pool struct {
	Messages chan *Message
	Clients  map[string]*Client
}

func NewPool() *Pool {
	return &Pool{
		Messages: make(chan *Message),
		Clients:  make(map[string]*Client),
	}
}

func (p *Pool) Start() {

	for {

		msg := <-p.Messages
		recevierId := "123"

		val, ok := p.Clients[recevierId]
		if ok {
			val.conn.WriteJSON((msg.Message))

		}

	}

}
