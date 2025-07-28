package http

import (
	"encoding/json"
	"net/http"

	"desafio_clean_architecture/internal/application/usecase"
)

type OrderHandler struct {
	ListOrdersUC *usecase.ListOrdersUseCase
}

func NewOrderHandler(listUC *usecase.ListOrdersUseCase) *OrderHandler {
	return &OrderHandler{ListOrdersUC: listUC}
}

func (h *OrderHandler) ListOrders(w http.ResponseWriter, r *http.Request) {
	orders, err := h.ListOrdersUC.Execute()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}
