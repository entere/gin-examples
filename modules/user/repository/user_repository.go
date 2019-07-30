package repository

import (
	"micro-gin/model"
	"micro-gin/modules/user"

	"github.com/jinzhu/gorm"
)

type userRepository struct {
}

// NewUserRepository ...
func NewUserRepository() user.Repository {
	return &userRepository{}
}

func (ur *userRepository) GetByID(id string) (*model.User, error) {
	var user model.User
	db := model.DB.Self.Raw("select a.*,b.identity_type,b.identifier from users a inner join user_auths b  on a.id=b.user_id where a.id =?", id).Scan(&user)
	if db.Error != nil {
		return nil, db.Error
	}
	return &model.User{}, nil
}

func (ur *userRepository) Store(user *model.User) error {
	if err := model.DB.Self.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

// Signup ... 注册
func (ur *userRepository) Signup(user *model.User, userAuth *model.UserAuth) error {
	tx := model.DB.Self.Begin()
	if err := model.DB.Self.Create(&user).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := model.DB.Self.Create(&userAuth).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (ur *userRepository) ExistByID(id string) (bool, error) {
	var user model.User
	err := model.DB.Self.Select("id").Where("id = ?", id).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if user.ID != "" {
		return true, nil
	}
	return false, nil
}
