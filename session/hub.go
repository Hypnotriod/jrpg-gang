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
	Port                  string `json:"port"`
	ReadBufferSize        int    `json:"readBufferSize"`
	WriteBufferSize       int    `json:"writeBufferSize"`
	BroadcasterPoolSize   int    `json:"broadcasterPoolSize"`
	BroadcastQueueSize    int    `json:"broadcastQueueSize"`
	MaxMessageSize        int64  `json:"maxMessageSize"`
	UserOfflineTimeoutSec int64  `json:"userOfflineTimeoutSec"`
}

type Hub struct {
	mu            sync.RWMutex
	config        HubConfig
	server        *http.Server
	upgrader      *websocket.Upgrader
	controller    *controller.GameController
	clients       map[engine.UserId]*Client
	leaveTimers   map[engine.UserId]*time.Timer
	broadcastPool chan broadcast
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
		Addr: ":" + config.Port,
	}
	hub.broadcastPool = make(chan broadcast, config.BroadcastQueueSize)
	controller.RegisterBroadcaster(hub)
	http.HandleFunc("/ws", hub.serveWsRequest)
	return hub
}

func (h *Hub) Start() error {
	for n := h.config.BroadcasterPoolSize; n > 0; n-- {
		go h.broadcastGameMessageRoutine(h.broadcastPool)
	}
	return h.server.ListenAndServe()
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
	h.mu.Lock()
	if oldClient, ok := h.clients[client.userId]; ok {
		oldClient.Kick()
	}
	h.clients[client.userId] = client
	if timer, ok := h.leaveTimers[client.userId]; ok && timer.Stop() {
		delete(h.leaveTimers, client.userId)
		h.mu.Unlock()
		h.controller.ConnectionStatusChanged(client.userId, false)
		log.Info("Client back online: ", client.userId)
		return
	}
	h.mu.Unlock()
	log.Info("Register Client: ", client.userId)
}

func (h *Hub) unregisterClient(userId engine.UserId) {
	if userId == engine.UserIdEmpty {
		log.Info("Client left without joining")
		return
	}
	h.mu.Lock()
	delete(h.clients, userId)
	h.leaveTimers[userId] = time.AfterFunc(time.Duration(h.config.UserOfflineTimeoutSec)*time.Second, func() {
		h.mu.Lock()
		if _, ok := h.leaveTimers[userId]; !ok {
			h.mu.Unlock()
			return
		}
		delete(h.leaveTimers, userId)
		h.mu.Unlock()
		h.controller.Leave(userId)
		log.Info("Unregister Client by timeout: ", userId)
	})
	h.mu.Unlock()
	h.controller.ConnectionStatusChanged(userId, true)
	log.Info("Client went offline: ", userId)
}

func (h *Hub) getClient(userId engine.UserId) *Client {
	h.mu.RLock()
	client, ok := h.clients[userId]
	h.mu.RUnlock()
	if !ok {
		return nil
	}
	return client
}
