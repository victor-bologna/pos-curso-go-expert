package database

import "github.com/victor-bologna/pos-curso-go-expert-apis/internal/entity"

type UserInterface interface {
	CreateUser(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}

type ProductInterface interface {
	CreateProduct(product *entity.Product) error
	FindAll(page, limit int, sort string) ([]entity.Product, error)
	FindByID(id string) (*entity.Product, error)
	Update(product *entity.Product) error
	DeleteByID(id string) error
}
