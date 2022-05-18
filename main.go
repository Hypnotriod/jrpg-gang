package main

import (
	"jrpg-gang/controller"
	"jrpg-gang/session"
	"net/http"
)

func main() {
	config := session.HubConfig{
		Addres:          ":3000",
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		MaxMessageSize:  4096,
	}

	http.Handle("/", http.FileServer(http.Dir("./public")))

	cntrl := controller.NewGameController()
	hub := session.NewHub(config, cntrl)
	hub.Start()
}
