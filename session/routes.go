package session

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

func (h *Hub) initRoutes(router *mux.Router) {
	router.HandleFunc("/google/oauth2", h.auth.HandleGoogleAuth2).Methods(http.MethodGet)
	router.HandleFunc("/google/oauth2/callback", h.auth.HandleGoogleAuth2Callback).Methods(http.MethodGet)
	router.HandleFunc("/configuration", h.handleConfigurationRequest).
		Methods(http.MethodPost, http.MethodOptions).Headers(HeaderContentType, ContentTypeApplicationJson)
}

func (h *Hub) handleConfigurationRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Add(HeaderContentType, ContentTypeApplicationJson)
	r.Body = http.MaxBytesReader(w, r.Body, ConfigurationPayloadMaxSize)
	if data, err := io.ReadAll(r.Body); err == nil {
		w.Write(h.controller.HandleConfigurationRequest(data))
	} else {
		http.Error(w, "", http.StatusBadRequest)
	}
}
