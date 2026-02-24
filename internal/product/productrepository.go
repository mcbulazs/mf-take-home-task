package product

import "database/sql"

type SQLProductRepository struct {
	DB *sql.DB
}

func NewSQLProductRepository(db *sql.DB) *SQLProductRepository {
	return &SQLProductRepository{
		DB: db,
	}
}

func (r *SQLProductRepository) ListProducts() ([]Product, error) {
	return nil, nil
}

func (r *SQLProductRepository) GetTopProducts(limit int) ([]Product, error) {
	return nil, nil
}

func (r *SQLProductRepository) GetLowStockProducts(threshold int) ([]Product, error) {
	return nil, nil
}

func (r *SQLProductRepository) ApplyMovement(id string, sku string, movement int, reason string) (applied bool, err error) {
	return false, nil
}

func (r *SQLProductRepository) CountProducts() (int, error) {
	return 0, nil
}

func (r *SQLProductRepository) SumStock() (int, error) {
	return 0, nil
}
