package application

import (
    "holamundo/src/categories/domain/entities"
    "holamundo/src/categories/domain"
)

type CreateCategoryUseCase struct {
    repo domain.CategoryRepository
}

func NewCreateCategoryUseCase(repo domain.CategoryRepository) *CreateCategoryUseCase {
    return &CreateCategoryUseCase{repo: repo}
}

func (uc *CreateCategoryUseCase) Execute(category *entities.Category) error {
    return uc.repo.Save(category)
}