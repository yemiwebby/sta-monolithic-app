package handlers

import "net/http"

func BuggyEndpoint(w http.ResponseWriter, r *http.Request) {
	// Simulate a crash
	panic("Intentional bug triggered")
}
