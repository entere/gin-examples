package model

// UserAuth ...
type UserAuth struct {
	BaseModel
	ID           string `json:"id"`
	Identifier   string `json:"identifier"`
	Credential   string `json:"credential"`
	IdentityType string `json:"identity_type"`
	UserID       string `json:"user_id"`
}

// TableName ...
func (u *UserAuth) TableName() string {
	return "user_auths"
}

// // AddUserAuth 注册时向user_auths表添加注册信息
// func AddUserAuth(data map[string]interface{}) error {
// 	userAuth := UserAuth{
// 		ID:           data["id"].(string),
// 		Identifier:   data["identifier"].(string),
// 		Credential:   data["credential"].(string),
// 		IdentityType: data["identity_type"].(string),
// 		UserID:       data["user_id"].(string),
// 	}
// 	if err := DB.Self.Create(&userAuth).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }

// // ExistIdentifierByType ...判断注册名是否存在
// func ExistIdentifierByType(identifier string, identityType string) (bool, error) {
// 	var userAuth UserAuth
// 	err := DB.Self.Select("user_id").Where("identifier = ?", identifier).Where("identity_type = ?", identityType).First(&userAuth).Error
// 	if err != nil && err != gorm.ErrRecordNotFound {
// 		return false, err
// 	}
// 	if userAuth.UserID != "" {
// 		return true, nil
// 	}

// 	return false, nil

// }

// // GetUserAuth ...
// func GetUserAuth(identifier string, identityType string) (*UserAuth, error) {
// 	var userAuth UserAuth
// 	err := DB.Self.Where("identifier = ?", identifier).Where("identity_type = ?", identityType).First(&userAuth).Error
// 	if err != nil && err != gorm.ErrRecordNotFound {
// 		return nil, err
// 	}

// 	return &userAuth, nil

// }
