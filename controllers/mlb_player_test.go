package controllers

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	e "github.com/EloYaniel/academy-go-q42021/entities"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockMLBService struct {
	mock.Mock
}

func (m *mockMLBService) GetMLBPlayers() ([]e.MLBPlayer, error) {
	args := m.Called()

	return args.Get(0).([]e.MLBPlayer), args.Error(1)
}

func (m *mockMLBService) GetMLBPlayerByID(id int) (*e.MLBPlayer, error) {
	args := m.Called()

	return args.Get(0).(*e.MLBPlayer), args.Error(1)
}

func (m *mockMLBService) GetMLBPlayerDesired(filterType string, totalItems int, itemsPerWorker int) ([]e.MLBPlayer, error) {
	args := m.Called()

	return args.Get(0).([]e.MLBPlayer), args.Error(1)
}

func Test_MLBPlayerController_GetMLBPlayers_Suite(t *testing.T) {
	testCases := []struct {
		name                 string
		statusCode           int
		expectedServiceCalls int
		hasError             bool
		serviceError         error
		serviceResponse      []e.MLBPlayer
		errorMessage         string
	}{
		{
			name:                 "Should return players",
			statusCode:           http.StatusOK,
			expectedServiceCalls: 1,
			hasError:             false,
			serviceResponse: []e.MLBPlayer{
				{
					ID:       1,
					Name:     "Adam Donachie",
					Team:     "BAL",
					Position: "Catcher",
					Height:   74,
					Weight:   180,
					Age:      22.99,
				},
			},
			serviceError: nil,
		},
		{
			name:                 "Should return internal server on service error",
			statusCode:           http.StatusInternalServerError,
			expectedServiceCalls: 1,
			hasError:             true,
			serviceResponse:      nil,
			serviceError:         errors.New("unknown error"),
			errorMessage:         "Internal server error",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "/mlb-players", nil)
			m := new(mockMLBService)
			m.On("GetMLBPlayers").Return(tc.serviceResponse, tc.serviceError)
			ctr := NewMLBPlayerController(m)

			ctr.GetMLBPlayers(w, r)

			if tc.hasError {
				assert.Contains(t, w.Body.String(), tc.errorMessage)
			}

			assert.Equal(t, w.Code, tc.statusCode)
			assert.Equal(t, "application/json", w.Result().Header.Get("Content-Type"))
			m.AssertNumberOfCalls(t, "GetMLBPlayers", tc.expectedServiceCalls)
		})
	}
}

func Test_MLBPlayerController_GetMLBPlayerByID_Suite(t *testing.T) {
	testCases := []struct {
		name                 string
		idParam              string
		statusCode           int
		expectedServiceCalls int
		hasError             bool
		serviceError         error
		serviceResponse      *e.MLBPlayer
		errorMessage         string
	}{
		{
			name:                 "Should return player",
			idParam:              "1",
			statusCode:           http.StatusOK,
			expectedServiceCalls: 1,
			hasError:             false,
			serviceResponse: &e.MLBPlayer{
				ID:       1,
				Name:     "Adam Donachie",
				Team:     "BAL",
				Position: "Catcher",
				Height:   74,
				Weight:   180,
				Age:      22.99,
			},
			serviceError: nil,
		},
		{
			name:                 "Should return internal server on service error",
			idParam:              "1",
			statusCode:           http.StatusInternalServerError,
			expectedServiceCalls: 1,
			hasError:             true,
			serviceResponse:      nil,
			serviceError:         errors.New("unknown error"),
			errorMessage:         "Internal server error",
		},
		{
			name:                 "Should return not found if no player",
			idParam:              "10",
			statusCode:           http.StatusNotFound,
			expectedServiceCalls: 1,
			hasError:             false,
			serviceResponse:      nil,
			serviceError:         nil,
			errorMessage:         "User not found",
		},
		{
			name:                 "Should return bad request if no player id provided",
			idParam:              "1a2b",
			statusCode:           http.StatusBadRequest,
			expectedServiceCalls: 0,
			hasError:             true,
			serviceResponse:      nil,
			serviceError:         nil,
			errorMessage:         "Player ID provided must be of type integer",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "/mlb-player/{id}", nil)
			r = mux.SetURLVars(r, map[string]string{"id": tc.idParam})
			m := new(mockMLBService)
			m.On("GetMLBPlayerByID").Return(tc.serviceResponse, tc.serviceError)
			ctr := NewMLBPlayerController(m)

			ctr.GetMLBPlayerByID(w, r)
			res := w.Result()
			body, _ := ioutil.ReadAll(res.Body)

			if tc.hasError {
				assert.Contains(t, string(body), tc.errorMessage)
			}

			assert.Equal(t, w.Code, tc.statusCode)
			assert.Equal(t, "application/json", res.Header.Get("Content-Type"))
			m.AssertNumberOfCalls(t, "GetMLBPlayerByID", tc.expectedServiceCalls)
		})
	}
}

