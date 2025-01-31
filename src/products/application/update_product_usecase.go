package application

import (
	"holamundo/src/products/domain"       
	"holamundo/src/products/domain/entities" 
)

type UpdateProductUseCase struct {
	repo domain.ProductRepository 
}

func NewUpdateProductUseCase(repo domain.ProductRepository) *UpdateProductUseCase {
	return &UpdateProductUseCase{repo: repo}
}

func (uc *UpdateProductUseCase) Execute(product *entities.Product) error {
	return uc.repo.Update(product) 
}
