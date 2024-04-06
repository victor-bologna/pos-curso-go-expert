package database

import (
	"database/sql"
	"fmt"
	"math/rand"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/victor-bologna/pos-curso-go-expert-apis/internal/entity"
	"github.com/victor-bologna/pos-curso-go-expert-apis/pkg/property"
	"gorm.io/gorm"
	_ "modernc.org/sqlite"
)

func openProductDB(t *testing.T) *gorm.DB {
	db, err := sql.Open("sqlite", "./data.db")
	if err != nil {
		t.Error(err)
	}

	if err = db.AutoMigrate(&entity.Product{}); err != nil {
		t.Error(err)
	}
	return db
}

func TestCreateProduct(t *testing.T) {
	db := openProductDB(t)
	product, err := entity.NewProduct("Test", 10.0)
	assert.NoError(t, err)
	productDB := NewProductDB(db)

	err = productDB.Create(product)
	assert.NoError(t, err)

	var productResult entity.Product
	err = db.Where("id = ?", product.ID).Find(&productResult).Error
	assert.NoError(t, err)
	assert.Equal(t, product.ID, productResult.ID)
	assert.Equal(t, product.Name, productResult.Name)
	assert.Equal(t, product.Price, productResult.Price)
	assert.NotNil(t, productResult.CreatedAt)
}

func TestFindAllProducts(t *testing.T) {
	var products []entity.Product
	db := openProductDB(t)
	for i := 1; i < 26; i++ {
		product, err := entity.NewProduct(fmt.Sprintf("Product %v", i), rand.Float64()*100)
		assert.NoError(t, err)
		products = append(products, *product)
	}
	productDB := NewProductDB(db)
	db.Create(products)

	productsResult, err := productDB.FindAll(1, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, productsResult, 10)
	assert.Equal(t, "Product 1", productsResult[0].Name)
	assert.Equal(t, "Product 10", productsResult[9].Name)

	productsResult, err = productDB.FindAll(2, 10, "desc")
	assert.NoError(t, err)
	assert.Len(t, productsResult, 10)
	assert.Equal(t, "Product 11", productsResult[0].Name)
	assert.Equal(t, "Product 20", productsResult[9].Name)

	productsResult, err = productDB.FindAll(3, 10, "test")
	assert.NoError(t, err)
	assert.Len(t, productsResult, 5)
	assert.Equal(t, "Product 21", productsResult[0].Name)
	assert.Equal(t, "Product 25", productsResult[4].Name)

	productsResult, err = productDB.FindAll(0, 0, "")
	assert.NoError(t, err)
	assert.Len(t, productsResult, 25)
	assert.Equal(t, "Product 1", productsResult[0].Name)
	assert.Equal(t, "Product 25", productsResult[24].Name)
}

func TestProductFindByID(t *testing.T) {
	db := openProductDB(t)
	product, err := entity.NewProduct("Test", 15.23)
	assert.NoError(t, err)

	db.Create(product)
	productDB := NewProductDB(db)

	productResult, err := productDB.FindByID(product.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, product.ID, productResult.ID)
	assert.Equal(t, product.Name, productResult.Name)
	assert.Equal(t, product.Price, productResult.Price)
	assert.NotNil(t, productResult.CreatedAt)
}

func TestProductUpdate(t *testing.T) {
	db := openProductDB(t)
	product, err := entity.NewProduct("Test", 15.23)
	assert.NoError(t, err)

	db.Create(product)
	productDB := NewProductDB(db)

	product.Price = 40.50
	err = productDB.Update(product)
	assert.NoError(t, err)

	var product2 entity.Product
	err = productDB.DB.Where("ID = ?", product.ID.String()).First(&product2).Error
	assert.NoError(t, err)
	assert.Equal(t, product.Price, product2.Price)
}

func TestProductUpdate_NotFound(t *testing.T) {
	db := openProductDB(t)
	product, err := entity.NewProduct("Test", 15.23)
	assert.NoError(t, err)
	db.Create(product)
	product.ID, _ = property.ParseID("")

	productDB := NewProductDB(db)

	err = productDB.Update(product)
	assert.Error(t, err)
	assert.EqualError(t, err, "record not found")
}

func TestProductUpdate_PriceZero(t *testing.T) {
	db := openProductDB(t)
	product, err := entity.NewProduct("Test", 15.23)
	assert.NoError(t, err)
	db.Create(product)

	err = db.Where("id = ?", product.ID).Find(&product).Error
	assert.NoError(t, err)
	productDB := NewProductDB(db)

	product.Price = 0
	err = productDB.Update(product)
	assert.Error(t, err)
	assert.EqualError(t, err, entity.ErrProductPriceRequired.Error())
}

func TestProductUpdate_PriceInvalid(t *testing.T) {
	db := openProductDB(t)
	product, err := entity.NewProduct("Test", 15.23)
	assert.NoError(t, err)
	db.Create(product)

	err = db.Where("id = ?", product.ID).Find(&product).Error
	assert.NoError(t, err)
	productDB := NewProductDB(db)

	product.Price = -10
	err = productDB.Update(product)
	assert.Error(t, err)
	assert.EqualError(t, err, entity.ErrProductPriceInvalid.Error())
}

func TestProductDeleteByID(t *testing.T) {
	db := openProductDB(t)
	product, err := entity.NewProduct("Test", 15.23)
	assert.NoError(t, err)

	db.Create(product)
	productDB := NewProductDB(db)

	err = productDB.DeleteByID(product.ID.String())
	assert.NoError(t, err)

	var product2 entity.Product
	err = productDB.DB.Where("ID = ?", product.ID.String()).First(&product2).Error
	assert.Error(t, err)
	assert.Equal(t, uuid.Nil, product2.ID)
	assert.Equal(t, "", product2.Name)
	assert.Equal(t, 0.0, product2.Price)
}

func TestProductDeleteByID_NotFound(t *testing.T) {
	db := openProductDB(t)
	product, err := entity.NewProduct("Test", 15.23)
	assert.NoError(t, err)

	productDB := NewProductDB(db)

	err = productDB.DeleteByID(product.ID.String())
	assert.Error(t, err)
}
