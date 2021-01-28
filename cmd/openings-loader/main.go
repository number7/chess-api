package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/jackc/pgx/v4"
)

const ErrorStr = "ERROR"

type appConfig struct {
	File     string
	Host     string
	Port     int
	Username string
	Password string
	DBName   string
}

var cfg appConfig

func init() {
	parseFlags(&cfg)
}

func main() {
	mustValidateFlags(cfg)

	ctx := context.Background()
	url := makePostgresURL(cfg)

	conn, err := pgx.Connect(ctx, url)
	if err != nil {
		panic(err)
	}
	defer conn.Close(ctx)
}

func parseFlags(cfg *appConfig) {
	flag.StringVar(&cfg.File, "file", ErrorStr, "Name and path of file to parse")
	flag.StringVar(&cfg.Host, "host", ErrorStr, "Name and path of file to parse")
	flag.IntVar(&cfg.Port, "port", 0, "Name and path of file to parse")
	flag.StringVar(&cfg.Username, "uid", ErrorStr, "Name and path of file to parse")
	flag.StringVar(&cfg.Password, "pwd", ErrorStr, "Name and path of file to parse")
	flag.StringVar(&cfg.DBName, "dbname", ErrorStr, "Name and path of file to parse")

	flag.Parse()
}

func mustValidateFlags(cfg appConfig) {
	var badParms []string

	if cfg.File == ErrorStr {
		badParms = append(badParms, "file")
	}
	if cfg.Host == ErrorStr {
		badParms = append(badParms, "host")
	}
	if cfg.Port == 0 {
		badParms = append(badParms, "port")
	}
	if cfg.Username == ErrorStr {
		badParms = append(badParms, "uid")
	}
	if cfg.Password == ErrorStr {
		badParms = append(badParms, "pwd")
	}
	if cfg.DBName == ErrorStr {
		badParms = append(badParms, "dbname")
	}

	if len(badParms) > 0 {
		fmt.Printf("These required flags were not passed in: %s\n\n", strings.Join(badParms, ", "))
		flag.Usage()
		os.Exit(1)
	}
}

func makePostgresURL(info appConfig) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s", info.Username, info.Password, info.Host, info.Port, info.DBName)
}
