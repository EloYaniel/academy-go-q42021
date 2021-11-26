package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	e "github.com/EloYaniel/academy-go-q42021/entities"
	"github.com/gorilla/mux"
)

type mlbPlayerService interface {
	GetMLBPlayers() ([]e.MLBPlayer, error)
	GetMLBPlayerByID(id int) (*e.MLBPlayer, error)
}

type errorMessage struct {
	Message string `json:"message"`
}

type MLBPlayerController struct {
	service mlbPlayerService
}

func NewMLBPlayerController(service mlbPlayerService) *MLBPlayerController {
	return &MLBPlayerController{service: service}
}

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
