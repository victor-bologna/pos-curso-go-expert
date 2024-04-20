package usecase

import (
	"context"
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"github.com/victor-bologna/pos-curso-go-expert-uow/internal/repository"
)

func TestAddCourse(t *testing.T) {
	database, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	assert.NoError(t, err)

	database.Exec("DROP TABLE if exists `courses`;")
	database.Exec("DROP TABLE if exists `categories`;")

	database.Exec("CREATE TABLE IF NOT EXISTS `categories` (id int PRIMARY KEY AUTO_INCREMENT, name varchar(255) NOT NULL);")
	database.Exec("CREATE TABLE IF NOT EXISTS `courses` (id int PRIMARY KEY AUTO_INCREMENT, name varchar(255) NOT NULL, category_id INTEGER NOT NULL, FOREIGN KEY (category_id) REFERENCES categories(id));")

	input := InputCourseUseCase{
		CourseName:       "Java",
		CategoryName:     "Backend",
		CourseCategoryID: 2,
	}

	ctx := context.Background()

	useCase := NewAddCourseUseCase(repository.NewCourseRepository(database), repository.NewCategoryRepository(database))

	err = useCase.Execute(ctx, input)
	assert.NoError(t, err)
}
