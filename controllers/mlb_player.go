package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/EloYaniel/academy-go-q42021/entities"
	r "github.com/EloYaniel/academy-go-q42021/repositories/implementations"
	"github.com/EloYaniel/academy-go-q42021/services"
	"github.com/gorilla/mux"
)

type errorMessage struct {
	Message string `json:"message"`
}

var service services.MLBPlayerService = *services.NewMLBPlayerService(&r.CSVMLBPlayerRepository{})

func GetMLBPlayers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	players, err := service.GetMLBPlayers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errorMessage{
			Message: "Internal server error",
		})

		return
	}
	json.NewEncoder(w).Encode(players)
}

func GetMLBPlayerByID(w http.ResponseWriter, r *http.Request) {
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
	player, err := service.GetMLBPlayerByID(id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errorMessage{
			Message: "Internal server error",
		})

		return
	}

	if (*player == entities.MLBPlayer{}) {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(errorMessage{
			Message: "Player not found",
		})

		return
	}

	json.NewEncoder(w).Encode(player)
}
