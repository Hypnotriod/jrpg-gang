package session

import (
	"jrpg-gang/controller"
	"jrpg-gang/engine"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

type HubConfig struct {
	Addres                string `json:"addres"`
	ReadBufferSize        int    `json:"readBufferSize"`
	WriteBufferSize       int    `json:"writeBufferSize"`
	MaxMessageSize        int64  `json:"maxMessageSize"`
	UserOfflineTimeoutSec int64  `json:"userOfflineTimeoutSec"`
}

type Hub struct {
	sync.RWMutex
	config      HubConfig
	server      *http.Server
	upgrader    *websocket.Upgrader
	controller  *controller.GameController
	clients     map[engine.UserId]*Client
	leaveTimers map[engine.UserId]*time.Timer
}

func NewHub(config HubConfig, controller *controller.GameController) *Hub {
	hub := &Hub{}
	hub.config = config
	hub.controller = controller
	hub.clients = make(map[engine.UserId]*Client)
	hub.leaveTimers = make(map[engine.UserId]*time.Timer)
	hub.upgrader = &websocket.Upgrader{
		CheckOrigin:     hub.checkOrigin,
		ReadBufferSize:  config.ReadBufferSize,
		WriteBufferSize: config.WriteBufferSize,
	}
	hub.server = &http.Server{
		Addr: config.Addres,
	}
	controller.RegisterBroadcaster(hub)
	http.HandleFunc("/ws", hub.serveWsRequest)
	return hub
}

func (h *Hub) Start() error {
	return h.server.ListenAndServe()
}

func (h *Hub) BroadcastGameMessage(userIds []engine.UserId, message string) {
	for _, userId := range userIds {
		if client := h.getClient(userId); client != nil {
			client.WriteMessage(message)
		}
	}
}

func (h *Hub) checkOrigin(r *http.Request) bool {
	return true
}

func (h *Hub) serveWsRequest(writer http.ResponseWriter, request *http.Request) {
	conn, err := h.upgrader.Upgrade(writer, request, nil)
	if err != nil {
		log.Error("Can't serve ws request:", err)
		return
	}
	log.Info("Connection established")
	NewClient(conn, h).Serve()
}

func (h *Hub) registerClient(client *Client) {
	h.Lock()
	if oldClient, ok := h.clients[client.userId]; ok {
		oldClient.Kick()
	}
	h.clients[client.userId] = client
	if timer, ok := h.leaveTimers[client.userId]; ok && timer.Stop() {
		delete(h.leaveTimers, client.userId)
		h.Unlock()
		h.controller.ConnectionStatusChanged(client.userId, false)
		log.Info("Client back online: ", client.userId)
		return
	}
	h.Unlock()
	log.Info("Register Client: ", client.userId)
}

func (h *Hub) unregisterClient(userId engine.UserId) {
	if userId == engine.UserIdEmpty {
		log.Info("Client left without joining")
		return
	}
	h.Lock()
	delete(h.clients, userId)
	h.leaveTimers[userId] = time.AfterFunc(time.Duration(h.config.UserOfflineTimeoutSec)*time.Second, func() {
		h.Lock()
		delete(h.leaveTimers, userId)
		h.Unlock()
		h.controller.Leave(userId)
		log.Info("Unregister Client by timeout: ", userId)
	})
	h.Unlock()
	h.controller.ConnectionStatusChanged(userId, true)
	log.Info("Client went offline: ", userId)
}

func (h *Hub) getClient(userId engine.UserId) *Client {
	h.RLock()
	client, ok := h.clients[userId]
	h.RUnlock()
	if !ok {
		return nil
	}
	return client
}
