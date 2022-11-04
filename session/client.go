package session

import (
	"jrpg-gang/engine"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

type Client struct {
	mu            sync.Mutex
	conn          *websocket.Conn
	hub           *Hub
	userId        engine.UserId
	noUserIdTimer *time.Timer
	kicked        bool
	left          bool
}

func NewClient(connection *websocket.Conn, hub *Hub) *Client {
	c := &Client{}
	c.conn = connection
	c.hub = hub
	c.userId = engine.UserIdEmpty
	c.kicked = false
	c.left = false
	c.conn.SetReadLimit(c.hub.config.MaxMessageSize)
	return c
}

func (c *Client) WriteMessage(message string) {
	c.mu.Lock()
	err := c.conn.WriteMessage(websocket.TextMessage, []byte(message))
	c.mu.Unlock()
	if err != nil && websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
		log.Error("Client (", c.userId, ") write message error:", err)
	}
}

func (c *Client) Serve() {
	c.begin()
	for {
		mt, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Error("Client (", c.userId, ") read message error:", err)
			}
			break
		}
		if mt != websocket.TextMessage {
			break
		}
		userId, response := c.hub.controller.HandleRequest(c.userId, message)
		if userId != engine.UserIdEmpty {
			c.manageUserId(userId)
		}
		c.WriteMessage(response)
	}
	c.complete()
}

func (c *Client) Kick() {
	defer c.mu.Unlock()
	c.mu.Lock()
	if !c.left && !c.kicked {
		c.conn.Close()
	}
	c.kicked = true
}

func (c *Client) begin() {
	c.noUserIdTimer = time.AfterFunc(time.Duration(c.hub.config.UserWithoutIdTimeoutSec)*time.Second, func() {
		c.Kick()
		log.Info("Client didn't obtain the id and was kicked")
	})
}

func (c *Client) manageUserId(userId engine.UserId) {
	if c.noUserIdTimer.Stop() {
		c.userId = userId
		c.hub.registerClient(c)
	} else {
		c.hub.controller.Leave(userId)
	}
}

func (c *Client) complete() {
	defer c.mu.Unlock()
	c.mu.Lock()
	c.left = true
	if !c.kicked {
		c.hub.unregisterClient(c.userId)
		c.conn.Close()
	}
}
