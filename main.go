package main

import (
	"encoding/json"
	"net/http"
)

// Define Item struct
type Item struct {
	ID   string `json:"ID"`
	Name string `json:"Name"`
}

// In-memory storage for items
var items = []Item{}

const Dport = ":8012"

// created by Samhitha Dubbaka - 500225971
// Items handler to handle get and post requests for items
func itemsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getItems(w, r)
	case http.MethodPost:
		addItems(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// Function to get all the items
// created by Abhisheik Yadla
func getItems(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(items)
}
