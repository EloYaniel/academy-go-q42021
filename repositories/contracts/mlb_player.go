package repositories

import e "github.com/EloYaniel/academy-go-q42021/entities"

type MLBPlayerRepository interface {
	// GetMLBPlayers gets all MLB Players.
	GetMLBPlayers() ([]e.MLBPlayer, error)

	// GetMLBPlayerByID get a Player by its ID
	GetMLBPlayerByID(id int) (*e.MLBPlayer, error)

	// GetMLBPlayerDesired gets MLB Players and filetered by its params.
	GetMLBPlayerDesired(filterType string, totalItems int, itemsPerWorker int) ([]e.MLBPlayer, error)
}
