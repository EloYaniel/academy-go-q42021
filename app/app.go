package app

import (
	"github.com/EloYaniel/academy-go-q42021/apiclient"
	ctr "github.com/EloYaniel/academy-go-q42021/controllers"
	repo "github.com/EloYaniel/academy-go-q42021/repositories/implementations"
	srv "github.com/EloYaniel/academy-go-q42021/services"
	"github.com/gorilla/mux"
)

func InitApp() *mux.Router {
	apiclient := apiclient.GetHttpApiClientInstance()

	csvmlbrepository := repo.NewCSVMLBPlayerRepository("data/mlb_players.csv")
	csvuserrepository := repo.NewCSVUserRepository("data/users.csv")

	mlbplayerservice := srv.NewMLBPlayerService(csvmlbrepository)
	userservice := srv.NewUserService(csvuserrepository, apiclient, "https://reqres.in/api/users")

	healthcontroller := ctr.NewHealthController()
	mlbplayercontroller := ctr.NewMLBPlayerController(mlbplayerservice)
	usercontroller := ctr.NewUserController(userservice)

	r := mux.NewRouter()
	r.HandleFunc("/health", healthcontroller.CheckHealth)
	r.HandleFunc("/mlb-players", mlbplayercontroller.GetMLBPlayers)
	r.HandleFunc("/mlb-players/{id}", mlbplayercontroller.GetMLBPlayerByID)
	r.HandleFunc("/users", usercontroller.GetUsers)
	r.HandleFunc("/users/{id}", usercontroller.GetUserByID)

	return r
}
