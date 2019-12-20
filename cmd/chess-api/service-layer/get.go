package service

import (
	data "github.com/number7/chess-api/cmd/chess-api/data-layer/structs"
	"github.com/number7/chess-api/cmd/chess-api/service-layer/structs"
)

//go:generate counterfeiter . GetRepo
type GetRepo interface {
	Get(string) (data.GameData, error)
}

type GetService struct {
	repo GetRepo
}

func New(repo GetRepo) GetService {
	return GetService{
		repo: repo,
	}
}

func (s GetService) Get(req structs.GetRequest) (structs.GetResponse, error) {
	_, err := s.repo.Get(req.ID)
	if err != nil {
		return structs.GetResponse{}, err
	}

	resp := structs.GetResponse{
		ID:        "1",
		Location:  "Helsinki",
		Date:      "20190610 120000 -5:00",
		Round:     1,
		WhiteName: "Elvis Hitler",
		BlackName: "Casanova Frankenstein",
		Moves:     "that kill",
		Result:    "dance-off",
	}
	return resp, nil
}
