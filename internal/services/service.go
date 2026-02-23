package services

import "database/sql"

type ProductService struct {
	DB *sql.DB
}

func NewProductService(db *sql.DB) *ProductService {
	return &ProductService{
		DB: db,
	}
}

func (s ProductService) List() error {
	return nil
}

func (s ProductService) IncreseQuantity(id string, sku string, qty int, reason string) error {
	return nil
}

func (s ProductService) DecreaseQuantity(id string, sku string, qty int, reason string) error {
	return nil
}

func (s ProductService) Report(top int, lowStock int) error {
	return nil
}
