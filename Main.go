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

	//order requests
	r.HandleFunc("/orders/", orderHandler.GetAllOrders).Methods("GET") //kitchen side maybe

	r.HandleFunc("/orders/{id}", orderHandler.GetOrder).Methods("GET")
	r.HandleFunc("/orders", orderHandler.CreateOrder).Methods("POST")
	r.HandleFunc("/orders/{id}", orderHandler.DeleteOrder).Methods("DELETE")
	r.HandleFunc("/orders/{id}", orderHandler.UpdateOrder).Methods("POST")
	r.HandleFunc("orders/{id}", orderHandler.ChangeOrderStatus).Methods("PATCH")
	//menu requests
	r.HandleFunc("/menu/", menuHandler.getAllMenu).Methods("GET")
	r.HandleFunc("menu/{id}", menuHandler.getMenuItem).Methods("GET")
	//could potentially add other request types for menu

	http.Handle("/", r)

}
