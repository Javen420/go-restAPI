package main

import (
	"net/http"
	"restAPI/handlers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	orderHandler := handlers.NewOrderHandler()
	menuHandler := handlers.NewMenuHandler()

	//order requests
	r.HandleFunc("/orders/", orderHandler.GetAllOrders).Methods("GET") //kitchen side maybe

	r.HandleFunc("/orders/{id}", orderHandler.GetOrder).Methods("GET")
	r.HandleFunc("/orders", orderHandler.CreateOrder).Methods("POST")
	r.HandleFunc("/orders/{id}", orderHandler.DeleteOrder).Methods("DELETE")
	r.HandleFunc("/orders/{id}", orderHandler.UpdateOrder).Methods("POST")
	r.HandleFunc("/orders/{id}", orderHandler.ChangeOrderStatus).Methods("PATCH")
	//menu requests
	r.HandleFunc("/menu/", menuHandler.GetAllMenu).Methods("GET")
	r.HandleFunc("/menu/{id}", menuHandler.GetMenu).Methods("GET")
	//could potentially add other request types for menu

	http.Handle("/", r)

}
