package product

import (
	"github.com/hyperyuri/repository-pattern-go/interfaces"
	"github.com/hyperyuri/repository-pattern-go/models/product"
)

type ProductService struct {
	repository interfaces.IProductRepository
}

func (s *ProductService) CreateProduct(data ProductDTO) (product.Product, error) {
	prod := product.NewProduct(data.Name, data.Price)

	return s.repository.Store(*prod)
}

func (s *ProductService) UpdateProduct(data ProductDTO) error {
	prod := product.NewProductWithID(data.Id, data.Name, data.Price)

	err := s.repository.Update(*prod)

	if err != nil {
		return err
	}

	return nil
}

func (s *ProductService) DeleteProduct(id uint) (product.Product, error) {
	return s.repository.Delete(id)
}

func (s *ProductService) ShowProduct(id uint) (product.Product, error) {
	return s.repository.Get(id)
}

func (s *ProductService) ShowAllProducts() ([]product.Product, error) {
	return s.repository.GetAll()
}

func NewProductService(r interfaces.IProductRepository) *ProductService {
	return &ProductService{
		repository: r,
	}
}
