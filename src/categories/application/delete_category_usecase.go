package application

import "holamundo/src/categories/domain"

type DeleteCategoryUseCase struct {
	repo domain.CategoryRepository
}

func NewDeleteCategoryUseCase(repo domain.CategoryRepository) *DeleteCategoryUseCase {
	return &DeleteCategoryUseCase{repo: repo}
}

func (uc *DeleteCategoryUseCase) Execute(id int32) error {
	return uc.repo.Delete(id)
}