package handlers

import (
	"encoding/json"
	"net/http"
	"restAPI/models"

	"github.com/gorilla/mux"
)

type orderHandler struct {
	// db *sql.DB  // add this later when you have a database
}

func NewOrderHandler() *orderHandler {
	return &orderHandler{}
}

func (h *orderHandler) GetOrder(w http.ResponseWriter, r *http.Request) {
	// 1. Get the ID from the URL
	vars := mux.Vars(r)
	id := vars["id"]

	// 2. Do something with it (fetch from DB later)

	// 3. Write response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"order_id": id})
}

func (h *orderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	// 1. Decode the request body
	var items []models.Item
	json.NewDecoder(r.Body).Decode(&items)

	// 2. Create the order
	order, err := models.NewOrder(items)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// 3. Write response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(order)
}
