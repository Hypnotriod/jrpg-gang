package session

import (
	"jrpg-gang/controller"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type HubConfig struct {
	Addres          string `json:"addres"`
	ReadBufferSize  int    `json:"readBufferSize"`
	WriteBufferSize int    `json:"writeBufferSize"`
	MaxMessageSize  int64  `json:"maxMessageSize"`
	PongWaitSec     int64  `json:"pongWaitSec"`
}

type Hub struct {
	config     HubConfig
	server     *http.Server
	upgrader   *websocket.Upgrader
	controller *controller.GameController
}

func NewHub(config HubConfig, controller *controller.GameController) *Hub {
	hub := &Hub{}
	hub.config = config
	hub.controller = controller
	hub.upgrader = &websocket.Upgrader{
		CheckOrigin:     hub.checkOrigin,
		ReadBufferSize:  config.ReadBufferSize,
		WriteBufferSize: config.WriteBufferSize,
	}
	hub.server = &http.Server{
		Addr: config.Addres,
	}
	http.HandleFunc("/ws", hub.serverWsRequest)
	return hub
}

func (h *Hub) Strart(config HubConfig) error {
	return h.server.ListenAndServe()
}

func (h *Hub) checkOrigin(r *http.Request) bool {
	return true
}

func (h *Hub) serverWsRequest(writer http.ResponseWriter, request *http.Request) {
	conn, err := h.upgrader.Upgrade(writer, request, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

}
