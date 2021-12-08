package repositories

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"

	e "github.com/EloYaniel/academy-go-q42021/entities"
)

type CSVUserRepository struct {
	filePath string
}

func NewCSVUserRepository(filePath string) *CSVUserRepository {
	return &CSVUserRepository{filePath: filePath}
}

func (repo *CSVUserRepository) SaveUsers(users []e.User) error {
	csvFile, err := os.Create(repo.filePath)

	if err != nil {
		return errors.New("error opening o creating the file")
	}
	defer csvFile.Close()
	csvwriter := csv.NewWriter(csvFile)
	defer csvwriter.Flush()

	fileinfo, _ := csvFile.Stat()

	if fileinfo.Size() == 0 {
		err = csvwriter.Write([]string{"Id", "Email", "FirstName", "LastName", "Avatar"})

		if err != nil {
			return errors.New((fmt.Sprint("error writing user to file:", err.Error())))
		}
	}

	for _, user := range users {
		row := []string{strconv.Itoa(user.ID), user.Email, user.FirstName, user.LastName, user.Avatar}
		err = csvwriter.Write(row)

		if err != nil {
			return errors.New((fmt.Sprint("error writing user to file:", err.Error(), "User", row)))
		}
	}

	return nil
}

func (repo *CSVUserRepository) GetUsers() ([]e.User, error) {
	f, err := os.Open(repo.filePath)

	if err != nil {
		return nil, errors.New("error opening the file")
	}
	defer f.Close()
	data, err := csv.NewReader(f).ReadAll()

	if err != nil {
		return nil, errors.New("error reading the file")
	}
	var users []e.User
	for i, line := range data {
		if i != 0 {
			id, err := strconv.Atoi(line[0])

			if err != nil {
				return nil, errors.New("error casting ID")
			}
			users = append(users, e.User{
				ID:        id,
				Email:     line[1],
				FirstName: line[2],
				LastName:  line[3],
				Avatar:    line[4],
			})
		}

	}

	return users, nil
}

func (repo *CSVUserRepository) GetUserByID(id int) (*e.User, error) {
	users, err := repo.GetUsers()

	if err != nil {
		return nil, errors.New("error getting user")
	}

	for _, u := range users {
		if u.ID == id {
			return &u, nil
		}
	}

	return nil, nil
}
