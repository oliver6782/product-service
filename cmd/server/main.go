package main

import (
	"log"
	"net/http"
	"product-service/internal/handler"
	"product-service/pkg/db"
	"product-service/internal/config"
	"product-service/internal/service"
	"product-service/internal/repository"
	"github.com/gorilla/mux"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	db, err := db.Connect(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Initialize service and handler
	productRepository := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepository)
	handler := handler.NewHandler(productService)

	r := mux.NewRouter()

	// Set up routes
	r.HandleFunc("/product", handler.GetProducts).Methods("GET")
	r.HandleFunc("/product", handler.CreateProduct).Methods("POST")
	r.HandleFunc("/product/{id}", handler.GetProductById).Methods("GET")
	r.HandleFunc("/product/{id}", handler.UpdateProduct).Methods("PUT")
	r.HandleFunc("/product/{id}", handler.DeleteProduct).Methods("DELETE")

	// Start the server
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
