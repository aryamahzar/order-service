package routes

import (
	"net/http"
	"os"

	"order-service/auth"
	"order-service/handlers"

	"github.com/gorilla/mux"
)

func SetupRoutes(orderHandler *handlers.OrderHandler, authHandler *auth.AuthHandler) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/auth/token", authHandler.GenerateToken).Methods("POST")
	apiRouter := router.PathPrefix("/").Subrouter()
	if os.Getenv("ENABLE_AUTH") == "true" {
		apiRouter.Use(auth.JWTAuth) // Apply the authentication middleware
		apiRouter.Handle("/orders", auth.RequirePermission("create_order", http.HandlerFunc(orderHandler.CreateOrder))).Methods("POST")
		apiRouter.Handle("/orders/{id}", auth.RequirePermission("view_order", http.HandlerFunc(orderHandler.GetOrder))).Methods("GET")
		apiRouter.Handle("/orders/{id}", auth.RequirePermission("update_order", http.HandlerFunc(orderHandler.UpdateOrder))).Methods("PUT")
		//apiRouter.Handle("/orders/{id}", auth.RequirePermission("delete_order", http.HandlerFunc(orderHandler.DeleteOrder))).Methods("DELETE")
		apiRouter.Handle("/orders", auth.RequirePermission("list_orders", http.HandlerFunc(orderHandler.ListOrders))).Methods("GET")
	} else {
		apiRouter.HandleFunc("/orders", orderHandler.CreateOrder).Methods("POST")
		apiRouter.HandleFunc("/orders/{id}", orderHandler.GetOrder).Methods("GET")
		apiRouter.HandleFunc("/orders/{id}", orderHandler.UpdateOrder).Methods("PUT")
		//apiRouter.HandleFunc("/orders/{id}", orderHandler.DeleteOrder).Methods("DELETE")
		apiRouter.HandleFunc("/orders", orderHandler.ListOrders).Methods("GET")
	}

	return router
}
