package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/yemiwebby/sta-monolithic-app/handlers"
)

func main() {
	r := mux.NewRouter()

	// Routes for each service
	r.HandleFunc("/", handlers.HomePage).Methods("GET")
	r.HandleFunc("/validate-session", logRequest(handlers.ValidateSession)).Methods("POST")
	r.HandleFunc("/products", logRequest(handlers.GetProducts)).Methods("GET")
	r.HandleFunc("/cart", logRequest(handlers.AddToCart)).Methods("POST")
	r.HandleFunc("/checkout", logRequest(handlers.Checkout)).Methods("POST")
	r.HandleFunc("/send", logRequest(handlers.SendEmail)).Methods("POST")


	// Simulated intentional bug for demonstration
	r.HandleFunc("/buggy-endpoint", logRequest(handlers.BuggyEndpoint)).Methods("GET")


	// Start the server
	log.Println("Monolithic App running on port 8080...")
	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}

func logRequest(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("Started %s %s", r.Method, r.URL.Path)
		handler(w, r)
		duration := time.Since(start)
		log.Printf("Completed %s %s in %v", r.Method, r.URL.Path, duration)
	}
}
