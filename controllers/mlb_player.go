package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	e "github.com/EloYaniel/academy-go-q42021/entities"
	"github.com/gorilla/mux"
)

var allowedTypeFilters = map[string]bool{"even": true, "odd": true}

type mlbPlayerService interface {
	GetMLBPlayers() ([]e.MLBPlayer, error)
	GetMLBPlayerByID(id int) (*e.MLBPlayer, error)
	GetMLBPlayerDesired(filterType string, totalItems int, itemsPerWorker int) ([]e.MLBPlayer, error)
}

type errorMessage struct {
	Message string `json:"message"`
}

// MLBPlayerController struct handles api controller.
type MLBPlayerController struct {
	service mlbPlayerService
}

// MLBPlayerController function creates an instance of NewMLBPlayerController.
func NewMLBPlayerController(service mlbPlayerService) *MLBPlayerController {
	return &MLBPlayerController{service: service}
}

// GetMLBPlayers handles list of MLB Players
func (ctr *MLBPlayerController) GetMLBPlayers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	players, err := ctr.service.GetMLBPlayers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errorMessage{
			Message: "Internal server error",
		})

		return
	}
	json.NewEncoder(w).Encode(players)
}

// GetMLBPlayers handles MLB Players by ID.
func (ctr *MLBPlayerController) GetMLBPlayerByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorMessage{
			Message: "Player ID provided must be of type integer",
		})

		return
	}
	player, err := ctr.service.GetMLBPlayerByID(id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errorMessage{
			Message: "Internal server error",
		})

		return
	}

	if player == nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(errorMessage{
			Message: "Player not found",
		})

		return
	}

	json.NewEncoder(w).Encode(player)
}

// GetMLBPlayerDesired handles list of MLB Players by filters.
func (ctr *MLBPlayerController) GetMLBPlayerDesired(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	filterType := r.FormValue("type")

	if v, ok := allowedTypeFilters[filterType]; !ok || !v {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorMessage{
			Message: "type param value is not allowed",
		})

		return
	}
	itemsperworkers, err := strconv.Atoi(r.FormValue("items_per_workers"))

	if err != nil || itemsperworkers <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorMessage{
			Message: "items_per_workers param must be a positive integer",
		})

		return
	}

	items, err := strconv.Atoi(r.FormValue("items"))

	if err != nil || items <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorMessage{
			Message: "items param must be a positive integer",
		})

		return
	}

	if itemsperworkers > items {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorMessage{
			Message: "items_per_workers param must be less or equal items param",
		})

		return
	}
	players, err := ctr.service.GetMLBPlayerDesired(filterType, items, itemsperworkers)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errorMessage{
			Message: "Internal server error",
		})

		return
	}

	json.NewEncoder(w).Encode(struct {
		Count   int           `json:"total"`
		Players []e.MLBPlayer `json:"players"`
	}{
		len(players),
		players,
	})
}
