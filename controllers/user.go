package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	e "github.com/EloYaniel/academy-go-q42021/entities"
	"github.com/gorilla/mux"
)

type userService interface {
	GetUsers() ([]e.User, error)
	GetUserByID(id int) (*e.User, error)
}

type UserController struct {
	service userService
}

func NewUserController(service userService) *UserController {
	return &UserController{service: service}
}

func (ctr *UserController) GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	users, err := ctr.service.GetUsers()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errorMessage{
			Message: "Internal server error",
		})

		return
	}
	json.NewEncoder(w).Encode(users)
}

func (ctr *UserController) GetUserByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorMessage{
			Message: "User ID provided must be of type integer",
		})

		return
	}
	user, err := ctr.service.GetUserByID(id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errorMessage{
			Message: "Internal server error",
		})

		return
	}

	if user == nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(errorMessage{
			Message: "User not found",
		})

		return
	}

	json.NewEncoder(w).Encode(user)
}
