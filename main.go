package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"order-service/auth"
	"order-service/handlers"
	"order-service/repository"
	"order-service/routes"
	"order-service/service"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	// MongoDB connection setup
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// MongoDB connection setup
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Fatal("MONGO_URI environment variable not set")
	}

	log.Printf("mongoURI: %+v\n", mongoURI)

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Fatalf("Failed to disconnect from MongoDB: %v", err)
		}
	}()

	// Check the connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Could not connect to MongoDB:", err)
	} else {
		log.Println("Connected to MongoDB!")
	}

	// Get a handle to your database
	db := client.Database("e-commerce")

	// ... rest of your main function ...

	orderRepository := repository.NewOrderRepository(db)
	orderService := service.NewOrderService(*orderRepository)
	apiGatewayURL := "https://api-gateway:8002"
	orderHandler := handlers.NewOrderHandler(*orderService, apiGatewayURL)
	authHandler := new(auth.AuthHandler)

	r := routes.SetupRoutes(orderHandler, authHandler)

	// Apply authentication middleware if ENABLE_AUTH is true
	if os.Getenv("ENABLE_AUTH") == "true" {
		r.Use(auth.JWTAuth)
	}

	log.Println("Server listening on port 8003")
	http.ListenAndServe(":8003", r)
}
