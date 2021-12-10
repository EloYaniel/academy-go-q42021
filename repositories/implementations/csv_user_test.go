package repositories

import (
	"errors"
	"os"
	"testing"

	e "github.com/EloYaniel/academy-go-q42021/entities"
	"github.com/stretchr/testify/assert"
)

var user1 = e.User{
	ID:        1,
	Email:     "george.bluth@reqres.in",
	FirstName: "George",
	LastName:  "Bluth",
	Avatar:    "https://reqres.in/img/faces/1-image.jpg",
}
var user2 = e.User{
	ID:        2,
	Email:     "janet.weaver@reqres.in",
	FirstName: "Janet",
	LastName:  "Weaver",
	Avatar:    "https://reqres.in/img/faces/2-image.jpg",
}

func Test_NewCSVUserRepository_ShouldReturnDiffInstances(t *testing.T) {
	instance := NewCSVUserRepository("./here.csv")
	instance2 := NewCSVUserRepository("./there.csv")

	assert.NotNil(t, instance)
	assert.NotNil(t, instance2)
	assert.NotSame(t, instance, instance2)
}

func Test_SaveUsers_ShouldSaveUsers(t *testing.T) {
	users := []e.User{user1, user2}
	testCases := []struct {
		name          string
		filePath      string
		users         []e.User
		createdFile   bool
		expectedError error
	}{
		{
			name:          "Should return error when can't create file",
			filePath:      "",
			users:         users,
			createdFile:   false,
			expectedError: errors.New("error opening o creating the file"),
		},
		{
			name:          "Should save users",
			filePath:      "../../data/test/saved-users-test.csv",
			users:         users,
			createdFile:   true,
			expectedError: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			repo := NewCSVUserRepository(tc.filePath)

			err := repo.SaveUsers(tc.users)

			assert.Equal(t, tc.expectedError, err)

			if tc.createdFile {
				file, err := os.Open(repo.filePath)
				assert.Nil(t, err)
				defer file.Close()
				st, err := file.Stat()
				assert.Nil(t, err)
				assert.Greater(t, st.Size(), int64(0))
				os.Remove(repo.filePath)
			}
		})
	}
}

func Test_GetUsers_Suite(t *testing.T) {
	users := []e.User{
		user1,
		user2,
	}
	testCases := []struct {
		name             string
		filePath         string
		expectedError    error
		expectedResponse []e.User
	}{
		{
			name:             "Should return the users",
			filePath:         "../../data/test/users-test.csv",
			expectedResponse: users,
			expectedError:    nil,
		},
		{
			name:             "Should return error when open file",
			filePath:         "",
			expectedResponse: nil,
			expectedError:    errors.New("error opening the file"),
		},
		{
			name:             "Should return error when casting ID",
			filePath:         "../../data/test/users-with-wrong-id-test.csv",
			expectedResponse: nil,
			expectedError:    errors.New("error casting ID"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			repo := NewCSVUserRepository(tc.filePath)

			users, err := repo.GetUsers()

			assert.Equal(t, tc.expectedResponse, users)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func Test_GetUserByID_Suite(t *testing.T) {
	users := []e.User{
		user1,
		user2,
	}
	testCases := []struct {
		name             string
		filePath         string
		userID           int
		users            []e.User
		expectedError    error
		expectedResponse *e.User
	}{
		{
			name:             "Should return the user",
			filePath:         "../../data/test/users-test.csv",
			userID:           1,
			users:            users,
			expectedResponse: &user1,
			expectedError:    nil,
		},
		{
			name:             "Should return no user and no error",
			filePath:         "../../data/test/users-test.csv",
			users:            users,
			userID:           3,
			expectedResponse: nil,
			expectedError:    nil,
		},
		{
			name:             "Should return no user and error",
			filePath:         "",
			users:            users,
			userID:           1,
			expectedResponse: nil,
			expectedError:    errors.New("error getting user"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repo := NewCSVUserRepository(tc.filePath)

			user, err := repo.GetUserByID(tc.userID)

			assert.Equal(t, tc.expectedResponse, user)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}
