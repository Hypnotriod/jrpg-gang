package session

import (
	"io"
	"jrpg-gang/auth"
	"jrpg-gang/controller"
	"net/http"
)

const MaxPayloadSize = 256

func InitRoutes(controller *controller.GameController, auth *auth.Authenticator) {
	http.HandleFunc("GET /google/oauth2", auth.HandleGoogleAuth2)
	http.HandleFunc("GET /google/oauth2/callback", auth.HandleGoogleAuth2Callback)
	http.HandleFunc("POST /configuration", handleConfigurationRequest(controller))
}

func handleConfigurationRequest(controller *controller.GameController) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		r.Body = http.MaxBytesReader(w, r.Body, MaxPayloadSize)
		if data, err := io.ReadAll(r.Body); err == nil {
			w.Write(controller.HandleConfigurationRequest(data))
		} else {
			http.Error(w, "", http.StatusBadRequest)
		}
	}
}
