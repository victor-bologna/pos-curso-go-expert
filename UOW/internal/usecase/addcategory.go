package usecase

import (
	"context"

	"github.com/victor-bologna/pos-curso-go-expert-uow/internal/entity"
	"github.com/victor-bologna/pos-curso-go-expert-uow/internal/repository"
)

type InputCategoryUseCase struct {
	Name string
}

type AddCategoryUseCase struct {
	CategoryRepository repository.CategoryRepositoryInterface
}

func NewCategoryUseCase(categoryRepository repository.CategoryRepositoryInterface) *AddCategoryUseCase {
	return &AddCategoryUseCase{CategoryRepository: categoryRepository}
}

func (a *AddCategoryUseCase) Execute(ctx context.Context, input InputCategoryUseCase) error {
	category := entity.Category{
		Name: input.Name,
	}

	err := a.CategoryRepository.Insert(ctx, category)
	if err != nil {
		return err
	}

	return nil
}
