package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/victor-bologna/pos-curso-go-expert-sqlc/internal/db"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

func main() {
	ctx := context.Background()
	mysql, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	if err != nil {
		panic(err)
	}
	defer mysql.Close()

	queries := db.New(mysql)

	// createCategory(ctx, queries)

	//updateCategory(ctx, queries)

	deleteCategory(ctx, queries)

	listCategories(ctx, queries)
}

func deleteCategory(ctx context.Context, queries *db.Queries) {
	queries.DeleteCategoryByID(ctx, "40b4c810-32b6-4a75-afda-dca90396b9ba")
}

func updateCategory(ctx context.Context, queries *db.Queries) {
	queries.UpdateCategoryByID(ctx, db.UpdateCategoryByIDParams{
		Name: "Backend",
		Description: sql.NullString{
			String: "Backend developent updated",
			Valid:  true,
		},
		ID: "40b4c810-32b6-4a75-afda-dca90396b9ba",
	})
}

func listCategories(ctx context.Context, queries *db.Queries) {
	categories, err := queries.ListCategories(ctx)
	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		fmt.Println(category)
	}
}

func createCategory(ctx context.Context, queries *db.Queries) {
	err := queries.CreateCategory(ctx, db.CreateCategoryParams{
		ID:   uuid.New().String(),
		Name: "Backend",
		Description: sql.NullString{
			String: "Backend development",
			Valid:  true,
		},
	})

	if err != nil {
		panic(err)
	}
}
