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
	Port                    string `json:"port"`
	TlsKey                  string `json:"tlsKey"`
	TlsCert                 string `json:"tlsCert"`
	ReadBufferSize          int    `json:"readBufferSize"`
	WriteBufferSize         int    `json:"writeBufferSize"`
	BroadcasterPoolSize     int    `json:"broadcasterPoolSize"`
	BroadcastQueueSize      int    `json:"broadcastQueueSize"`
	MaxMessageSize          int64  `json:"maxMessageSize"`
	UserOfflineTimeoutSec   int64  `json:"userOfflineTimeoutSec"`
	UserWithoutIdTimeoutSec int64  `json:"userWithoutIdTimeoutSec"`
}

type Hub struct {
	mu            sync.RWMutex
	config        HubConfig
	server        *http.Server
	upgrader      *websocket.Upgrader
	controller    *controller.GameController
	clients       map[engine.PlayerId]*Client
	offlineTimers map[engine.PlayerId]*time.Timer
	broadcastPool chan broadcast
}

func NewHub(config HubConfig, controller *controller.GameController) *Hub {
	hub := &Hub{}
	hub.config = config
	hub.controller = controller
	hub.clients = make(map[engine.PlayerId]*Client)
	hub.offlineTimers = make(map[engine.PlayerId]*time.Timer)
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
	if len(h.config.TlsKey) > 0 && len(h.config.TlsCert) > 0 {
		return h.server.ListenAndServeTLS(h.config.TlsCert, h.config.TlsKey)
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
	if oldClient, ok := h.clients[client.playerId]; ok {
		oldClient.Kick()
	}
	h.clients[client.playerId] = client
	if timer, ok := h.offlineTimers[client.playerId]; ok {
		delete(h.offlineTimers, client.playerId)
		timer.Stop()
		h.mu.Unlock()
		h.controller.ConnectionStatusChanged(client.playerId, false)
		log.Info("Client back online: ", client.playerId)
		return
	}
	h.mu.Unlock()
	log.Info("Register Client: ", client.playerId)
}

func (h *Hub) unregisterClient(playerId engine.PlayerId) {
	if playerId == engine.PlayerIdEmpty {
		log.Info("Client left without joining")
		return
	}
	h.mu.Lock()
	delete(h.clients, playerId)
	h.setupUserOfflineTimeout(playerId)
	h.mu.Unlock()
	h.controller.ConnectionStatusChanged(playerId, true)
	log.Info("Client went offline: ", playerId)
}

func (h *Hub) setupUserOfflineTimeout(playerId engine.PlayerId) {
	h.offlineTimers[playerId] = time.AfterFunc(time.Duration(h.config.UserOfflineTimeoutSec)*time.Second, func() {
		h.mu.Lock()
		if _, ok := h.offlineTimers[playerId]; !ok {
			h.mu.Unlock()
			return
		}
		delete(h.offlineTimers, playerId)
		h.mu.Unlock()
		h.controller.Leave(playerId)
		log.Info("Unregister Client by timeout: ", playerId)
	})
}

func (h *Hub) getClient(playerId engine.PlayerId) *Client {
	h.mu.RLock()
	client, ok := h.clients[playerId]
	h.mu.RUnlock()
	if !ok {
		return nil
	}
	return client
}
