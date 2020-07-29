package clients

import (
	"errors"

	"github.com/jackc/pgx"
	"github.com/number7/chess-api/cmd/chess-api/config"
)

type PostgresClient struct {
	conn pgx.Conn
}

func (c PostgresClient) RunQuery(dbinfo config.DBInfo) error {
	return errors.New("i ain't implemented")
}
