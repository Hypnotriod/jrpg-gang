package session

import (
	"jrpg-gang/engine"
	"time"

	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

type Client struct {
	conn     *websocket.Conn
	hub      *Hub
	userId   engine.UserId
	pongWait time.Duration
}

func NewClient(connection *websocket.Conn, hub *Hub) *Client {
	c := &Client{}
	c.conn = connection
	c.hub = hub
	c.pongWait = time.Second * time.Duration(hub.config.PongWaitSec)
	c.userId = engine.UserIdEmpty
	c.conn.SetReadLimit(c.hub.config.MaxMessageSize)
	c.conn.SetPongHandler(c.handlePong)
	return c
}

func (c *Client) Serve() {
	defer c.conn.Close()
	c.conn.SetReadDeadline(time.Now().Add(c.pongWait))
	for {
		mt, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("Client read message error: %v", err)
			}
			break
		}
		if mt != websocket.TextMessage {
			continue
		}
		data := string(message)
		userId, response := c.hub.controller.HandleRequest(c.userId, data)
		if userId != engine.UserIdEmpty {
			c.userId = userId
		}
		c.conn.WriteMessage(websocket.TextMessage, []byte(response))
	}
}

func (c *Client) handlePong(appData string) error {
	c.conn.SetReadDeadline(time.Now().Add(c.pongWait))
	return nil
}
