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
	var getReq structs.GetRequest

	resp, err := h.service.Get(getReq)
	if err != nil {
		Respond(writer, req, nil, err)
	}

	Respond(writer, req, resp, nil)
}
