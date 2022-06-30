package session

import (
	"jrpg-gang/controller"
	"jrpg-gang/engine"
	"jrpg-gang/util"
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
	leaveTimers map[engine.UserId]*util.Timer
}

func NewHub(config HubConfig, controller *controller.GameController) *Hub {
	hub := &Hub{}
	hub.config = config
	hub.controller = controller
	hub.clients = make(map[engine.UserId]*Client)
	hub.leaveTimers = make(map[engine.UserId]*util.Timer)
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
	if timer, ok := h.leaveTimers[client.userId]; ok {
		timer.Cancel()
		delete(h.leaveTimers, client.userId)
		log.Info("Client back online: ", client.userId)
	}
	h.Unlock()
	log.Info("Register Client: ", client.userId)
}

func (h *Hub) unregisterClient(userId engine.UserId) {
	if userId == engine.UserIdEmpty {
		log.Info("Client left without joining")
		return
	}
	log.Info("Client is offline: ", userId)
	h.Lock()
	delete(h.clients, userId)
	timer := util.NewTimer(time.Duration(h.config.UserOfflineTimeoutSec)*time.Second, func() {
		h.controller.Leave(userId)
		log.Info("Unregister Client by timeout: ", userId)
	})
	h.leaveTimers[userId] = timer
	h.Unlock()
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
