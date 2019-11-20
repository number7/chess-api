package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"

	"gkwkr/chess-api/cmd/chess-api/config"
	rest "gkwkr/chess-api/cmd/chess-api/handlers"
)

func main() {
	if len(os.Args) != 2 {
		logrus.Fatalln("you must provide a single argument representing the config file.")
	}
	appConfig, err := config.ParseConfig(os.Args[1])
	if err != nil {
		panic(fmt.Sprintf("Error: Failed to parse the config file. Error returned: \n%v", err))
	}

	address := fmt.Sprintf("%v:%v", appConfig.RestServer.Host, appConfig.RestServer.Port)

	restHandler := rest.Handler{}

	router := chi.NewRouter()
	router.Route(rest.ResourcePath, restHandler.Routes)

	err = http.ListenAndServe(address, router)
	if err != nil {
		fmt.Printf("error serving HTTP site:\n%v", err)
	}
}
