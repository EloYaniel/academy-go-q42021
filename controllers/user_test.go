package controllers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	e "github.com/EloYaniel/academy-go-q42021/entities"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockUserService struct {
	mock.Mock
}

func (m *mockUserService) GetUsers() ([]e.User, error) {
	args := m.Called()

	return args.Get(0).([]e.User), args.Error(1)
}

func (m *mockUserService) GetUserByID(id int) (*e.User, error) {
	args := m.Called()

	return args.Get(0).(*e.User), args.Error(1)
}

func Test_UserController_GetUsers_Suite(t *testing.T) {
	testCases := []struct {
		name                 string
		statusCode           int
		expectedServiceCalls int
		hasError             bool
		serviceError         error
		serviceResponse      []e.User
		errorMessage         string
	}{
		{
			name:                 "Should return users",
			statusCode:           http.StatusOK,
			expectedServiceCalls: 1,
			hasError:             false,
			serviceResponse: []e.User{{
				ID:        1,
				Email:     "e@gmail.com",
				FirstName: "First",
				LastName:  "Last",
				Avatar:    "FL",
			}},
			serviceError: nil,
		},
		{
			name:                 "Should return internal server on service error",
			statusCode:           http.StatusInternalServerError,
			expectedServiceCalls: 1,
			hasError:             true,
			serviceResponse:      nil,
			serviceError:         errors.New("unknown error"),
			errorMessage:         "Internal server error",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "/users", nil)
			m := new(mockUserService)
			m.On("GetUsers").Return(tc.serviceResponse, tc.serviceError)
			ctr := NewUserController(m)

			ctr.GetUsers(w, r)

			if tc.hasError {
				assert.Contains(t, w.Body.String(), tc.errorMessage)
			}

			assert.Equal(t, w.Code, tc.statusCode)
			assert.Equal(t, "application/json", w.Result().Header.Get("Content-Type"))
			m.AssertNumberOfCalls(t, "GetUsers", tc.expectedServiceCalls)
		})
	}
}

func Test_UserController_GetUserByID_Suite(t *testing.T) {
	testCases := []struct {
		name                 string
		idParam              string
		statusCode           int
		expectedServiceCalls int
		hasError             bool
		serviceError         error
		serviceResponse      *e.User
		errorMessage         string
	}{
		{
			name:                 "Should return user",
			idParam:              "1",
			statusCode:           http.StatusOK,
			expectedServiceCalls: 1,
			hasError:             false,
			serviceResponse: &e.User{
				ID:        1,
				Email:     "e@gmail.com",
				FirstName: "First",
				LastName:  "Last",
				Avatar:    "FL",
			},
			serviceError: nil,
		},
		{
			name:                 "Should return internal server on service error",
			idParam:              "1",
			statusCode:           http.StatusInternalServerError,
			expectedServiceCalls: 1,
			hasError:             true,
			serviceResponse:      nil,
			serviceError:         errors.New("unknown error"),
			errorMessage:         "Internal server error",
		},
		{
			name:                 "Should return not found if no user",
			idParam:              "10",
			statusCode:           http.StatusNotFound,
			expectedServiceCalls: 1,
			hasError:             false,
			serviceResponse:      nil,
			serviceError:         nil,
			errorMessage:         "User not found",
		},
		{
			name:                 "Should bad request if no user id provided",
			idParam:              "1a2b",
			statusCode:           http.StatusBadRequest,
			expectedServiceCalls: 0,
			hasError:             true,
			serviceResponse:      nil,
			serviceError:         nil,
			errorMessage:         "User ID provided must be of type integer",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "/users/{id}", nil)
			r = mux.SetURLVars(r, map[string]string{"id": tc.idParam})
			m := new(mockUserService)
			m.On("GetUserByID").Return(tc.serviceResponse, tc.serviceError)
			ctr := NewUserController(m)

			ctr.GetUserByID(w, r)

			if tc.hasError {
				assert.Contains(t, w.Body.String(), tc.errorMessage)
			}

			assert.Equal(t, w.Code, tc.statusCode)
			assert.Equal(t, "application/json", w.Result().Header.Get("Content-Type"))
			m.AssertNumberOfCalls(t, "GetUserByID", tc.expectedServiceCalls)
		})
	}
}
