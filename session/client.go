package session

import (
	"jrpg-gang/engine"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

type Client struct {
	mu              sync.Mutex
	conn            *websocket.Conn
	hub             *Hub
	playerId        engine.PlayerId
	noPlayerIdTimer *time.Timer
	kicked          bool
	left            bool
}

func NewClient(connection *websocket.Conn, hub *Hub) *Client {
	c := &Client{}
	c.conn = connection
	c.hub = hub
	c.playerId = engine.PlayerIdEmpty
	c.kicked = false
	c.left = false
	c.conn.SetReadLimit(c.hub.config.MaxMessageSize)
	return c
}

func (c *Client) WriteMessage(message []byte) {
	c.mu.Lock()
	if c.left || c.kicked {
		c.mu.Unlock()
		return
	}
	err := c.conn.WriteMessage(websocket.TextMessage, message)
	c.mu.Unlock()
	if err != nil && websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
		log.Error("Client (", c.Info(), ") write message error:", err)
	}
}

func (c *Client) Serve() {
	c.begin()
	for {
		mt, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Error("Client (", c.Info(), ") read message error:", err)
			}
			break
		}
		if mt != websocket.TextMessage {
			break
		}
		playerId, response := c.hub.controller.HandleRequest(c.playerId, message)
		if playerId != engine.PlayerIdEmpty {
			c.managePlayerId(playerId)
		}
		c.WriteMessage(response)
	}
	c.complete()
}

func (c *Client) Kick() {
	c.mu.Lock()
	defer c.mu.Unlock()
	if !c.left && !c.kicked {
		c.conn.Close()
	}
	c.kicked = true
}

func (c *Client) Info() string {
	return c.conn.RemoteAddr().String() + " " + string(c.playerId)
}

func (c *Client) begin() {
	c.noPlayerIdTimer = time.AfterFunc(time.Duration(c.hub.config.UserWithoutIdTimeoutSec)*time.Second, func() {
		c.Kick()
		log.Info("Client didn't obtain the id and was kicked: ", c.Info())
	})
}

func (c *Client) managePlayerId(playerId engine.PlayerId) {
	if c.noPlayerIdTimer.Stop() {
		c.playerId = playerId
		c.hub.registerClient(c)
	} else {
		c.hub.controller.Leave(playerId)
	}
}

func (c *Client) complete() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.left = true
	if !c.kicked {
		c.hub.unregisterClient(c)
		c.conn.Close()
	}
}
