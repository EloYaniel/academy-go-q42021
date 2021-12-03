package repositories

import e "github.com/EloYaniel/academy-go-q42021/entities"

type UserRepository interface {
	SaveUsers(users []e.User) error
	GetUsers() ([]e.User, error)
	GetUserByID(id int) (*e.User, error)
}
