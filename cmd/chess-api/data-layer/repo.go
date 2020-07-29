package data

import (
	"errors"

	"github.com/jackc/pgx/v4"
	data "github.com/number7/chess-api/cmd/chess-api/data-layer/structs"
)

type Repo struct {
	pgClient *pgx.Conn
}

// TODO: gonna need a client that talks to Postgres here
func NewRepo(conn *pgx.Conn) Repo {
	return Repo{
		pgClient: conn,
	}
}

func (r Repo) Get(string) (data.GameData, error) {
	return data.GameData{}, errors.New("not implemented")
}
