package service

import (
	"errors"

	"github.com/number7/chess-api/cmd/chess-api/service-layer/structs"
)

type ChessRepo interface {
}

type Service struct {
	repo ChessRepo
}

func NewChessService(repo ChessRepo) Service {
	return Service{
		repo: repo,
	}
}

func (s Service) Get(req structs.GetRequest) (structs.GetResponse, error) {
	return structs.GetResponse{}, errors.New("method not implemented")
}
