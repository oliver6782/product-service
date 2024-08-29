package dto
import (
	"product-service/internal/model"
)

type ProductDTO struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"` 
}

func ToProductDTO(product model.Product) ProductDTO {
	return ProductDTO{
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
	}
}

func ToProduct(productDTO ProductDTO) model.Product {
	return model.Product{
		Name:        productDTO.Name,
		Description: productDTO.Description,
		Price:       productDTO.Price,
		Stock:		 productDTO.Stock,
	}
}
