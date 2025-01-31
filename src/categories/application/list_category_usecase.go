package application

import (
    "holamundo/src/categories/domain/entities"
    "holamundo/src/categories/domain"
)

type ListCategoryUseCase struct {
    repo domain.CategoryRepository
}

func NewListCategoryUseCase(repo domain.CategoryRepository) *ListCategoryUseCase {
    return &ListCategoryUseCase{repo: repo}
}

func (uc *ListCategoryUseCase) Execute() ([]entities.Category, error) {
    return uc.repo.GetAll()
}