func Test_MLBPlayerController_GetMLBPlayerDesired_Suite(t *testing.T) {
	testCases := []struct {
		name                 string
		typeParam            string
		ipwParam             string
		itemsParam           string
		statusCode           int
		expectedServiceCalls int
		hasError             bool
		serviceError         error
		serviceResponse      []e.MLBPlayer
		errorMessage         string
	}{
		{
			name:                 "Should return players with odd type",
			typeParam:            "odd",
			ipwParam:             "5",
			itemsParam:           "20",
			statusCode:           http.StatusOK,
			expectedServiceCalls: 1,
			hasError:             false,
			serviceResponse: []e.MLBPlayer{
				{
					ID:       1,
					Name:     "Adam Donachie",
					Team:     "BAL",
					Position: "Catcher",
					Height:   74,
					Weight:   180,
					Age:      22.99,
				},
			},
			serviceError: nil,
		},
		{
			name:                 "Should return players with even type",
			typeParam:            "even",
			ipwParam:             "5",
			itemsParam:           "20",
			statusCode:           http.StatusOK,
			expectedServiceCalls: 1,
			hasError:             false,
			serviceResponse: []e.MLBPlayer{
				{
					ID:       1,
					Name:     "Adam Donachie",
					Team:     "BAL",
					Position: "Catcher",
					Height:   74,
					Weight:   180,
					Age:      22.99,
				},
			},
			serviceError: nil,
		},
		{
			name:                 "Should return internal server on service error",
			typeParam:            "odd",
			ipwParam:             "5",
			itemsParam:           "20",
			statusCode:           http.StatusInternalServerError,
			expectedServiceCalls: 1,
			hasError:             true,
			serviceResponse:      nil,
			serviceError:         errors.New("unknown error"),
			errorMessage:         "Internal server error",
		},
		{
			name:                 "Should return bad request if not allowed type",
			typeParam:            "even-odd",
			ipwParam:             "5",
			itemsParam:           "20",
			statusCode:           http.StatusBadRequest,
			expectedServiceCalls: 0,
			hasError:             true,
			serviceResponse:      nil,
			serviceError:         nil,
			errorMessage:         "type param value is not allowed",
		},
		{
			name:                 "Should return bad request if items params is not a number",
			typeParam:            "even",
			ipwParam:             "5",
			itemsParam:           "20abc",
			statusCode:           http.StatusBadRequest,
			expectedServiceCalls: 0,
			hasError:             true,
			serviceResponse:      nil,
			serviceError:         nil,
			errorMessage:         "items param must be a positive integer",
		},
		{
			name:                 "Should return bad request if items params is less than 0",
			typeParam:            "even",
			ipwParam:             "5",
			itemsParam:           "-1",
			statusCode:           http.StatusBadRequest,
			expectedServiceCalls: 0,
			hasError:             true,
			serviceResponse:      nil,
			serviceError:         nil,
			errorMessage:         "items param must be a positive integer",
		},
		{
			name:                 "Should return bad request if items_per_workers params is not a number",
			typeParam:            "even",
			ipwParam:             "5abc",
			itemsParam:           "20",
			statusCode:           http.StatusBadRequest,
			expectedServiceCalls: 0,
			hasError:             true,
			serviceResponse:      nil,
			serviceError:         nil,
			errorMessage:         "items_per_workers param must be a positive integer",
		},
		{
			name:                 "Should return bad request if items_per_workers params is less than 0",
			typeParam:            "even",
			ipwParam:             "-1",
			itemsParam:           "20",
			statusCode:           http.StatusBadRequest,
			expectedServiceCalls: 0,
			hasError:             true,
			serviceResponse:      nil,
			serviceError:         nil,
			errorMessage:         "items_per_workers param must be a positive integer",
		},
		{
			name:                 "Should return bad request if items_per_workers params is less than items params",
			typeParam:            "even",
			ipwParam:             "21",
			itemsParam:           "20",
			statusCode:           http.StatusBadRequest,
			expectedServiceCalls: 0,
			hasError:             true,
			serviceResponse:      nil,
			serviceError:         nil,
			errorMessage:         "items_per_workers param must be less or equal items param",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "/random-mlb-players", nil)
			q := r.URL.Query()
			q.Add("items_per_workers", tc.ipwParam)
			q.Add("items", tc.itemsParam)
			q.Add("type", tc.typeParam)
			r.URL.RawQuery = q.Encode()
			m := new(mockMLBService)
			m.On("GetMLBPlayerDesired").Return(tc.serviceResponse, tc.serviceError)
			ctr := NewMLBPlayerController(m)

			ctr.GetMLBPlayerDesired(w, r)
			res := w.Result()
			body, _ := ioutil.ReadAll(res.Body)

			if tc.hasError {
				assert.Contains(t, string(body), tc.errorMessage)
			}

			assert.Equal(t, tc.statusCode, res.StatusCode)
			assert.Equal(t, "application/json", res.Header.Get("Content-Type"))
			m.AssertNumberOfCalls(t, "GetMLBPlayerDesired", tc.expectedServiceCalls)
		})
	}
}
