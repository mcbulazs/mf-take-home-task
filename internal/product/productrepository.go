package product

import (
	"database/sql"

	scripts "mcbulazs/mf-take-home-task/sql"
)

type SQLProductRepository struct {
	DB *sql.DB
}

func NewSQLProductRepository(db *sql.DB) *SQLProductRepository {
	return &SQLProductRepository{
		DB: db,
	}
}

func (r *SQLProductRepository) ListProducts() ([]Product, error) {
	rows, err := r.DB.Query(scripts.ListProductsSQL)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var product Product
		err = rows.Scan(&product.SKU, &product.Name, &product.Stock)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
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
