package services

import (
	"log"

	"github.com/EloYaniel/academy-go-q42021/apiclient"
	e "github.com/EloYaniel/academy-go-q42021/entities"
	repo "github.com/EloYaniel/academy-go-q42021/repositories/contracts"
)

// UserService struct handles Users business logic.
type UserService struct {
	repo      repo.UserRepository
	apiClient apiclient.ApiClient
	userURL   string
}

// NewUserService function return an instance of UserService
func NewUserService(repo repo.UserRepository, client apiclient.ApiClient, userURL string) *UserService {
	return &UserService{repo: repo, apiClient: client, userURL: userURL}
}

// GetUsers gets all Users
func (s *UserService) GetUsers() ([]e.User, error) {
	users, err := s.repo.GetUsers()

	if err != nil {
		log.Println(err)
	}

	if len(users) > 0 {
		return users, nil
	}

	resp := struct {
		Data []e.User
	}{}
	err = s.apiClient.Get(s.userURL, nil, &resp)

	if err != nil {
		log.Println(err)
		return nil, err
	}
	users = resp.Data
	err = s.repo.SaveUsers(users)

	if err != nil {
		log.Println(err)
	}

	return users, nil
}

// GetUserByID get a User by its ID
func (s *UserService) GetUserByID(id int) (*e.User, error) {
	return s.repo.GetUserByID(id)
}
