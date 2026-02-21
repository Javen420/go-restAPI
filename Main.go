package main

import (
	"encoding/json"
	"net/http"
	"restAPI/handlers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	orderHandler := handlers.NewOrderHandler()

	r.HandleFunc("/orders", orderHandler.CreateOrder).Methods("POST")
	r.HandleFunc("/orders/{id}", orderHandler.GetOrder).Methods("GET")

	http.Handle("/", r)

}
