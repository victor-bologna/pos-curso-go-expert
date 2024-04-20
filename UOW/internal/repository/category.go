package repository

import (
	"context"
	"database/sql"

	"github.com/victor-bologna/pos-curso-go-expert-uow/internal/db"
	"github.com/victor-bologna/pos-curso-go-expert-uow/internal/entity"
)

type CategoryRepositoryInterface interface {
	Insert(ctx context.Context, category entity.Category) error
}

type CategoryRepository struct {
	DB      *sql.DB
	Queries *db.Queries
}

func NewCategoryRepository(database *sql.DB) *CategoryRepository {
	return &CategoryRepository{DB: database, Queries: db.New(database)}
}

func (cr *CategoryRepository) Insert(ctx context.Context, category entity.Category) error {
	return cr.Queries.CreateCategory(ctx, db.CreateCategoryParams{
		Name: category.Name,
	})
}
