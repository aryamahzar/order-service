package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"order-service/models"
	"order-service/service"
	"order-service/utils"

	"github.com/gorilla/mux"
)

type OrderHandler struct {
	orderService service.OrderService
}

func NewOrderHandler(orderService service.OrderService) *OrderHandler {
	return &OrderHandler{
		orderService: orderService,
	}
}

func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var order models.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	createdOrder, err := h.orderService.CreateOrder(r.Context(), &order)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, createdOrder)
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
