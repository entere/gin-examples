package repository

import (
	"micro-gin/model"
	"micro-gin/modules/user_auth"

	"github.com/jinzhu/gorm"
)

type userAuthRepository struct {
}

// NewUserAuthRepository ...
func NewUserAuthRepository() user_auth.Repository {
	return &userAuthRepository{}
}

func (ur *userAuthRepository) GetByID(id string) (*model.UserAuth, error) {
	var userAuth model.UserAuth
	db := model.DB.Self.Raw("select a.*,b.identity_type,b.identifier from users a inner join user_auths b  on a.id=b.user_id where a.id =?", id).Scan(&userAuth)
	if db.Error != nil {
		return nil, db.Error
	}
	return &userAuth, nil
}

func (ur *userAuthRepository) Store(userAuth *model.UserAuth) error {
	if err := model.DB.Self.Create(&userAuth).Error; err != nil {
		return err
	}
	return nil
}

func (ur *userAuthRepository) ExistByID(id string) (bool, error) {
	var userAuth model.UserAuth
	err := model.DB.Self.Select("id").Where("id = ?", id).First(&userAuth).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if userAuth.ID != "" {
		return true, nil
	}
	return false, nil
}

func (ur *userAuthRepository) ExistByIdentifierAndIdentityType(identifier string, identityType string) (bool, error) {
	var userAuth model.UserAuth
	err := model.DB.Self.Select("user_id").Where("identifier = ?", identifier).Where("identity_type = ?", identityType).First(&userAuth).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if userAuth.UserID != "" {
		return true, nil
	}
	return false, nil
}

func (ur *userAuthRepository) GetByIdentifierAndIdentityType(identifier string, identityType string) (*model.UserAuth, error) {
	var userAuth model.UserAuth
	err := model.DB.Self.Select("*").Where("identifier = ?", identifier).Where("identity_type = ?", identityType).First(&userAuth).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &userAuth, nil
}
