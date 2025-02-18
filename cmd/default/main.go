package main

import (
	"flag"
	"io"
	"jrpg-gang/auth"
	"jrpg-gang/controller"
	"jrpg-gang/persistance"
	"jrpg-gang/session"
	"jrpg-gang/util"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)

func configs() (authConfig auth.AuthenticatorConfig, hubConfig session.HubConfig, persistanceConfig persistance.PersistanceConfig) {
	googleClientId := os.Getenv("GOOGLE_CLIENT_ID")
	googleClientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")
	if len(googleClientId) == 0 || len(googleClientSecret) == 0 {
		log.Fatal("Google credentials are not specified for the environment")
	}
	googleAuthCallbackUrl := util.Getenv("HOST_URL", "http://localhost:8080") + "/google/oauth2/callback"
	authRedirectUrl := util.Getenv("AUTH_REDIRECT_URL", "http://localhost:8080")
	authRequestTimeoutSec := flag.Int64("authRequestTimeoutSec", 10, "HTTP request timeout in seconds")
	authStateCacheTimeoutMin := flag.Int64("authStateCacheTimeoutMin", 30, "State cache timeout in minutes")

	port := util.Getenv("PORT", "8080")
	key := flag.String("key", "", "path to TLS key.pem")
	cert := flag.String("cert", "", "path to TLS cert.pem")
	rBuffSize := flag.Int("rBuffSize", 1024, "ws read buffer size")
	wBuffSize := flag.Int("wBuffSize", 4096, "ws write buffer size")
	broadcasterPoolSize := flag.Int("broadcasterPoolSize", 32, "broadcaster routines pool size")
	broadcastQueueSize := flag.Int("broadcastQueueSize", 4096, "broadcast channel queue size")
	maxMessageSize := flag.Int64("maxMessageSize", 1024, "max message size sent by peer")
	userOfflineTimeoutSec := flag.Int64("userOfflineTimeoutSec", 10, "user offline timeout in seconds")
	userWithoutIdTimeoutSec := flag.Int64("userWithoutIdTimeoutSec", 2, "user hasn't obtained id timeout in seconds")
	pingTimeoutSec := flag.Int64("pingTimeoutSec", 30, "server ping timeout in seconds")
	readDeadlineSec := flag.Int64("readDeadlineSec", 35, "client read deadline timeout in seconds (including pong messages)")
	writeDeadlineSec := flag.Int64("writeDeadlineSec", 2, "client write deadline timeout in seconds")

	persistanceCacheTimeoutMin := flag.Int64("persistanceCacheTimeoutMin", 30, "Persistance cache timeout in minutes")
	mongodbUri := util.Getenv("MONGODB_URI", "mongodb://localhost:27017")
	requestTimeoutSec := flag.Int64("dbRequestTimeoutSec", 10, "Database request timeout seconds")

	flag.Parse()

	authConfig = auth.AuthenticatorConfig{
		RequestTimeoutSec:    *authRequestTimeoutSec,
		StateCacheTimeoutMin: *authStateCacheTimeoutMin,
		GoogleClientId:       googleClientId,
		GoogleClientSecret:   googleClientSecret,
		GoogleCallbackUrl:    googleAuthCallbackUrl,
		RedirectUrl:          authRedirectUrl,
	}

	hubConfig = session.HubConfig{
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
		PingTimeoutSec:          *pingTimeoutSec,
		ReadDeadlineSec:         *readDeadlineSec,
		WriteDeadlineSec:        *writeDeadlineSec,
	}

	persistanceConfig = persistance.PersistanceConfig{
		PersistanceCacheTimeoutMin: *persistanceCacheTimeoutMin,
		MongoDBConfig: persistance.MongoDBConfig{
			Uri:               mongodbUri,
			RequestTimeoutSec: *requestTimeoutSec,
		},
	}

	return
}

func main() {
	authConfig, hubConfig, persistanceConfig := configs()
	persistance := persistance.NewPersistance(persistanceConfig)
	controller := controller.NewGameController(persistance)
	auth := auth.NewAuthenticator(authConfig, controller)
	http.HandleFunc("/google/oauth2", auth.HandleGoogleAuth2)
	http.HandleFunc("/google/oauth2/callback", auth.HandleGoogleAuth2Callback)
	http.HandleFunc("/configuration", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		data, _ := io.ReadAll(r.Body)
		w.Write(controller.HandleConfigurationRequest(data))
	})

	hub := session.NewHub(hubConfig, controller)
	hub.Start()
}
