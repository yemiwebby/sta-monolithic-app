package handlers

import (
	"encoding/json"
	"net/http"
)

var products = []map[string]interface{}{
	{"id": "1", "name": "Laptop", "price": 1200.0, "stock": 5},
	{"id": "2", "name": "Phone", "price": 800.0, "stock": 10},
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(products)
}
