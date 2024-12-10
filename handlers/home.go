package handlers

import (
	"net/http"
)

// HomePage serves as the entry point for the application.
func HomePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>Monolithic App</title>
		</head>
		<body>
			<h1>Welcome to the Monolithic App</h1>
			<p>All services are running within this single application.</p>
			<ul>
				<li><a href="/validate-session">Validate Session</a> (POST)</li>
				<li><a href="/products">Fetch Products</a> (GET)</li>
				<li><a href="/cart">Add to Cart</a> (POST)</li>
				<li><a href="/checkout">Checkout</a> (POST)</li>
				<li><a href="/send">Send Email</a> (POST)</li>
				<li><a href="/buggy-endpoint">Trigger Bug</a> (GET)</li>
			</ul>
		</body>
		</html>
	`))
}