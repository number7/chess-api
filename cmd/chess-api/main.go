package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v4"
	"github.com/jackc/pgx/v4"
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

	ctx := context.Background()
	address := fmt.Sprintf("%v:%v", appConfig.RestServer.Host, appConfig.RestServer.Port)

	url := makePostgresURL(appConfig.DBInfo)
	conn, err := pgx.Connect(ctx, url)
	if err != nil {
		panic(err)
	}
	defer conn.Close(ctx)

	chessRepo := data.NewRepo(conn)
	svc := service.NewChessService(chessRepo)
	restHandler := rest.NewRestHandler(svc)

	router := chi.NewRouter()
	router.Route(rest.ResourcePath, restHandler.Routes)

	fmt.Print("Serving up chess games, dude...\n")
	fmt.Printf("on http://%v:%v\n", appConfig.RestServer.Host, appConfig.RestServer.Port)
	err = http.ListenAndServe(address, router)
	if err != nil {
		fmt.Printf("error serving HTTP site:\n%v", err)
	}
}

func makePostgresURL(info config.DBInfo) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s", info.Username, info.Password, info.Host, info.Port, info.DBName)
}
