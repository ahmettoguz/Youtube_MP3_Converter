package health

import (
	"encoding/json"
	"net/http"
	"time"
)

// HealthResponse defines the JSON structure
type HealthResponse struct {
	Status string `json:"status"`
	Time   string `json:"time"`
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	resp := HealthResponse{
		Status: "UP",
		Time:   time.Now().Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(resp)
}
