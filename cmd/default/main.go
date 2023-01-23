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
	key := flag.String("key", "", "path to TLS key.pem")
	cert := flag.String("cert", "", "path to TLS cert.pem")
	rBuffSize := flag.Int("rBuffSize", 1024, "ws read buffer size")
	wBuffSize := flag.Int("wBuffSize", 4096, "ws write buffer size")
	broadcasterPoolSize := flag.Int("broadcasterPoolSize", 32, "broadcaster routines pool size")
	broadcastQueueSize := flag.Int("broadcastQueueSize", 4096, "broadcast channel queue size")
	maxMessageSize := flag.Int64("maxMessageSize", 1024, "max message size sent by peer")
	userOfflineTimeoutSec := flag.Int64("userOfflineTimeoutSec", 10, "user offline timeout in secnds")
	userWithoutIdTimeoutSec := flag.Int64("userWithoutIdTimeoutSec", 1, "user hasn't obtained id timeout in secnds")
	flag.Parse()

	return session.HubConfig{
		Port:                    port,
		TlsKey:                  *key,
		TlsCert:                 *cert,
		ReadBufferSize:          *rBuffSize,
		WriteBufferSize:         *wBuffSize,
		BroadcasterPoolSize:     *broadcasterPoolSize,
		BroadcastQueueSize:      *broadcastQueueSize,
		MaxMessageSize:          *maxMessageSize,
		UserOfflineTimeoutSec:   *userOfflineTimeoutSec,
		UserWithoutIdTimeoutSec: *userWithoutIdTimeoutSec,
	}
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./public")))

	config := config()
	cntrl := controller.NewGameController()
	hub := session.NewHub(config, cntrl)
	hub.Start()
}
