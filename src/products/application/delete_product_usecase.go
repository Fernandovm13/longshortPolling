package application

import "holamundo/src/products/domain" 

type DeleteProductUseCase struct {
	repo domain.ProductRepository 
}

func NewDeleteProductUseCase(repo domain.ProductRepository) *DeleteProductUseCase {
	return &DeleteProductUseCase{repo: repo}
}

func (uc *DeleteProductUseCase) Execute(id int32) error {
	return uc.repo.Delete(id) 
}
