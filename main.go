package main

import (
	"jrpg-gang/controller"
	"jrpg-gang/session"
	"log"
)

func main() {
	config := session.HubConfig{
		Addres:          ":3000",
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		MaxMessageSize:  4096,
		PongWaitSec:     60,
	}

	cntrl := controller.NewGameController()
	hub := session.NewHub(config, cntrl)
	err := hub.Start()
	if err != nil {
		log.Printf("Can't start: %v\r\n", err)
	}
}
