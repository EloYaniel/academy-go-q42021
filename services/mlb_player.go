package services

import (
	"log"

	e "github.com/EloYaniel/academy-go-q42021/entities"
	r "github.com/EloYaniel/academy-go-q42021/repositories/contracts"
)

type MLBPlayerService struct {
	repository r.MLBPlayerRepository
}

func NewMLBPlayerService(r r.MLBPlayerRepository) *MLBPlayerService {
	return &MLBPlayerService{repository: r}
}

func (s *MLBPlayerService) GetMLBPlayers() ([]e.MLBPlayer, error) {
	players, err := s.repository.GetMLBPlayers()

	if err != nil {
		log.Println(err)
	}

	return players, err
}

func (s *MLBPlayerService) GetMLBPlayerByID(id int) (*e.MLBPlayer, error) {
	player, err := s.repository.GetMLBPlayerByID(id)

	if err != nil {
		log.Println(err)
	}

	return player, err
}
