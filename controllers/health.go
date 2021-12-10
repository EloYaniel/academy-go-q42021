package controllers

import (
	"encoding/json"
	"net/http"
)

// HealthController struct handles api controller.
type HealthController struct{}

// NewHealthController function creates an instance of HealthController.
func NewHealthController() *HealthController {
	return &HealthController{}
}

// CheckHealth checks API health.
func (HealthController) CheckHealth(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("API is healthy")
}
