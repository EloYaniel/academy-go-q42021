package controllers

import (
	"encoding/json"
	"net/http"
)

type HealthController struct{}

func NewHealthController() *HealthController {
	return &HealthController{}
}

func (HealthController) CheckHealth(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("API is healthy")
}
