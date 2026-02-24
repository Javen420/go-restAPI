package handlers

import (
	"encoding/json"
	"net/http"
	"restAPI/service"
)

type menuHandler struct {
	service *service.MenuService
}

func NewMenuHandler(service *service.MenuService) *menuHandler {
	return &menuHandler{service: service}
}

func (h *menuHandler) GetAllMenu(w http.ResponseWriter, r *http.Request) {
	menu, err := h.service.GetFullMenu(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(menu)
}
