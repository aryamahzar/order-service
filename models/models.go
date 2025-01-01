package models

import "time"

type Order struct {
	OrderID         string      `json:"orderId" bson:"orderId"`
	UserID          string      `json:"userId" bson:"userId"`
	Status          OrderStatus `json:"status" bson:"status"`
	OrderDate       time.Time   `json:"orderDate" bson:"orderDate"`
	TotalAmount     float64     `json:"totalAmount" bson:"totalAmount"`
	Items           []OrderItem `json:"items" bson:"items"`
	ShippingAddress Address     `json:"shippingAddress" bson:"shippingAddress"`
	BillingAddress  Address     `json:"billingAddress" bson:"billingAddress"`
}

type OrderItem struct {
	ProductID string  `json:"product_id" bson:"product_id"`
	Quantity  int     `json:"quantity" bson:"quantity"`
	UnitPrice float64 `json:"unitPrice" bson:"unitPrice"`
}

type Address struct {
	Street     string `json:"street" bson:"street"`
	City       string `json:"city" bson:"city"`
	State      string `json:"state" bson:"state"`
	PostalCode string `json:"postalCode" bson:"postalCode"`
	Country    string `json:"country" bson:"country"`
}

type OrderStatus string

const (
	OrderStatusCreated   OrderStatus = "CREATED"
	OrderStatusPaid      OrderStatus = "PAID"
	OrderStatusShipped   OrderStatus = "SHIPPED"
	OrderStatusDelivered OrderStatus = "DELIVERED"
	OrderStatusCancelled OrderStatus = "CANCELLED"
)

type Product struct {
	ProductID string `json:"product_id"`
	Name      string `json:"name"`
	Quantity  int    `json:"quantity"`
}
