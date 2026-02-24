package product_test

import (
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"

	"mcbulazs/mf-take-home-task/internal/product"
	scripts "mcbulazs/mf-take-home-task/sql"
)

func newMockDB(t *testing.T) (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherRegexp))
	require.NoError(t, err)
	return db, mock
}

func TestListProducts(t *testing.T) {
	db, mock := newMockDB(t)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"sku", "name", "stock"}).
		AddRow("sku1", "Product 1", 10).
		AddRow("sku2", "Product 2", 5)

	mock.ExpectQuery(regexp.QuoteMeta(scripts.ListProductsSQL)).WillReturnRows(rows)

	repo := product.NewSQLProductRepository(db)
	products, err := repo.ListProducts()
	require.NoError(t, err)
	require.Len(t, products, 2)
	require.Equal(t, "sku1", products[0].SKU)
	require.Equal(t, 5, products[1].Stock)

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestApplyMovement_Increase(t *testing.T) {
	db, mock := newMockDB(t)
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(scripts.GetStockMovementSQL)).
		WithArgs("req-001", "sku1", 5).
		WillReturnError(sql.ErrNoRows)

	mock.ExpectQuery(regexp.QuoteMeta(scripts.UpdateProductStockSQL)).
		WithArgs("sku1", 5).
		WillReturnRows(sqlmock.NewRows([]string{"stock"}).AddRow(15))

	mock.ExpectExec(regexp.QuoteMeta(scripts.InsertStockMovementSQL)).
		WithArgs("req-001", "sku1", 5, "production").
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectCommit()

	repo := product.NewSQLProductRepository(db)
	applied, newVal, err := repo.ApplyMovement("req-001", "sku1", 5, "production")
	require.NoError(t, err)
	require.True(t, applied)
	require.Equal(t, 15, newVal)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestApplyMovement_AlreadyApplied(t *testing.T) {
	db, mock := newMockDB(t)
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(scripts.GetStockMovementSQL)).
		WithArgs("req-002", "sku1", 3).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))
	mock.ExpectRollback()

	repo := product.NewSQLProductRepository(db)
	applied, _, err := repo.ApplyMovement("req-002", "sku1", 3, "order")
	require.NoError(t, err)
	require.False(t, applied)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetTopProducts(t *testing.T) {
	db, mock := newMockDB(t)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"sku", "name", "stock"}).
		AddRow("sku1", "Product 1", 50).
		AddRow("sku2", "Product 2", 40)

	mock.ExpectQuery(regexp.QuoteMeta(scripts.TopNProductsSQL)).
		WithArgs(2).
		WillReturnRows(rows)

	repo := product.NewSQLProductRepository(db)
	products, err := repo.GetTopProducts(2)
	require.NoError(t, err)
	require.Len(t, products, 2)
	require.Equal(t, 50, products[0].Stock)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetLowStockProducts(t *testing.T) {
	db, mock := newMockDB(t)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"sku", "name", "stock"}).
		AddRow("sku3", "Product 3", 2).
		AddRow("sku4", "Product 4", 1)

	mock.ExpectQuery(regexp.QuoteMeta(scripts.LowestProductsUnderSQL)).
		WithArgs(5).
		WillReturnRows(rows)

	repo := product.NewSQLProductRepository(db)
	products, err := repo.GetLowStockProducts(5)
	require.NoError(t, err)
	require.Len(t, products, 2)
	require.Equal(t, "sku3", products[0].SKU)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestCountProducts(t *testing.T) {
	db, mock := newMockDB(t)
	defer db.Close()

	mock.ExpectQuery(regexp.QuoteMeta(scripts.CountProductsSQL)).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(42))

	repo := product.NewSQLProductRepository(db)
	count, err := repo.CountProducts()
	require.NoError(t, err)
	require.Equal(t, 42, count)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestSumStock(t *testing.T) {
	db, mock := newMockDB(t)
	defer db.Close()

	mock.ExpectQuery(regexp.QuoteMeta(scripts.SumProductStockSQL)).
		WillReturnRows(sqlmock.NewRows([]string{"sum"}).AddRow(123))

	repo := product.NewSQLProductRepository(db)
	sum, err := repo.SumStock()
	require.NoError(t, err)
	require.Equal(t, 123, sum)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestApplyMovement_NegativeStock(t *testing.T) {
	db, mock := newMockDB(t)
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(scripts.GetStockMovementSQL)).
		WithArgs("req-003", "sku5", -100).
		WillReturnError(sql.ErrNoRows)

	mock.ExpectQuery(regexp.QuoteMeta(scripts.UpdateProductStockSQL)).
		WithArgs("sku5", -100).
		WillReturnError(sql.ErrNoRows)
	mock.ExpectRollback()

	repo := product.NewSQLProductRepository(db)
	applied, _, err := repo.ApplyMovement("req-003", "sku5", -100, "order")
	require.Error(t, err)
	require.False(t, applied)
	require.Equal(t, "SKU not found or stock would go to negative", err.Error())
	require.NoError(t, mock.ExpectationsWereMet())
}
