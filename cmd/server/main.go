package main

import (
	"net"

	"github.com/gorilla/mux"
	"log"
	"net/http"
	"product-service/internal/config"
	"product-service/internal/handler"
	"product-service/internal/repository"
	"product-service/internal/service"
	"product-service/pkg/db"
	pb "product-service/api/gen/go/grpc"
	"google.golang.org/grpc"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	database, err := db.Connect(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Initialize service and handler
	productRepository := repository.NewProductRepository(database)
	productService := service.NewProductService(productRepository)
	httpHandler := handler.NewHandler(productService)

	r := mux.NewRouter()

	// Set up routes
	r.HandleFunc("/product", httpHandler.GetProducts).Methods("GET")
	r.HandleFunc("/product", httpHandler.CreateProduct).Methods("POST")
	r.HandleFunc("/product/{id}", httpHandler.GetProductById).Methods("GET")
	r.HandleFunc("/product/{id}", httpHandler.UpdateProduct).Methods("PUT")
	r.HandleFunc("/product/{id}", httpHandler.DeleteProduct).Methods("DELETE")

	// Set up gRPC server
	grpcServer := grpc.NewServer()
	pb.RegisterProductServiceServer(grpcServer, &grpcHandler{productService: productService})

	// Start HTTP server
	go func() {
		log.Println("Starting HTTP server on :8080")
		if err := http.ListenAndServe(":8080", r); err != nil {
			log.Fatalf("Failed to start HTTP server: %v", err)
		}
	}()

	// Start gRPC server
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen for gRPC: %v", err)
	}
	log.Println("Starting gRPC server on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start gRPC server: %v", err)
	}
}

// grpcHandler implements the gRPC ProductServiceServer interface
type grpcHandler struct {
	pb.UnimplementedProductServiceServer
	productService *service.ProductService
}

// // Implement your gRPC methods here, for example:
// func (h *grpcHandler) GetProductInfo(ctx context.Context, req *pb.ProductRequest) (*pb.ProductReply, error) {
// 	product, err := h.productService.GetProductByID(req.Id)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &pb.ProductReply{
// 		Id:          product.ID,
// 		Name:        product.Name,
// 		Description: product.Description,
// 		Price:       float32(product.Price),
// 	}, nil
// }

