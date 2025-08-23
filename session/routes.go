package session

import (
	"io"
	"jrpg-gang/util"
	"net/http"

	"github.com/gorilla/mux"
)

func (h *Hub) initRoutes(router *mux.Router) {
	router.HandleFunc("/google/oauth2", h.auth.HandleGoogleAuth2).Methods(http.MethodGet)
	router.HandleFunc("/google/oauth2/callback", h.auth.HandleGoogleAuth2Callback).Methods(http.MethodGet)
	router.HandleFunc("/configuration", h.handleConfigurationRequest).Methods(http.MethodPost).
		Headers(HeaderContentType, ContentTypeApplicationJson)
	router.NotFoundHandler = http.HandlerFunc(h.handleNotFoundRequest)
}

func (h *Hub) handleConfigurationRequest(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, ConfigurationPayloadMaxSize)
	if data, err := io.ReadAll(r.Body); err == nil {
		w.Header().Add(HeaderContentType, ContentTypeApplicationJson)
		w.Write(h.controller.HandleConfigurationRequest(util.GetIP(r), data))
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func (h *Hub) handleNotFoundRequest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}
