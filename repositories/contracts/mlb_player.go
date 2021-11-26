package repositories

import e "github.com/EloYaniel/academy-go-q42021/entities"

type MLBPlayerRepository interface {
	GetMLBPlayers() ([]e.MLBPlayer, error)
	GetMLBPlayerByID(id int) (*e.MLBPlayer, error)
}
