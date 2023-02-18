package main

import (
	"flag"
	"jrpg-gang/auth"
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

func hubConfig() session.HubConfig {
	port := getenv("PORT", "3000")
	key := flag.String("key", "", "path to TLS key.pem")
	cert := flag.String("cert", "", "path to TLS cert.pem")
	rBuffSize := flag.Int("rBuffSize", 1024, "ws read buffer size")
	wBuffSize := flag.Int("wBuffSize", 4096, "ws write buffer size")
	broadcasterPoolSize := flag.Int("broadcasterPoolSize", 32, "broadcaster routines pool size")
	broadcastQueueSize := flag.Int("broadcastQueueSize", 4096, "broadcast channel queue size")
	maxMessageSize := flag.Int64("maxMessageSize", 1024, "max message size sent by peer")
	userOfflineTimeoutSec := flag.Int64("userOfflineTimeoutSec", 10, "user offline timeout in seconds")
	userWithoutIdTimeoutSec := flag.Int64("userWithoutIdTimeoutSec", 1, "user hasn't obtained id timeout in seconds")
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

func authConfig() auth.AuthenticatorConfig {
	googleClientId := os.Getenv("GOOGLE_CLIENT_ID")
	googleClientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")
	googleRedirectUrl := os.Getenv("HOST_URL") + "/google/oauth2/callback"
	httpRequestTimeoutSec := flag.Int64("httpRequestTimeoutSec", 5, "HTTP request timeout in seconds")
	flag.Parse()

	return auth.AuthenticatorConfig{
		HttpRequestTimeoutSec: *httpRequestTimeoutSec,
		GoogleClientId:        googleClientId,
		GoogleClientSecret:    googleClientSecret,
		GoogleRedirectUrl:     googleRedirectUrl,
	}
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./public")))

	authConfig := authConfig()
	auth := auth.NewAuthenticator(authConfig)
	http.HandleFunc("/google/oauth2", auth.HandleGoogleAuth2)
	http.HandleFunc("/google/oauth2/callback", auth.HandleGoogleAuth2Callback)

	hubConfig := hubConfig()
	cntrl := controller.NewGameController()
	hub := session.NewHub(hubConfig, cntrl)
	hub.Start()
}
