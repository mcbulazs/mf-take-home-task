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

	return getProductsFromSQLRows(rows)
}

func (r *SQLProductRepository) ApplyMovement(id string, sku string, movement int, reason string) (applied bool, err error) {
	return false, nil
}

func (r *SQLProductRepository) GetTopProducts(limit int) ([]Product, error) {
	rows, err := r.DB.Query(scripts.TopNProductsSQL, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return getProductsFromSQLRows(rows)
}

func (r *SQLProductRepository) GetLowStockProducts(threshold int) ([]Product, error) {
	rows, err := r.DB.Query(scripts.LowestProductsUnderSQL, threshold)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return getProductsFromSQLRows(rows)
}

func (r *SQLProductRepository) CountProducts() (int, error) {
	var count int
	err := r.DB.QueryRow(scripts.CountProductsSQL).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *SQLProductRepository) SumStock() (int, error) {
	var sum int
	err := r.DB.QueryRow(scripts.SumProductStockSQL).Scan(&sum)
	if err != nil {
		return 0, err
	}
	return sum, nil
}

func getProductsFromSQLRows(rows *sql.Rows) ([]Product, error) {
	var products []Product
	for rows.Next() {
		var product Product
		err := rows.Scan(&product.SKU, &product.Name, &product.Stock)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}
