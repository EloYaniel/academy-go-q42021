package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Health_Suite(t *testing.T) {
	testCases := []struct {
		name       string
		statusCode int
	}{
		{
			name:       "Should return healthy api",
			statusCode: http.StatusOK,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "/health", nil)

			NewHealthController().CheckHealth(w, r)

			assert.Equal(t, tc.statusCode, w.Code)
		})
	}
}
