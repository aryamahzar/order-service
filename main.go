package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"order-service/handlers"
	"order-service/repository"
	"order-service/routes"
	"order-service/service"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// MongoDB connection setup
	mongoURI := os.Getenv("MONGODB_URI")
	if mongoURI == "" {
		mongoURI = "mongodb://localhost:27017" // fallback in case MONGO_URI is not set
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
	orderHandler := handlers.NewOrderHandler(*orderService)

	r := routes.SetupRoutes(orderHandler)

	log.Println("Server listening on port 8080")
	http.ListenAndServe(":8080", r)
}
