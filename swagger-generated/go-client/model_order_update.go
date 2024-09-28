/*
 * Order Service API
 *
 * API for managing orders in the e-commerce platform
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type OrderUpdate struct {
	Status *OrderStatus `json:"status,omitempty"`
	ShippingAddress *Address `json:"shippingAddress,omitempty"`
	BillingAddress *Address `json:"billingAddress,omitempty"`
}
