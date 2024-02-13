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
	pingTimer       *time.Timer
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
	c.conn.SetWriteDeadline(time.Now().Add(time.Duration(c.hub.config.WriteDeadlineSec) * time.Second))
	err := c.conn.WriteMessage(websocket.TextMessage, message)
	c.mu.Unlock()
	if err != nil {
		if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
			log.Error("Client (", c.Info(), ") write message error:", err)
		}
		c.Kick()
	}
}

func (c *Client) Ping() {
	c.mu.Lock()
	if c.left || c.kicked {
		c.mu.Unlock()
		return
	}
	// log.Info("Client (", c.Info(), ") ping")
	c.conn.SetWriteDeadline(time.Now().Add(time.Duration(c.hub.config.WriteDeadlineSec) * time.Second))
	err := c.conn.WriteMessage(websocket.PingMessage, []byte{})
	c.mu.Unlock()
	if err != nil {
		if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
			log.Error("Client (", c.Info(), ") ping error:", err)
		}
		c.Kick()
	}
}

func (c *Client) Serve() {
	c.begin()
	for {
		c.updateReadDeadline()
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
		if c.playerId == engine.PlayerIdEmpty {
			if playerId == engine.PlayerIdEmpty {
				c.WriteMessage(response)
				break
			}
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
		log.Info("Client didn't obtain the id in time and was kicked: ", c.Info())
	})
	c.conn.SetPongHandler(func(string) error {
		return c.updateReadDeadline()
	})
}

func (c *Client) updateReadDeadline() error {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.pingTimer != nil {
		c.pingTimer.Reset(time.Duration(c.hub.config.PingTimeoutSec) * time.Second)
	} else {
		c.pingTimer = time.AfterFunc(time.Duration(c.hub.config.PingTimeoutSec)*time.Second, c.Ping)
	}
	return c.conn.SetReadDeadline(time.Now().Add(time.Duration(c.hub.config.ReadDeadlineSec) * time.Second))
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
	c.noPlayerIdTimer.Stop()
	c.pingTimer.Stop()
	c.left = true
	if !c.kicked {
		c.hub.unregisterClient(c)
		c.conn.Close()
	}
}
