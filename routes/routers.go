package routes

import (
	"net/http"
	"order-service/auth"
	"order-service/handlers"

	"github.com/gorilla/mux"
)

func SetupRoutes(orderHandler *handlers.OrderHandler, authHandler *auth.AuthHandler) *mux.Router {
	r := mux.NewRouter()

	// Auth routes
	r.HandleFunc("/auth/token", authHandler.GenerateToken).Methods("POST")
	apiRouter := r.PathPrefix("/").Subrouter()
	apiRouter.Use(auth.JWTAuth) // Apply the authentication middleware

	// Order routes (apply middleware to protected routes)
	apiRouter.Handle("/orders", auth.RequirePermission("CreateOrder", http.HandlerFunc(orderHandler.CreateOrder))).Methods("POST")
	apiRouter.Handle("/orders", auth.RequirePermission("ListOrders", http.HandlerFunc(orderHandler.ListOrders))).Methods("GET")
	apiRouter.Handle("/orders/{orderId}", auth.RequirePermission("GetOrder", http.HandlerFunc(orderHandler.GetOrder))).Methods("GET")
	apiRouter.Handle("/orders/{orderId}", auth.RequirePermission("UpdateOrder", http.HandlerFunc(orderHandler.UpdateOrder))).Methods("PUT")
	apiRouter.Handle("/orders/{orderId}/cancel", auth.RequirePermission("CancelOrder", http.HandlerFunc(orderHandler.CancelOrder))).Methods("POST")

	return r
}
