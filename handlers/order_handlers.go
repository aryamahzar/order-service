package handlers

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"order-service/models"
	"order-service/service"
	"order-service/utils"

	"github.com/gorilla/mux"
)

type OrderHandler struct {
	orderService  service.OrderService
	apiGatewayURL string
	httpClient    *http.Client
}

func NewOrderHandler(orderService service.OrderService, apiGatewayURL string) *OrderHandler {
	// Create an HTTP client with a custom transport to skip SSL verification
	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	return &OrderHandler{
		orderService:  orderService,
		apiGatewayURL: apiGatewayURL,
		httpClient:    httpClient,
	}
}

func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var order models.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	// Validate order details
	if order.UserID == "" || len(order.Items) == 0 {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid order details")
		return
	}

	// Check product availability with Inventory Service via API Gateway
	for _, item := range order.Items {

		fmt.Printf("item : %+v\n", item)

		available, err := h.checkProductAvailability(item.ProductID, item.Quantity)
		fmt.Println("available : ", available)
		fmt.Println("err : ", err)
		if err != nil || !available {
			utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Product %s is not available", item.ProductID))
			return
		}
	}

	// Calculate total amount
	totalAmount := 0.0
	for _, item := range order.Items {
		totalAmount += item.UnitPrice * float64(item.Quantity)
	}
	order.TotalAmount = totalAmount

	createdOrder, err := h.orderService.CreateOrder(r.Context(), &order)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, createdOrder)
}

func (h *OrderHandler) checkProductAvailability(productID string, quantity int) (bool, error) {

	// Configure a custom transport to skip SSL verification
	insecureTransport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true, // Ignore insecure SSL certificates
		},
	}

	insecureClient := &http.Client{Transport: insecureTransport}

	url := fmt.Sprintf("%s/products/%s", h.apiGatewayURL, productID)
	fmt.Println("url : ", url)
	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return false, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := insecureClient.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("failed to check product availability")
	}

	var product struct {
		ProductID string `json:"product_id"`
		Quantity  int    `json:"quantity"`
		Name      string `json:"name"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&product); err != nil {
		fmt.Println("error decode check availability  :", err)
		return false, err
	}
	fmt.Printf("product : %+v\n", product)
	return product.Quantity >= quantity, nil

}

func (h *OrderHandler) ListOrders(w http.ResponseWriter, r *http.Request) {
	// Parse query parameters for filtering and pagination
	userID := r.URL.Query().Get("userId")
	statusStr := r.URL.Query().Get("status")
	pageStr := r.URL.Query().Get("page")
	sizeStr := r.URL.Query().Get("size")

	// Convert query parameters to appropriate types
	page, _ := strconv.ParseInt(pageStr, 10, 64)
	if page == 0 {
		page = 1 // Default to page 1 if not provided or invalid
	}
	size, _ := strconv.ParseInt(sizeStr, 10, 64)
	if size == 0 {
		size = 20 // Default to size 20 if not provided or invalid
	}

	var status *models.OrderStatus
	if statusStr != "" {
		statusValue := models.OrderStatus(statusStr)
		status = &statusValue
	}

	orders, totalCount, err := h.orderService.ListOrders(r.Context(), &userID, status, page, size)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	response := map[string]interface{}{
		"content":       orders,
		"totalPages":    (totalCount + size - 1) / size, // Calculate total pages
		"totalElements": totalCount,
	}

	utils.RespondWithJSON(w, http.StatusOK, response)
}

func (h *OrderHandler) GetOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderID := vars["orderId"]

	order, err := h.orderService.GetOrder(r.Context(), orderID)
	if err != nil {
		if err.Error() == "order not found" {
			utils.RespondWithError(w, http.StatusNotFound, err.Error())
		} else {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, order)
}

func (h *OrderHandler) UpdateOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderID := vars["orderId"]

	var order models.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}
	order.OrderID = orderID // Ensure the orderId is set in the order object

	if err := h.orderService.UpdateOrder(r.Context(), &order); err != nil {
		if err.Error() == "order not found" {
			utils.RespondWithError(w, http.StatusNotFound, err.Error())
		} else {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, order)
}

func (h *OrderHandler) CancelOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderID := vars["orderId"]

	updatedOrder, err := h.orderService.CancelOrder(r.Context(), orderID)
	if err != nil {
		if err.Error() == "order not found" {
			utils.RespondWithError(w, http.StatusNotFound, err.Error())
		} else {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, updatedOrder)
}
