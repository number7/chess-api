package data

import (
	"errors"

	data "github.com/number7/chess-api/cmd/chess-api/data-layer/structs"
)

type Repo struct {
}

// TODO: gonna need a client that talks to Postgres here
func NewRepo() Repo {
	return Repo{}
}

func (r Repo) Get(string) (data.GameData, error) {
	return data.GameData{}, errors.New("not implemented")
}
