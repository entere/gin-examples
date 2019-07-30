package user

import (
	"micro-gin/model"
)

// Repository ...
type Repository interface {
	GetByID(id string) (*model.User, error)
	Store(um *model.User) error
	Signup(um *model.User, uam *model.UserAuth) error
	ExistByID(id string) (bool, error)
}
