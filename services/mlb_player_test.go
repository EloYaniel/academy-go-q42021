package services

import (
	"errors"
	"testing"

	e "github.com/EloYaniel/academy-go-q42021/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockMLBPlayerRepository struct {
	mock.Mock
}

func (m *mockMLBPlayerRepository) GetMLBPlayers() ([]e.MLBPlayer, error) {
	args := m.Called()

	return args.Get(0).([]e.MLBPlayer), args.Error(1)
}

func (m *mockMLBPlayerRepository) GetMLBPlayerByID(id int) (*e.MLBPlayer, error) {
	args := m.Called()

	return args.Get(0).(*e.MLBPlayer), args.Error(1)
}

func (m *mockMLBPlayerRepository) GetMLBPlayerDesired(filterType string, totalItems int, itemsPerWorker int) ([]e.MLBPlayer, error) {
	args := m.Called()

	return args.Get(0).([]e.MLBPlayer), args.Error(1)
}

func Test_NewMLBPlayerService_ShouldReturnInstance(t *testing.T) {
	instance := NewMLBPlayerService(&mockMLBPlayerRepository{})
	instance2 := NewMLBPlayerService(&mockMLBPlayerRepository{})

	assert.NotNil(t, instance)
	assert.NotSame(t, instance, instance2)
}

func Test_GetMLBPlayers_Suite(t *testing.T) {
	testCases := []struct {
		name     string
		response []e.MLBPlayer
		err      error
	}{
		{
			name:     "Should return error when repo has error",
			response: nil,
			err:      errors.New("Error getting players"),
		},
		{name: "Should return players",
			response: []e.MLBPlayer{
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
					ID:       2,
					Name:     "Paul Bako",
					Team:     "BAL",
					Position: "Catcher",
					Height:   74,
					Weight:   215,
					Age:      34.69,
				},
			},
			err: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repoMock := new(mockMLBPlayerRepository)
			repoMock.On("GetMLBPlayers").Return(tc.response, tc.err)
			service := NewMLBPlayerService(repoMock)

			resp, err := service.GetMLBPlayers()

			if err != nil {
				assert.Equal(t, tc.err, err)
				assert.Nil(t, resp)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tc.response, resp)
			}

		})
	}
}

func Test_GetMLBPlayerByID_Suite(t *testing.T) {
	testCases := []struct {
		name     string
		response *e.MLBPlayer
		err      error
	}{
		{
			name:     "Should return error when repo has error",
			response: nil,
			err:      errors.New("Error getting player"),
		},
		{
			name: "Should return players",
			response: &e.MLBPlayer{
				ID:       1,
				Name:     "Adam Donachie",
				Team:     "BAL",
				Position: "Catcher",
				Height:   74,
				Weight:   180,
				Age:      22.99,
			},
			err: nil,
		},
		{
			name:     "Should return no error nor player if not found",
			response: nil,
			err:      nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repoMock := new(mockMLBPlayerRepository)
			repoMock.On("GetMLBPlayerByID").Return(tc.response, tc.err)
			service := NewMLBPlayerService(repoMock)

			resp, err := service.GetMLBPlayerByID(1)

			if err != nil {
				assert.Equal(t, tc.err, err)
			}
			if resp != nil {
				assert.Equal(t, tc.response, resp)
			}

		})
	}
}

func Test_GetMLBPlayerDesired_Suite(t *testing.T) {
	testCases := []struct {
		name     string
		response []e.MLBPlayer
		err      error
	}{
		{
			name:     "Should return error when repo has error",
			response: nil,
			err:      errors.New("Error getting players"),
		},
		{
			name: "Should return players",
			response: []e.MLBPlayer{
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
					ID:       2,
					Name:     "Paul Bako",
					Team:     "BAL",
					Position: "Catcher",
					Height:   74,
					Weight:   215,
					Age:      34.69,
				},
			},
			err: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repoMock := new(mockMLBPlayerRepository)
			repoMock.On("GetMLBPlayerDesired").Return(tc.response, tc.err)
			service := NewMLBPlayerService(repoMock)

			resp, err := service.GetMLBPlayerDesired("even", 20, 5)

			if err != nil {
				assert.Equal(t, tc.err, err)
				assert.Nil(t, resp)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tc.response, resp)
			}

		})
	}
}
