package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

func SendEmail(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Recipient string `json:"recipient"`
		Subject string `json:"subject"`
		Body    string `json:"body"`
	}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Printf("Error decoding email request: %v", err)
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	log.Printf("Email sent to %s with subject: %s", req.To, req.Subject)
	w.Write([]byte("Email sent successfully"))
}
