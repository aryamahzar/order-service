package main

import (
	"order-service/handlers"

	"github.com/gorilla/mux"
)

func setupRoutes(orderHandler *handlers.OrderHandler) *mux.Router {
	r := mux.NewRouter()

	// Order routes
	r.HandleFunc("/orders", orderHandler.CreateOrder).Methods("POST")
	r.HandleFunc("/orders", orderHandler.ListOrders).Methods("GET")
	r.HandleFunc("/orders/{orderId}", orderHandler.GetOrder).Methods("GET")
	r.HandleFunc("/orders/{orderId}", orderHandler.UpdateOrder).Methods("PUT")
	r.HandleFunc("/orders/{orderId}/cancel", orderHandler.CancelOrder).Methods("POST")

	return r
}
