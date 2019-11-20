package rest

import (
	"github.com/go-chi/chi"
	"net/http"
)

type Handler struct {
}

const ResourcePath = "/chess-api/v1alpha"

func (h Handler) Routes(router chi.Router) {
	router.Get("/heartbeat", h.heartbeat)
}

func (h Handler) heartbeat(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("I'm alive."))
}
