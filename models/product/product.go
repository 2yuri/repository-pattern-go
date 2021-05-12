package product

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name"`
	Price     float32        `json:"price"`
	CreatedAt time.Time      `json:"created"`
	UpdatedAt time.Time      `json:"updated"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted"`
}

func NewProduct(name string, price float32) *Product {
	return &Product{
		Name:  name,
		Price: price,
	}
}

func NewProductWithID(id uint, name string, price float32) *Product {
	return &Product{
		ID:    id,
		Name:  name,
		Price: price,
	}
}
