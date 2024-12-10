package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

func ValidateSession(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Token string `json:"token"`
	}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Printf("Error decoding request body: %v", err)
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	if req.Token == "Bearer dummy-token" {
		log.Println("Session validated successfully")
		json.NewEncoder(w).Encode(map[string]bool{"valid": true})
		return
	}
	log.Println("Session validation failed")
	http.Error(w, "Invalid token", http.StatusUnauthorized)
}
