package model

import "errors"

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	Description string  `json:"description"`
}

func (p *Product) SetPrice(price float64) error {
	if price < 0 {
		return errors.New("price cannot be negative")
	}
	p.Price = price
	return nil
}

func (p *Product) GetPrice() float64 {
	return p.Price
}

func (p *Product) SetStock(stock int) error {
	if stock < 0 {
		return errors.New("stock cannot be negative")
	}
	p.Stock = stock
	return nil
}

func (p * Product) GetStock() int {
	return p.Stock
}