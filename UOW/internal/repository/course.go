package repository

import (
	"context"
	"database/sql"

	"github.com/victor-bologna/pos-curso-go-expert-uow/internal/db"
	"github.com/victor-bologna/pos-curso-go-expert-uow/internal/entity"
)

type CourseRepositoryInterface interface {
	Insert(ctx context.Context, course entity.Course) error
}

type CourseRepository struct {
	DB      *sql.DB
	Queries *db.Queries
}

func NewCourseRepository(database *sql.DB) *CourseRepository {
	return &CourseRepository{DB: database, Queries: db.New(database)}
}

func (cr *CourseRepository) Insert(ctx context.Context, course entity.Course) error {
	return cr.Queries.CreateCourse(ctx, db.CreateCourseParams{
		CategoryID: int32(course.CategoryID),
		Name:       course.Name,
	})
}
