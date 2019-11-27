package rest

import (
	"net/http"

	"github.com/number7/chess-api/cmd/chess-api/service-layer/structs"
)

type GetService interface {
	Get(structs.GetRequest) (structs.GetResponse, error)
}
type GetHandler struct {
	service GetService
}

func (h GetHandler) Get(writer http.ResponseWriter, req *http.Request) {
}
