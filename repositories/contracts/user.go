package repositories

import e "github.com/EloYaniel/academy-go-q42021/entities"

type UserRepository interface {
	// SaveUsers saves all users
	SaveUsers(users []e.User) error

	// GetUsers gets all Users
	GetUsers() ([]e.User, error)

	// GetUserByID get a User by its ID
	GetUserByID(id int) (*e.User, error)
}
