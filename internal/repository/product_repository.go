package repository

import (
	"gorm.io/gorm"
	"product-service/internal/model"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) GetProducts() ([]model.Product, error) {
	var products []model.Product
	result := r.db.Find(&products)
	return products, result.Error
}

func (r *ProductRepository) CreateProduct(product model.Product) (model.Product, error) {
	if err := r.db.Save(&product).Error; err != nil {
		return model.Product{}, err
	}
	return product, nil
}

func (r *ProductRepository) GetProductById(id string) (model.Product, error) {
	var product model.Product
	result := r.db.First(&product, id)
	return product, result.Error
}

func (r *ProductRepository) UpdateProduct(id string, product model.Product) (model.Product, error) {
	if err := r.db.First(&product, id).Error; err != nil {
		return model.Product{}, err
	}

	result := r.db.Save(&product)
	return product, result.Error
}

func (r *ProductRepository) DeleteProduct(id string) error {
	result := r.db.Delete(&model.Product{}, id)
	return result.Error
}


