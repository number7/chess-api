package service

type ChessRepo interface {
	GetRepo
}

type Service struct {
	GetService
}

func NewChessService(repo ChessRepo) Service {
	return Service{
		GetService: GetService{
			repo: repo,
		},
	}
}
