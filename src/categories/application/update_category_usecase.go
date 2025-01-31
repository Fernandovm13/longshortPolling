package application

import (
    "holamundo/src/categories/domain/entities"
    "holamundo/src/categories/domain"
)

type UpdateCategoryUseCase struct {
    repo domain.CategoryRepository
}

func NewUpdateCategoryUseCase(repo domain.CategoryRepository) *UpdateCategoryUseCase {
    return &UpdateCategoryUseCase{repo: repo}
}

func (uc *UpdateCategoryUseCase) Execute(category *entities.Category) error {
    return uc.repo.Update(category)
}