package usecase

import (
	"context"

	"github.com/victor-bologna/pos-curso-go-expert-uow/internal/entity"
	"github.com/victor-bologna/pos-curso-go-expert-uow/internal/repository"
)

type InputCourseUseCase struct {
	CourseName       string
	CategoryName     string
	CourseCategoryID int
}

type AddCourseUseCase struct {
	CourseRepository    repository.CourseRepositoryInterface
	CategoryRespository repository.CategoryRepositoryInterface
}

func NewAddCourseUseCase(courseRepository repository.CourseRepositoryInterface,
	categoryRepository repository.CategoryRepositoryInterface) *AddCourseUseCase {
	return &AddCourseUseCase{
		CourseRepository:    courseRepository,
		CategoryRespository: categoryRepository,
	}
}

func (a *AddCourseUseCase) Execute(ctx context.Context, input InputCourseUseCase) error {
	category := entity.Category{
		Name: input.CategoryName,
	}

	err := a.CategoryRespository.Insert(ctx, category)
	if err != nil {
		return nil
	}

	course := entity.Course{
		Name:       input.CourseName,
		CategoryID: input.CourseCategoryID,
	}

	err = a.CourseRepository.Insert(ctx, course)
	if err != nil {
		return nil
	}

	return nil
}
