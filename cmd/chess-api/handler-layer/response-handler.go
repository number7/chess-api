package rest

import (
	"encoding/json"
	"errors"
	"net/http"
)

var ResourceDoesNotExist = errors.New("this resource does not exist")

func handleError(w http.ResponseWriter, err error) {
	switch err {
	case ResourceDoesNotExist:
		http.Error(w, err.Error(), http.StatusNotFound)
	default:
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Respond(writer http.ResponseWriter, request *http.Request, response interface{}, err error) {
	if err != nil {
		handleError(writer, err)
		return
	}

	b, err := json.Marshal(response)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	if status, ok := request.Context().Value("Status").(int); ok {
		writer.WriteHeader(status)
	}

	writer.Write(b)
}
