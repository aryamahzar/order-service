/*
 * Order Service API
 *
 * API for managing orders in the e-commerce platform
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type InlineResponse200 struct {
	Content []Order `json:"content,omitempty"`
	TotalPages int32 `json:"totalPages,omitempty"`
	TotalElements int32 `json:"totalElements,omitempty"`
}
