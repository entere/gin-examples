package service

import (
	"micro-gin/model"
	"micro-gin/modules/user"
)

type userService struct {
	userRepo user.Repository
}

// NewUserService ...
func NewUserService(ur user.Repository) user.Service {
	return &userService{
		userRepo: ur,
	}
}

func (us *userService) GetByID(id string) (*model.User, error) {
	user, err := us.userRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (us *userService) Store(user *model.User) error {

	if err := us.userRepo.Store(user); err != nil {
		return err
	}
	return nil
}

func (us *userService) Signup(user *model.User, userAuth *model.UserAuth) error {

	if err := us.userRepo.Signup(user, userAuth); err != nil {
		return err
	}
	return nil
}

func (us *userService) ExistByID(id string) (bool, error) {
	return us.userRepo.ExistByID(id)
}
