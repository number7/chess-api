package rest

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/number7/chess-api/cmd/chess-api/handler-layer/structs"
)

const ResourcePath = "/chess-api/v1alpha"

type ChessService interface {
	GetService
}

type Handler struct {
	GetHandler
}

func NewRestHandler(service ChessService) Handler {
	return Handler{
		GetHandler: GetHandler{service: service},
	}
}

func (h Handler) Routes(router chi.Router) {
	router.Get("/heartbeat", h.heartbeat)
	router.Get("/", h.Get)
}

func (h Handler) heartbeat(w http.ResponseWriter, r *http.Request) {
	resp := structs.Heartbeat{
		Message:   "I'm alive",
		Timestamp: time.Now(),
	}
	Respond(w, r, resp, nil)
}
