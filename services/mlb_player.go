package services

import (
	"log"

	e "github.com/EloYaniel/academy-go-q42021/entities"
	r "github.com/EloYaniel/academy-go-q42021/repositories/contracts"
)

// MLBPlayerService struct handles MLB Players business logic.
type MLBPlayerService struct {
	repository r.MLBPlayerRepository
}

// NewMLBPlayerService function return an instance of MLBPlayerService
func NewMLBPlayerService(r r.MLBPlayerRepository) *MLBPlayerService {
	return &MLBPlayerService{repository: r}
}

// GetMLBPlayers gets all MLB Players.
func (s *MLBPlayerService) GetMLBPlayers() ([]e.MLBPlayer, error) {
	players, err := s.repository.GetMLBPlayers()

	if err != nil {
		log.Println(err)
	}

	return players, err
}

// GetMLBPlayerByID get a Player by its ID
func (s *MLBPlayerService) GetMLBPlayerByID(id int) (*e.MLBPlayer, error) {
	player, err := s.repository.GetMLBPlayerByID(id)

	if err != nil {
		log.Println(err)
	}

	return player, err
}

// GetMLBPlayerDesired gets MLB Players and filetered by its params.
func (s *MLBPlayerService) GetMLBPlayerDesired(filterType string, totalItems int, itemsPerWorker int) ([]e.MLBPlayer, error) {
	players, err := s.repository.GetMLBPlayerDesired(filterType, totalItems, itemsPerWorker)

	if err != nil {
		log.Println(err)
	}

	return players, err
}
