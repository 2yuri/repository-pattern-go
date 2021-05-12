package migrations

import (
	"github.com/hyperyuri/repository-pattern-go/models/product"
	"github.com/hyperyuri/repository-pattern-go/models/user"
	"gorm.io/gorm"
)

func RunAutoMigrations(db *gorm.DB) {
	db.AutoMigrate(&product.Product{})
	db.AutoMigrate(&user.User{})
}
