package services

import (
	"github.com/EloYaniel/academy-go-q42021/apiclient"
	e "github.com/EloYaniel/academy-go-q42021/entities"
	repo "github.com/EloYaniel/academy-go-q42021/repositories/contracts"
)

type UserService struct {
	repo      repo.UserRepository
	apiClient apiclient.ApiClient
}

func NewUserService(repo repo.UserRepository, client apiclient.ApiClient) *UserService {
	return &UserService{repo: repo, apiClient: client}
}

func (s *UserService) GetUsers() ([]e.User, error) {
	users, _ := s.repo.GetUsers()

	if len(users) == 0 {
		resp := struct {
			Data []e.User
		}{}
		err := s.apiClient.Get("https://reqres.in/api/users", nil, &resp)

		if err != nil {
			return nil, err
		}

		users = resp.Data
		if err != nil {
			return nil, err
		}
		s.repo.SaveUsers(users)
	}

	return users, nil
}

func (s *UserService) GetUserByID(id int) (*e.User, error) {
	return s.repo.GetUserByID(id)
}
