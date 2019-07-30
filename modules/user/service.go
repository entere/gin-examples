package user

import "micro-gin/model"

// Service ...
type Service interface {
	GetByID(id string) (*model.User, error)
	Store(um *model.User) error
	Signup(um *model.User, uam *model.UserAuth) error
	ExistByID(id string) (bool, error)
}
