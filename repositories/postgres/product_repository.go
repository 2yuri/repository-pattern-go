package postgres

import (
	"fmt"

	"github.com/hyperyuri/repository-pattern-go/database"
	model "github.com/hyperyuri/repository-pattern-go/models/product"
)

type ProductRepository struct {
}

func (s *ProductRepository) Store(p model.Product) (model.Product, error) {
	db := database.GetDatabase()
	err := db.Create(&p).Error

	if err != nil {
		return model.Product{}, err
	}

	return p, nil
}

func (s *ProductRepository) Delete(id uint) (model.Product, error) {
	db := database.GetDatabase()
	n, err := s.Get(id)
	if err != nil {
		return model.Product{}, err
	}

	err = db.Delete(&n).Error
	if err != nil {
		return model.Product{}, fmt.Errorf("cannot delete file: %v", err)
	}

	return n, nil
}

func (s *ProductRepository) Update(p model.Product) error {
	db := database.GetDatabase()

	err := db.Save(&p).Error
	if err != nil {
		return fmt.Errorf("cannot update client on pg: %v", err)
	}

	return nil
}

func (s *ProductRepository) Get(id uint) (model.Product, error) {
	db := database.GetDatabase()
	var p model.Product
	err := db.First(&p, id).Error

	if err != nil {
		return model.Product{}, fmt.Errorf("cannot find product by id: %v", err)
	}

	return p, nil
}

func (s *ProductRepository) GetAll() ([]model.Product, error) {
	db := database.GetDatabase()
	var p []model.Product
	err := db.Find(&p).Error

	if err != nil {
		return nil, fmt.Errorf("cannot find products: %v", err)
	}

	return p, err
}
