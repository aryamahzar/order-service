package service

import (
	"context"
	"errors"
	"fmt"

	"order-service/models"
	"order-service/repository"

)




type OrderService struct {
	orderRepository repository.OrderRepository
}

func NewOrderService(orderRepository repository.OrderRepository) *OrderService {
	return &OrderService{
		orderRepository: orderRepository,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, order *models.Order) (*models.Order, error) {
	// Basic input validation (you might want more comprehensive validation)
	if order.UserID == "" {
		return nil, errors.New("userId is required")
	}
	if len(order.Items) == 0 {
		return nil, errors.New("at least one order item is required")
	}

	// Calculate total amount (assuming you have product prices available)
	totalAmount := 0.0
	for _, item := range order.Items {
		// ... fetch product price based on item.ProductID (you might need a ProductService for this) ...
		totalAmount += item.UnitPrice * float64(item.Quantity)
	}
	order.TotalAmount = totalAmount

	// Set initial order status
	order.Status = models.OrderStatusCreated

	// Persist the order
	if err := s.orderRepository.CreateOrder(ctx, order); err != nil {
		return nil, fmt.Errorf("failed to create order: %v", err)
	}

	return order, nil
}

func (s *OrderService) GetOrder(ctx context.Context, orderID string) (*models.Order, error) {
	return s.orderRepository.GetOrder(ctx, orderID)
}

func (s *OrderService) UpdateOrder(ctx context.Context, order *models.Order) error {
	// Check if the order exists
	existingOrder, err := s.orderRepository.GetOrder(ctx, order.OrderID)
	if err != nil {
		return err
	}

	// Perform any necessary status transition checks or business logic
	if existingOrder.Status == models.OrderStatusDelivered {
		return errors.New("cannot update a delivered order")
	}

	// ... other business logic or validation ...

	return s.orderRepository.UpdateOrder(ctx, order)
}

func (s *OrderService) ListOrders(ctx context.Context, userID *string, status *models.OrderStatus, page, size int64) ([]*models.Order, int64, error) {
	return s.orderRepository.ListOrders(ctx, userID, status, page, size)
}

// ... (Add CancelOrder and other necessary service methods) ...
func (s *OrderService) CancelOrder(ctx context.Context, orderID string) (*models.Order, error) {
	// Fetch the order
	order, err := s.orderRepository.GetOrder(ctx, orderID)
	if err != nil {
		return nil, err
	}

	// Check if the order can be canceled (based on its current status)
	if order.Status == models.OrderStatusDelivered || order.Status == models.OrderStatusCancelled {
		return nil, errors.New("cannot cancel an order that is already delivered or canceled")
	}

	// Update the order status to canceled
	order.Status = models.OrderStatusCancelled

	// ... Perform any other necessary cancellation logic (e.g., inventory adjustments, notifications) ...

	// Update the order in the repository
	if err := s.orderRepository.UpdateOrder(ctx, order); err != nil {
		return nil, fmt.Errorf("failed to cancel order: %v", err)
	}

	return order, nil
}
