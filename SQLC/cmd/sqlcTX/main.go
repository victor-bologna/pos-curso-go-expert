package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/victor-bologna/pos-curso-go-expert-sqlc/internal/db"

	_ "github.com/go-sql-driver/mysql"
)

type CourseDB struct {
	dbConn *sql.DB
	*db.Queries
}

func NewCourseDB(dbConn *sql.DB) *CourseDB {
	return &CourseDB{
		dbConn:  dbConn,
		Queries: db.New(dbConn),
	}
}

func (c *CourseDB) callTx(ctx context.Context, fn func(*db.Queries) error) error {
	tx, err := c.dbConn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := db.New(tx) // tx implementa funções da interface DBTX
	err = fn(q)     // fn é uma função definida (no caso create/update/delete/find course) na linha 44
	// que é executada dentro da função callTx
	if err != nil {
		if errRb := tx.Rollback(); errRb == nil {
			return fmt.Errorf("error on rollback: %v, original rollback: %w", errRb, err)
		}
		return err
	}
	return tx.Commit()
}

func (c *CourseDB) CreateCourseAndCategory(ctx context.Context, courseArgs db.CreateCourseParams,
	categoryArgs db.CreateCategoryParams) error {
	err := c.callTx(ctx, func(q *db.Queries) error { //Func para criar Course e Category, caso de algum erro,
		// ele da rollback
		var err error // var err do escopo da função
		err = q.CreateCategory(ctx, db.CreateCategoryParams{
			ID:          categoryArgs.ID,
			Name:        categoryArgs.Name,
			Description: categoryArgs.Description,
		})
		if err != nil {
			return err
		}
		err = q.CreateCourse(ctx, db.CreateCourseParams{
			ID:          courseArgs.ID,
			Name:        courseArgs.Name,
			Description: courseArgs.Description,
			Price:       courseArgs.Price,
			CategoryID:  categoryArgs.ID,
		})
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func main() {
	ctx := context.Background()
	mysql, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	if err != nil {
		panic(err)
	}
	defer mysql.Close()

	queries := db.New(mysql)

	listCourses(ctx, queries)

	//createCourseAndCategory(ctx, err, courseDB)
}

func listCourses(ctx context.Context, queries *db.Queries) {
	courses, err := queries.ListCourses(ctx)
	if err != nil {
		panic(err)
	}

	for _, course := range courses {
		fmt.Println(course)
	}
}

func createCourseAndCategory(ctx context.Context, err error, courseDB *CourseDB) {
	course := db.CreateCourseParams{
		ID:   uuid.New().String(),
		Name: "Java",
		Description: sql.NullString{
			String: "Java Development II err",
			Valid:  true,
		},
		Price: 10.20,
	}

	category := db.CreateCategoryParams{
		ID:   uuid.New().String(),
		Name: "Backend III",
		Description: sql.NullString{
			String: "Backend III development err",
			Valid:  true,
		},
	}

	err = courseDB.CreateCourseAndCategory(ctx, course, category)
	if err != nil {
		panic(err)
	}
}
