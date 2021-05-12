package interfaces

import (
	model "github.com/hyperyuri/repository-pattern-go/models/user"
)

type IUserRepository interface {
	Store(model.User) (model.User, error)
	Delete(id uint) (model.User, error)
	Update(model.User) error
	Get(id uint) (model.User, error)
	GetAll() ([]model.User, error)
}
