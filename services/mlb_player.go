package services

import (
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
	return s.repository.GetMLBPlayers()
}

func (s *MLBPlayerService) GetMLBPlayerByID(id int) (*e.MLBPlayer, error) {
	return s.repository.GetMLBPlayerByID(id)
}
