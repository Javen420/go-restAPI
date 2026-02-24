package main

import (
	"fmt"
	"log"
	"net/http"
	"restAPI/db"
	"restAPI/db/repository"
	"restAPI/handlers"
	"restAPI/service"

	"github.com/gorilla/mux"
)

func main() {

	pool := db.NewPool()
	r := mux.NewRouter()

	menuRepo := repository.NewMenuRepository(pool)
	menuService := service.NewMenuService(menuRepo)
	menuHandler := handlers.NewMenuHandler(menuService)

	orderRepo := repository.NewOrderRepository(pool)
	orderService := service.NewOrderService(orderRepo)
	orderHandler := handlers.NewOrderHandler(orderService)

	//order requests
	r.HandleFunc("/orders/", orderHandler.GetAllOrders).Methods("GET") //kitchen side maybe

	r.HandleFunc("/orders/{id}", orderHandler.GetOrder).Methods("GET")
	r.HandleFunc("/orders", orderHandler.CreateOrder).Methods("POST")
	r.HandleFunc("/orders/{id}", orderHandler.DeleteOrder).Methods("DELETE")
	r.HandleFunc("/orders/{id}", orderHandler.UpdateTotalPrice).Methods("POST")
	r.HandleFunc("/orders/{id}", orderHandler.ChangeOrderStatus).Methods("PATCH")

	//menu requests
	r.HandleFunc("/menu/", menuHandler.GetAllMenu).Methods("GET")
	fmt.Println("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
