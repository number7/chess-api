package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/number7/chess-api/cmd/chess-api/config"
	"github.com/number7/chess-api/cmd/chess-api/data-layer"
	rest "github.com/number7/chess-api/cmd/chess-api/handler-layer"
	"github.com/number7/chess-api/cmd/chess-api/service-layer"
	"github.com/sirupsen/logrus"
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

	chessRepo := data.NewRepo()
	svc := service.NewChessService(chessRepo)
	restHandler := rest.NewRestHandler(svc)

	router := chi.NewRouter()
	router.Route(rest.ResourcePath, restHandler.Routes)

	fmt.Print("Serving up chess games, dude...")
	err = http.ListenAndServe(address, router)
	if err != nil {
		fmt.Printf("error serving HTTP site:\n%v", err)
	}
}
