package utils

import (
	"encoding/json"
	"net/http"
)

// RespondJSON mengirimkan respon JSON dengan status code tertentu
func RespondJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

// RespondError mengirimkan pesan error JSON
func RespondError(w http.ResponseWriter, status int, message string) {
	RespondJSON(w, status, map[string]string{"error": message})
}
