package service

import (
	"micro-gin/model"
	"micro-gin/modules/user_auth"
)

type userAuthService struct {
	userAuthRepo user_auth.Repository
}

// NewUserAuthService ...
func NewUserAuthService(ur user_auth.Repository) user_auth.Service {
	return &userAuthService{
		userAuthRepo: ur,
	}
}

func (us *userAuthService) GetByID(id string) (*model.UserAuth, error) {
	user, err := us.userAuthRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (us *userAuthService) Store(userAuth *model.UserAuth) error {

	if err := us.userAuthRepo.Store(userAuth); err != nil {
		return err
	}
	return nil
}

func (us *userAuthService) ExistByID(id string) (bool, error) {
	return us.userAuthRepo.ExistByID(id)
}

func (us *userAuthService) ExistByIdentifierAndIdentityType(identifier string, identityType string) (bool, error) {
	return us.userAuthRepo.ExistByIdentifierAndIdentityType(identifier, identityType)
}

func (us *userAuthService) GetByIdentifierAndIdentityType(identifier string, identityType string) (*model.UserAuth, error) {
	return us.userAuthRepo.GetByIdentifierAndIdentityType(identifier, identityType)
}
