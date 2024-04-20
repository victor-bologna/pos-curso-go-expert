package usecase

import (
	"context"

	"github.com/victor-bologna/pos-curso-go-expert-uow/internal/entity"
	"github.com/victor-bologna/pos-curso-go-expert-uow/internal/repository"
	"github.com/victor-bologna/pos-curso-go-expert-uow/pkg/uow"
)

type AddCourseUseCaseUow struct {
	Uow uow.UowInterface
}

func NewAddCourseUseCaseUow(uow uow.UowInterface) *AddCourseUseCaseUow {
	return &AddCourseUseCaseUow{Uow: uow}
}

func (a *AddCourseUseCaseUow) Execute(ctx context.Context, input InputCourseUseCase) error {
	return a.Uow.Do(ctx, func(uow *uow.Uow) error {
		category := entity.Category{
			Name: input.CategoryName,
		}

		categoryRepo := a.getCategoryRepository(ctx)
		err := categoryRepo.Insert(ctx, category)
		if err != nil {
			return nil
		}

		course := entity.Course{
			Name:       input.CourseName,
			CategoryID: input.CourseCategoryID,
		}

		courseRepo := a.getCourseRepository(ctx)
		err = courseRepo.Insert(ctx, course)
		if err != nil {
			return err
		}
		return nil
	})

	// return nil
}

func (a *AddCourseUseCaseUow) getCategoryRepository(ctx context.Context) repository.CategoryRepositoryInterface {
	repo, err := a.Uow.GetRepository(ctx, "CategoryRepository")
	if err != nil {
		panic(err)
	}
	return repo.(repository.CategoryRepositoryInterface) //cast
}

func (a *AddCourseUseCaseUow) getCourseRepository(ctx context.Context) repository.CourseRepositoryInterface {
	repo, err := a.Uow.GetRepository(ctx, "CourseRepository")
	if err != nil {
		panic(err)
	}
	return repo.(repository.CourseRepositoryInterface) //cast
}
