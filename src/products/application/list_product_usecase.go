package application

import (
	"holamundo/src/products/domain" 
	"holamundo/src/products/domain/entities" 
)

type ListProductUseCase struct {
	repo domain.ProductRepository 
}

func NewListProductUseCase(repo domain.ProductRepository) *ListProductUseCase {
	return &ListProductUseCase{repo: repo}
}

func (uc *ListProductUseCase) Execute() ([]entities.Product, error) {
	return uc.repo.GetAll() 
}
