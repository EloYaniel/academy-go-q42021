package apiclient

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type responseBody struct {
	Name     string
	LastName string
}

func Test_GetHttpApiClientInstance_ShouldReturnInstance(t *testing.T) {
	instance := GetHttpApiClientInstance()

	assert.NotNil(t, instance)
}

func Test_GetHttpApiClientInstance_ShouldReturnSingleton(t *testing.T) {
	instance := GetHttpApiClientInstance()
	instance2 := GetHttpApiClientInstance()

	assert.Same(t, instance, instance2)
}

func Test_Get_Suite(t *testing.T) {
	testCases := []struct {
		name             string
		apiResponse      []byte
		expectedResponse interface{}
		statusCode       int
		expectedError    error
		hasError         bool
	}{
		{
			name:             "Should return error when parsing body",
			apiResponse:      []byte("malformed JSON"),
			expectedResponse: nil,
			statusCode:       200,
			expectedError:    errors.New("error parsing body response:"),
			hasError:         true,
		},
		{
			name:        "Should return parsed body",
			apiResponse: []byte("{\"Name\": \"Juan\", \"LastName\": \"Alonso\"}"),
			expectedResponse: responseBody{
				Name:     "Juan",
				LastName: "Alonso",
			},
			statusCode:    200,
			expectedError: nil,
			hasError:      false,
		},
		{
			name:             "Should return error if GET client fails",
			apiResponse:      nil,
			expectedResponse: nil,
			statusCode:       301,
			expectedError:    errors.New("301 response missing Location header"),
			hasError:         true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				res.WriteHeader(tc.statusCode)
				res.Write(tc.apiResponse)
			}))

			defer testServer.Close()

			client := GetHttpApiClientInstance()

			resp := responseBody{}

			err := client.Get(testServer.URL, nil, &resp)

			if tc.hasError {
				assert.Contains(t, err.Error(), tc.expectedError.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tc.expectedResponse, resp)
			}
		})
	}
}
