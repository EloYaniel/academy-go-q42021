package services

import (
	"log"

	"github.com/EloYaniel/academy-go-q42021/apiclient"
	e "github.com/EloYaniel/academy-go-q42021/entities"
	repo "github.com/EloYaniel/academy-go-q42021/repositories/contracts"
)

type UserService struct {
	repo      repo.UserRepository
	apiClient apiclient.ApiClient
	userURL   string
}

func NewUserService(repo repo.UserRepository, client apiclient.ApiClient, userURL string) *UserService {
	return &UserService{repo: repo, apiClient: client, userURL: userURL}
}

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
		return nil, err
	}

	return users, nil
}

func (s *UserService) GetUserByID(id int) (*e.User, error) {
	return s.repo.GetUserByID(id)
}
