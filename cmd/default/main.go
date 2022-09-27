package main

import (
	"flag"
	"jrpg-gang/controller"
	"jrpg-gang/session"
	"net/http"
	"os"
)

func getenv(key string, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}

func config() session.HubConfig {
	port := getenv("PORT", "3000")
	rBuffSize := flag.Int("rBuffSize", 1024, "ws read buffer size")
	wBuffSize := flag.Int("wBuffSize", 1024, "ws write buffer size")
	broadcastPoolSizeSize := flag.Int("broadcastPoolSizeSize", 32, "broadcast pool size")
	maxMessageSize := flag.Int64("maxMessageSize", 4096, "connection max message size")
	userOfflineTimeoutSec := flag.Int64("userOfflineTimeoutSec", 10, "user offline timeout in secnds")
	flag.Parse()

	return session.HubConfig{
		Port:                  port,
		ReadBufferSize:        *rBuffSize,
		WriteBufferSize:       *wBuffSize,
		BroadcastPoolSize:     *broadcastPoolSizeSize,
		MaxMessageSize:        *maxMessageSize,
		UserOfflineTimeoutSec: *userOfflineTimeoutSec,
	}
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./public")))

	config := config()
	cntrl := controller.NewGameController()
	hub := session.NewHub(config, cntrl)
	hub.Start()
}
