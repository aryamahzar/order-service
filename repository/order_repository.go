package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"order-service/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type OrderRepository struct {
	collection *mongo.Collection
}

func NewOrderRepository(db *mongo.Database) *OrderRepository {
	return &OrderRepository{
		collection: db.Collection("orders"),
	}
}

func (r *OrderRepository) CreateOrder(ctx context.Context, order *models.Order) error {
	// Generate a new ObjectId for the order
	order.OrderID = primitive.NewObjectID().Hex()
	order.OrderDate = time.Now()

	_, err := r.collection.InsertOne(ctx, order)
	return err
}

func (r *OrderRepository) GetOrder(ctx context.Context, orderID string) (*models.Order, error) {
	var order models.Order
	// Convert the orderID string to an ObjectId
	objID, err := primitive.ObjectIDFromHex(orderID)
	if err != nil {
		return nil, fmt.Errorf("invalid order ID: %v", err)
	}

	err = r.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&order)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("order not found")
		}
		return nil, err
	}
	return &order, nil
}

func (r *OrderRepository) UpdateOrder(ctx context.Context, order *models.Order) error {
	objID, err := primitive.ObjectIDFromHex(order.OrderID)
	if err != nil {
		return fmt.Errorf("invalid order ID: %v", err)
	}

	update := bson.M{
		"$set": bson.M{
			"status":          order.Status,
			"shippingAddress": order.ShippingAddress,
			"billingAddress":  order.BillingAddress,
			// ... other fields you want to update ...
		},
	}

	_, err = r.collection.UpdateOne(ctx, bson.M{"_id": objID}, update)
	return err
}

func (r *OrderRepository) ListOrders(ctx context.Context, userID *string, status *models.OrderStatus, page, size int64) ([]*models.Order, int64, error) {
	filter := bson.M{}
	if userID != nil {
		filter["userId"] = *userID
	}
	if status != nil {
		filter["status"] = *status
	}

	// Pagination options
	opts := options.Find().SetSkip((page - 1) * size).SetLimit(size)

	cursor, err := r.collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)

	var orders []*models.Order
	if err = cursor.All(ctx, &orders); err != nil {
		return nil, 0, err
	}

	// Get total count for pagination
	totalCount, err := r.collection.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	return orders, totalCount, nil
}

// ... (Add DeleteOrder and other necessary repository methods) ...
