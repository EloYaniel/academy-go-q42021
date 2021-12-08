package repositories

import (
	"errors"
	"testing"

	e "github.com/EloYaniel/academy-go-q42021/entities"
	"github.com/stretchr/testify/assert"
)

var player1 = e.MLBPlayer{
	ID:       1,
	Name:     "Adam Donachie",
	Team:     "BAL",
	Position: "Catcher",
	Height:   74,
	Weight:   180,
	Age:      22.99,
}

var player2 = e.MLBPlayer{
	ID:       2,
	Name:     "Paul Bako",
	Team:     "BAL",
	Position: "Catcher",
	Height:   74,
	Weight:   215,
	Age:      34.69,
}

func Test_CSVMLBPlayerRepository_ShouldReturnDiffInstances(t *testing.T) {
	instance := NewCSVMLBPlayerRepository("here.csv")
	instance2 := NewCSVMLBPlayerRepository("there.csv")

	assert.NotNil(t, instance)
	assert.NotNil(t, instance2)
	assert.NotSame(t, instance, instance2)
}

func Test_GetMLBPlayers_Suite(t *testing.T) {
	players := []e.MLBPlayer{
		player1,
		player2,
	}
	testCases := []struct {
		name             string
		filePath         string
		expectedError    error
		expectedResponse []e.MLBPlayer
	}{
		{
			name:             "Should return the players",
			filePath:         "../../data/test/players-test.csv",
			expectedResponse: players,
			expectedError:    nil,
		},
		{
			name:             "Should return error when open file",
			filePath:         "",
			expectedResponse: nil,
			expectedError:    errors.New("error opening the file"),
		},
		{
			name:             "Should return error when casting ID",
			filePath:         "../../data/test/players-with-wrong-id-test.csv",
			expectedResponse: nil,
			expectedError:    errors.New("error casting ID"),
		},
		{
			name:             "Should return error when casting Height",
			filePath:         "../../data/test/players-with-wrong-height-test.csv",
			expectedResponse: nil,
			expectedError:    errors.New("error casting Height"),
		},
		{
			name:             "Should return error when casting Weight",
			filePath:         "../../data/test/players-with-wrong-weight-test.csv",
			expectedResponse: nil,
			expectedError:    errors.New("error casting Weight"),
		},
		{
			name:             "Should return error when casting Age",
			filePath:         "../../data/test/players-with-wrong-age-test.csv",
			expectedResponse: nil,
			expectedError:    errors.New("error casting Age"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			repo := NewCSVMLBPlayerRepository(tc.filePath)

			users, err := repo.GetMLBPlayers()

			assert.Equal(t, tc.expectedResponse, users)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func Test_GetPlayerByID_Suite(t *testing.T) {
	players := []e.MLBPlayer{
		player1,
		player2,
	}
	testCases := []struct {
		name             string
		filePath         string
		playerID         int
		players          []e.MLBPlayer
		expectedError    error
		expectedResponse *e.MLBPlayer
	}{
		{
			name:             "Should return the player",
			filePath:         "../../data/test/players-test.csv",
			playerID:         1,
			players:          players,
			expectedResponse: &player1,
			expectedError:    nil,
		},
		{
			name:             "Should return no player and no error",
			filePath:         "../../data/test/players-test.csv",
			players:          players,
			playerID:         3,
			expectedResponse: nil,
			expectedError:    nil,
		},
		{
			name:             "Should return no player and error",
			filePath:         "",
			players:          players,
			playerID:         1,
			expectedResponse: nil,
			expectedError:    errors.New("error getting player"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repo := NewCSVMLBPlayerRepository(tc.filePath)

			player, err := repo.GetMLBPlayerByID(tc.playerID)

			assert.Equal(t, tc.expectedResponse, player)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func Test_GetMLBPlayerDesired_Suite(t *testing.T) {
	testCases := []struct {
		name             string
		filter           string
		totalItems       int
		itemsPerWorker   int
		filePath         string
		expectedError    error
		expectedResponse []e.MLBPlayer
	}{
		{
			name:           "Should return players with odd ID",
			filePath:       "../../data/mlb_players.csv",
			filter:         "odd",
			itemsPerWorker: 1,
			totalItems:     2,
			expectedResponse: []e.MLBPlayer{
				{
					ID:       1,
					Name:     "Adam Donachie",
					Team:     "BAL",
					Position: "Catcher",
					Height:   74,
					Weight:   180,
					Age:      22.99,
				},
				{

					ID:       3,
					Name:     "Ramon Hernandez",
					Team:     "BAL",
					Position: "Catcher",
					Height:   72,
					Weight:   210,
					Age:      30.78,
				},
			},
			expectedError: nil,
		},
		{
			name:           "Should return players with even ID",
			filePath:       "../../data/mlb_players.csv",
			filter:         "even",
			itemsPerWorker: 1,
			totalItems:     2,
			expectedResponse: []e.MLBPlayer{
				{
					ID:       2,
					Name:     "Paul Bako",
					Team:     "BAL",
					Position: "Catcher",
					Height:   74,
					Weight:   215,
					Age:      34.69,
				},
				{

					ID:       4,
					Name:     "Kevin Millar",
					Team:     "BAL",
					Position: "First Baseman",
					Height:   72,
					Weight:   210,
					Age:      35.43,
				},
			},
			expectedError: nil,
		},
		{
			name:             "Should return error when open file",
			filePath:         "",
			expectedResponse: nil,
			expectedError:    errors.New("error opening the file"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			repo := NewCSVMLBPlayerRepository(tc.filePath)

			users, err := repo.GetMLBPlayerDesired(tc.filter, tc.totalItems, tc.itemsPerWorker)

			assert.Equal(t, tc.expectedResponse, users)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}
