package session

import (
	"jrpg-gang/engine"
	"sync"

	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

type Client struct {
	sync.RWMutex
	conn   *websocket.Conn
	hub    *Hub
	userId engine.UserId
}

func NewClient(connection *websocket.Conn, hub *Hub) *Client {
	c := &Client{}
	c.conn = connection
	c.hub = hub
	c.userId = engine.UserIdEmpty
	c.conn.SetReadLimit(c.hub.config.MaxMessageSize)
	return c
}

func (c *Client) WriteMessage(message string) {
	c.Lock()
	err := c.conn.WriteMessage(websocket.TextMessage, []byte(message))
	c.Unlock()
	if err != nil && websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
		log.Error("Client (", c.userId, ") write message error:", err)
	}
}

func (c *Client) Serve() {
	defer func() {
		c.hub.unregisterClient(c.userId)
		c.conn.Close()
	}()
	for {
		mt, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Error("Client (", c.userId, ") read message error:", err)
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
			c.hub.registerClient(c)
		}
		c.WriteMessage(response)
	}
}
