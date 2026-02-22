package handlers

import (
	"encoding/json"
	"net/http"
	"restAPI/models"

	"github.com/gorilla/mux"
)

type menuHandler struct {
}

func NewMenuHandler() *menuHandler {
	return &menuHandler{}
}

func (h *menuHandler) GetAllMenu(w http.ResponseWriter, r *http.Request) {
	//db
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode([]models.MenuItem{})
}

func (h *menuHandler) GetMenu(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	//db

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"MenuItem": id})
}
