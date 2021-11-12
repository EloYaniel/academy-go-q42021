package app

import (
	"encoding/json"
	"net/http"

	controllers "github.com/EloYaniel/academy-go-q42021/controllers"
	"github.com/gorilla/mux"
)

func InitApp() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("Welcome to The Golang Bootcamp API")
	})

	r.HandleFunc("/mlb-players", controllers.GetMLBPlayers)
	r.HandleFunc("/mlb-players/{id}", controllers.GetMLBPlayerByID)

	return r
}
