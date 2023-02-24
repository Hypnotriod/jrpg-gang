package main

import (
	"flag"
	"jrpg-gang/auth"
	"jrpg-gang/controller"
	"jrpg-gang/persistance"
	"jrpg-gang/session"
	"jrpg-gang/util"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)

func hubConfig() session.HubConfig {
	port := util.Getenv("PORT", "8080")
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
	if len(googleClientId) == 0 || len(googleClientSecret) == 0 {
		log.Fatal("Google credentials are not specified for the environment")
	}
	googleRedirectUrl := util.Getenv("HOST_URL", "http://localhost:8080") + "/google/oauth2/callback"
	authRequestTimeoutSec := flag.Int64("authRequestTimeoutSec", 10, "HTTP request timeout in seconds")
	authStateCacheTimeoutMin := flag.Int64("authStateCacheTimeoutMin", 10, "State cache timeout in minutes")
	flag.Parse()

	return auth.AuthenticatorConfig{
		RequestTimeoutSec:    *authRequestTimeoutSec,
		StateCacheTimeoutMin: *authStateCacheTimeoutMin,
		GoogleClientId:       googleClientId,
		GoogleClientSecret:   googleClientSecret,
		GoogleRedirectUrl:    googleRedirectUrl,
	}
}

func dbConfig() persistance.MongoDBConfig {
	mongodbUri := util.Getenv("MONGODB_URI", "mongodb://localhost:27017")
	requestTimeoutSec := flag.Int64("dbRequestTimeoutSec", 10, "Database request timeout seconds")
	flag.Parse()
	return persistance.MongoDBConfig{
		Uri:               mongodbUri,
		RequestTimeoutSec: *requestTimeoutSec,
	}
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./public")))

	dbConfig := dbConfig()
	persistance := persistance.NewPersistance(dbConfig)

	hubConfig := hubConfig()
	controller := controller.NewGameController(persistance)

	authConfig := authConfig()
	auth := auth.NewAuthenticator(authConfig, controller)
	http.HandleFunc("/google/oauth2", auth.HandleGoogleAuth2)
	http.HandleFunc("/google/oauth2/callback", auth.HandleGoogleAuth2Callback)

	hub := session.NewHub(hubConfig, controller)
	hub.Start()
}
