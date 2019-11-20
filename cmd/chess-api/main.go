package main

import (
    "fmt"
    "github.com/go-chi/chi"
    rest "gkwkr/chess-api/cmd/chess-api/handlers"
    "net/http"
)

func main() {
    host := "localhost"  // put in a config
    port := "1066"       // put in a config
    address := fmt.Sprintf("%v:%v", host, port)

    restHandler := rest.Handler{}

    router := chi.NewRouter()
    router.Route(rest.ResourcePath, restHandler.Routes)
    http.ListenAndServe(address, router)
}
