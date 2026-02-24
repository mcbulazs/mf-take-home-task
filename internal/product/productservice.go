package product


import (
	"fmt"
)

type ProductRepository interface {
	ListProducts() ([]Product, error)
	GetTopProducts(limit int) ([]Product, error)
	GetLowStockProducts(threshold int) ([]Product, error)

	ApplyMovement(id string, sku string, movement int, reason string) (applied bool, err error)

	CountProducts() (int, error)
	SumStock() (int, error)
}

type ProductService struct {
	Repo ProductRepository
}

func NewProductService(repo ProductRepository) *ProductService {
	return &ProductService{
		Repo: repo,
	}
}

func (s *ProductService) List() error {
	products, err := s.Repo.ListProducts()
	if err != nil {
		return err
	}
	printProducts(products)

	return nil
}

func (s *ProductService) IncreaseQuantity(id string, sku string, qty int, reason string) error {
	ok, err := s.Repo.ApplyMovement(id, sku, qty, reason)
	if err != nil {
		return err
	}
	if !ok {
		fmt.Println("already applied")
		return nil
	}
	return nil
}

func (s *ProductService) DecreaseQuantity(id string, sku string, qty int, reason string) error {
	ok, err := s.Repo.ApplyMovement(id, sku, -qty, reason)
	if err != nil {
		return err
	}
	if !ok {
		fmt.Println("already applied")
		return nil
	}
	return nil
}

func (s *ProductService) Report(top int, lowStock int) error {
	count, err := s.Repo.CountProducts()
	if err != nil {
		return err
	}
	sum, err := s.Repo.SumStock()
	if err != nil {
		return err
	}
	topProducts, err := s.Repo.GetTopProducts(top)
	if err != nil {
		return err
	}
	lowProducts, err := s.Repo.GetLowStockProducts(lowStock)
	if err != nil {
		return err
	}
	fmt.Printf("Product count: %d\n", count)
	fmt.Printf("Product quantity sum: %d\n", sum)
	fmt.Printf("Top %d products by stock\n", top)
	printProducts(topProducts)
	fmt.Printf("Products with stock lower then %d\n", lowStock)
	printProducts(lowProducts)

	return nil
}

func printProducts(products []Product) {
	fmt.Println("SKU | Name | Stock")
	for _, p := range products {
		fmt.Printf("%s | %s | %d\n", p.SKU, p.Name, p.Stock)
	}
}
