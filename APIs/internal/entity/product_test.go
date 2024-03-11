package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	product, err := NewProduct("Test", 10.0)
	assert.Nil(t, err)
	assert.NotEmpty(t, product.ID)
	assert.NotEmpty(t, product.CreatedAt)
	assert.Equal(t, "Test", product.Name)
	assert.Equal(t, 10.0, product.Price)
	assert.Nil(t, product.Validate())
}

func TestValidateProduct_NameRequired(t *testing.T) {
	product, err := NewProduct("", 10.0)
	assert.Nil(t, product)
	assert.Equal(t, ErrProductNameRequired, err)
}

func TestValidateProduct_PriceInvalid(t *testing.T) {
	product, err := NewProduct("Test", -10.0)
	assert.Nil(t, product)
	assert.Equal(t, ErrProductPriceInvalid, err)
}

func TestValidateProduct_PriceRequired(t *testing.T) {
	product, err := NewProduct("Test", 0.0)
	assert.Nil(t, product)
	assert.Equal(t, ErrProductPriceRequired, err)
}
