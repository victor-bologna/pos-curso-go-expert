package main

import (
	"net/http"

	"github.com/glebarez/sqlite"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/victor-bologna/pos-curso-go-expert-apis/configs"
	"github.com/victor-bologna/pos-curso-go-expert-apis/internal/entity"
	"github.com/victor-bologna/pos-curso-go-expert-apis/internal/infra/database"
	"github.com/victor-bologna/pos-curso-go-expert-apis/internal/infra/webservers/handlers"
	"gorm.io/gorm"
)

func main() {
	_ = configs.LoadConfig(".")

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.User{}, &entity.Product{})

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	productDB := database.NewProductDB(db)
	productHandler := handlers.NewProductHandler(productDB)

	r.Post("/products", productHandler.CreateProduct)
	r.Get("/products/{id}", productHandler.FindProductByID)
	r.Get("/products", productHandler.FindAllProducts)
	r.Put("/products/{id}", productHandler.UpdateProduct)
	r.Delete("/products/{id}", productHandler.DeleteProduct)
	http.ListenAndServe(":8000", r)
}
