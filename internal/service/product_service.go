package service

import (
    "product-service/internal/model"
    "product-service/internal/repository"
    "product-service/internal/dto"
    "errors"
)

var ErrProductNotFound = errors.New("product not found")

type ProductService struct {
    repo *repository.ProductRepository 
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
    return &ProductService{
        repo: repo,
    }
}

func (s *ProductService) GetProducts() ([]model.Product, error) {
    products, err := s.repo.GetProducts()
    if err != nil {
        return nil, err
    }
    return products, nil
}

func (s *ProductService) CreateProduct(productDTO dto.ProductDTO) (model.Product, error) {
    product := dto.ToProduct(productDTO)
    createdProduct, err := s.repo.CreateProduct(product)
    if err != nil {
        return model.Product{}, err
    }
    return createdProduct, nil
}

func (s *ProductService) GetProductById(id string) (dto.ProductDTO, error) {
	product, err := s.repo.GetProductById(id)
	if err != nil {
		return dto.ProductDTO{}, err
	}
	return dto.ToProductDTO(product), nil
}

func (s *ProductService) UpdateProduct(id string, productDTO dto.ProductDTO) (dto.ProductDTO, error) {
	// First, check if the product exists
	existingProduct, err := s.repo.GetProductById(id)
	if err != nil {
		if errors.Is(err, ErrProductNotFound) {
			return dto.ProductDTO{}, ErrProductNotFound
		}
		return dto.ProductDTO{}, err
	}

	// Update the existing product with new data
	existingProduct.Name = productDTO.Name
	existingProduct.Description = productDTO.Description
	existingProduct.Price = productDTO.Price
	existingProduct.Stock = productDTO.Stock

	updatedProduct, err := s.repo.UpdateProduct(id, existingProduct)
	if err != nil {
		return dto.ProductDTO{}, err
	}

	return dto.ToProductDTO(updatedProduct), nil
}

func (s *ProductService) DeleteProduct(id string) error {
	// Check if the product exists before attempting to delete
	_, err := s.repo.GetProductById(id)
	if err != nil {
		if errors.Is(err, ErrProductNotFound) {
			return ErrProductNotFound
		}
		return err
	}
	
	return s.repo.DeleteProduct(id)
}
