package main

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

func main() {
	db, err := sql.Open("sqlite", "./database.db")
	if err != nil {
		panic(err)
	}

	/* Without Google Wire
	// Create Repo
	repository := product.NewProductRepository(db)

	// Create Use Case
	useCase := product.NewProductUseCase(repository)
	*/

	useCase := NewUseCase(db)

	// Get Product
	product, err := useCase.GetProductId(1)
	if err != nil {
		panic(err)
	}

	fmt.Println(product.Name)
}
