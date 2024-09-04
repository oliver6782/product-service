package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"product-service/internal/dto"
	"product-service/internal/service"
)

type Handler struct {
	productService *service.ProductService
}

// NewHandler constructor
func NewHandler(productService *service.ProductService) *Handler {
	return &Handler{
		productService: productService,
	}
}

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.productService.GetProducts()
	if err != nil {
		http.Error(w, "Failed to get products", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(products)
	if err != nil {
		return
	}
}

func (h *Handler) GetProductById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	product, err := h.productService.GetProductById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(product)
	if err != nil {
		return
	}
}

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var productDTO dto.ProductDTO
	if err := json.NewDecoder(r.Body).Decode(&productDTO); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	createdProduct, err := h.productService.CreateProduct(productDTO)
	if err != nil {
		http.Error(w, "Failed to create product", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(createdProduct)
	if err != nil {
		return
	}
}

func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var productDTO dto.ProductDTO
	if err := json.NewDecoder(r.Body).Decode(&productDTO); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	updatedProduct, err := h.productService.UpdateProduct(id, productDTO)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(updatedProduct)
	if err != nil {
		return
	}
}

func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if err := h.productService.DeleteProduct(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}


