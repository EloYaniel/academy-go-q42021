package app

import (
	ctr "github.com/EloYaniel/academy-go-q42021/controllers"
	repo "github.com/EloYaniel/academy-go-q42021/repositories/implementations"
	srv "github.com/EloYaniel/academy-go-q42021/services"
	"github.com/gorilla/mux"
)

func InitApp() *mux.Router {
	csvmlbrepository := repo.NewCSVMLBPlayerRepository("data/mlb_players.csv")

	service := srv.NewMLBPlayerService(csvmlbrepository)

	healthcontroller := ctr.NewHealthController()
	mlbplayercontroller := ctr.NewMLBPlayerController(service)

	r := mux.NewRouter()
	r.HandleFunc("/health", healthcontroller.CheckHealth)
	r.HandleFunc("/mlb-players", mlbplayercontroller.GetMLBPlayers)
	r.HandleFunc("/mlb-players/{id}", mlbplayercontroller.GetMLBPlayerByID)

	return r
}
