package handlers

import (
	"net/http"
)

func AddToCart(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Item added to cart"))
}

func Checkout(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Order placed successfully"))
}
