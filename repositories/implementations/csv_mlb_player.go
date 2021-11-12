package repositories

import (
	"encoding/csv"
	"errors"
	"log"
	"os"
	"strconv"

	e "github.com/EloYaniel/academy-go-q42021/entities"
)

// CSVMLBPlayerRepository implements MLBPlayerRepository interface
type CSVMLBPlayerRepository struct{}

func (repo *CSVMLBPlayerRepository) GetMLBPlayers() ([]e.MLBPlayer, error) {
	f, err := os.Open("data/mlb_players.csv")

	if err != nil {
		log.Println("error opening the file", err)
		return nil, errors.New("error opening the file")
	}
	defer f.Close()
	data, err := csv.NewReader(f).ReadAll()

	if err != nil {
		log.Println("error reading the file: ", err)
		return nil, errors.New("error reading the file")
	}
	var players []e.MLBPlayer
	for i, line := range data {
		if i != 0 {
			id, _ := strconv.Atoi(line[0])
			height, _ := strconv.Atoi(line[4])
			weight, _ := strconv.ParseFloat(line[5], 32)
			age, _ := strconv.ParseFloat(line[6], 32)

			players = append(players, e.MLBPlayer{
				ID:       id,
				Name:     line[1],
				Team:     line[2],
				Position: line[3],
				Height:   height,
				Weight:   float32(weight),
				Age:      float32(age),
			})
		}

	}

	return players, nil
}

func (repo *CSVMLBPlayerRepository) GetMLBPlayerByID(id int) (*e.MLBPlayer, error) {
	players, err := repo.GetMLBPlayers()

	if err != nil {
		return nil, errors.New("error getting player")
	}

	for _, p := range players {
		if p.ID == id {
			return &p, nil
		}
	}

	return &e.MLBPlayer{}, nil
}