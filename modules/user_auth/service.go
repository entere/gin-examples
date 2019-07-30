package user_auth

import "micro-gin/model"

// Service ...
type Service interface {
	GetByID(id string) (*model.UserAuth, error)
	Store(*model.UserAuth) error
	ExistByID(id string) (bool, error)
	ExistByIdentifierAndIdentityType(identifier string, identityType string) (bool, error)
	GetByIdentifierAndIdentityType(identifier string, identityType string) (*model.UserAuth, error)
}
