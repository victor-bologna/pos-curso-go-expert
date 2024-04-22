package product

import "database/sql"

type ProductRepositoryInterface interface {
	GetProductId(id int) (Product, error)
}

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) GetProductId(id int) (Product, error) {
	return Product{
		ID:   1,
		Name: "Product Name",
	}, nil
}
