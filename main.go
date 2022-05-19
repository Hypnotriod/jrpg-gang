package main

import (
	"flag"
	"jrpg-gang/controller"
	"jrpg-gang/session"
	"net/http"
)

func config() session.HubConfig {
	addres := flag.String("address", ":3000", "host address")
	rBuffSize := flag.Int("rBuffSize", 1024, "ws read buffer size")
	wBuffSize := flag.Int("wBuffSize", 1024, "ws write buffer size")
	maxMessageSize := flag.Int64("maxMessageSize", 4096, "connection max message size")
	flag.Parse()

	return session.HubConfig{
		Addres:          *addres,
		ReadBufferSize:  *rBuffSize,
		WriteBufferSize: *wBuffSize,
		MaxMessageSize:  *maxMessageSize,
	}
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./public")))

	config := config()
	cntrl := controller.NewGameController()
	hub := session.NewHub(config, cntrl)
	hub.Start()
}
