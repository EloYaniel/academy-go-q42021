package services

import (
	"errors"
	"testing"

	e "github.com/EloYaniel/academy-go-q42021/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockApiClient struct {
	mock.Mock
}

func (m *mockApiClient) Get(url string, params map[string]interface{}, response interface{}) error {
	args := m.Called()

	return args.Error(0)
}

type mockUserRepository struct {
	mock.Mock
}

func (m *mockUserRepository) SaveUsers(users []e.User) error {
	args := m.Called()

	return args.Error(0)
}

func (m *mockUserRepository) GetUserByID(id int) (*e.User, error) {
	args := m.Called()

	return args.Get(0).(*e.User), args.Error(1)
}

func (m *mockUserRepository) GetUsers() ([]e.User, error) {
	args := m.Called()

	return args.Get(0).([]e.User), args.Error(1)
}

func Test_NewUserService_ShouldReturnInstance(t *testing.T) {
	instance := NewUserService(&mockUserRepository{}, &mockApiClient{}, "http://user.com")
	instance2 := NewUserService(&mockUserRepository{}, &mockApiClient{}, "http://user.com")

	assert.NotNil(t, instance)
	assert.NotSame(t, instance, instance2)
}

func Test_GetUsers_Suite(t *testing.T) {
	testCases := []struct {
		name                       string
		response                   []e.User
		clientErr                  error
		getUsersRepoErr            error
		saveUsersRepoErr           error
		expectedGetUsersRepoCalls  int
		expectedSaveUsersRepoCalls int
		expectedClientCalls        int
	}{
		{
			name:                       "Should return error when client has error",
			clientErr:                  errors.New("Error getting users"),
			expectedSaveUsersRepoCalls: 0,
			expectedGetUsersRepoCalls:  1,
			expectedClientCalls:        1,
		},
		{
			name:                       "GetUser has error",
			clientErr:                  nil,
			getUsersRepoErr:            errors.New("Error saving users"),
			expectedSaveUsersRepoCalls: 1,
			expectedGetUsersRepoCalls:  1,
			expectedClientCalls:        1,
		},
		{
			name: "Should return users from GetUser",
			response: []e.User{
				{
					ID:        1,
					Email:     "e.gmail.com",
					FirstName: "el",
					LastName:  "pe",
					Avatar:    "EP",
				},
				{
					ID:        2,
					Email:     "f.gmail.com",
					FirstName: "fe",
					LastName:  "lm",
					Avatar:    "FL",
				},
			},
			expectedSaveUsersRepoCalls: 0,
			expectedGetUsersRepoCalls:  1,
			expectedClientCalls:        0,
		},
		{
			name:                       "Should return users from Client",
			getUsersRepoErr:            errors.New("unknown error"),
			expectedSaveUsersRepoCalls: 1,
			expectedGetUsersRepoCalls:  1,
			expectedClientCalls:        1,
		},
		{
			name:                       "Save user error",
			getUsersRepoErr:            errors.New("unknown error"),
			saveUsersRepoErr:           errors.New("Error saving user"),
			expectedSaveUsersRepoCalls: 1,
			expectedGetUsersRepoCalls:  1,
			expectedClientCalls:        1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repoMock := new(mockUserRepository)
			repoMock.On("GetUsers").Return(tc.response, tc.getUsersRepoErr)
			repoMock.On("SaveUsers").Return(tc.saveUsersRepoErr)
			clientMock := new(mockApiClient)
			clientMock.On("Get").Return(tc.clientErr)
			service := NewUserService(repoMock, clientMock, "http://user.com")

			resp, err := service.GetUsers()

			if err != nil {
				assert.Equal(t, tc.clientErr, err)
				assert.Nil(t, resp)

			} else {
				assert.Nil(t, err)
				assert.Equal(t, tc.response, resp)
			}
			repoMock.AssertNumberOfCalls(t, "SaveUsers", tc.expectedSaveUsersRepoCalls)
			repoMock.AssertNumberOfCalls(t, "GetUsers", tc.expectedGetUsersRepoCalls)
			clientMock.AssertNumberOfCalls(t, "Get", tc.expectedClientCalls)
		})
	}
}

func Test_GetUserByID_Suite(t *testing.T) {
	testCases := []struct {
		name     string
		response *e.User
		err      error
	}{
		{
			name:     "Should return error when repo has error",
			response: nil,
			err:      errors.New("Error getting user"),
		},
		{
			name: "Should return user by id",
			response: &e.User{
				ID:        1,
				Email:     "e.gmail.com",
				FirstName: "el",
				LastName:  "pe",
				Avatar:    "EP",
			},
			err: nil,
		},
		{
			name:     "Should return no error nor user if not found",
			response: nil,
			err:      nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repoMock := new(mockUserRepository)
			clientMock := new(mockApiClient)
			repoMock.On("GetUserByID").Return(tc.response, tc.err)
			service := NewUserService(repoMock, clientMock, "http://user.com")

			resp, err := service.GetUserByID(1)

			if err != nil {
				assert.Equal(t, tc.err, err)
			}
			if resp != nil {
				assert.Equal(t, tc.response, resp)
			}
			repoMock.AssertNumberOfCalls(t, "GetUserByID", 1)
		})
	}
}
