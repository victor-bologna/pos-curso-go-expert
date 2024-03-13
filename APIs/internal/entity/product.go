package entity

import (
	"errors"
	"time"

	"github.com/victor-bologna/pos-curso-go-expert-apis/pkg/property"
)

var (
	ErrProductIDRequired    = errors.New("Product ID is required")
	ErrProductNameRequired  = errors.New("Product name is required")
	ErrProductPriceRequired = errors.New("Product price is required")
	ErrProductIDInvalid     = errors.New("Product ID is invalid")
	ErrProductPriceInvalid  = errors.New("Product price invalid")
)

type Product struct {
	ID        property.ID
	Name      string
	Price     float64
	CreatedAt time.Time
}

func NewProduct(name string, price float64) (*Product, error) {
	product := &Product{
		ID:        property.NewUUID(),
		Name:      name,
		Price:     price,
		CreatedAt: time.Now(),
	}
	err := product.Validate()
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (p *Product) Validate() error {
	if p.ID.String() == "" {
		return ErrProductIDRequired
	}
	if _, err := property.ParseID(p.ID.String()); err != nil {
		return ErrProductIDInvalid
	}
	if p.Name == "" {
		return ErrProductNameRequired
	}
	if p.Price == 0.0 {
		return ErrProductPriceRequired
	}
	if p.Price < 0 {
		return ErrProductPriceInvalid
	}
	return nil
}
