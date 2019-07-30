package model

import (
	// _ mysql
	_ "github.com/go-sql-driver/mysql"
)

// User ...
type User struct {
	BaseModel
	// UserAuth
	ID       string `json:"id"`
	Nickname string `json:"nickname"`
	Realname string `json:"realname"`
	Gender   int    `json:"gender"`
	Age      int    `json:"age"`
	Avatar   string `json:"avatar"`
	Address  string `json:"address"`
	// UserAuths []UserAuth

	// Identifier string `json:"identifier"`
}

// TableName ...
func (user *User) TableName() string {
	return "users"
}

// // GetUser 根据 userID获取单个用户
// func GetUser(userID string) (*User, error) {
// 	var (
// 		user User
// 	)
// 	db := DB.Self.Raw("select a.*,b.identity_type,b.identifier from users a inner join user_auths b  on a.id=b.user_id where a.id =?", userID).Scan(&user)
// 	if db.Error != nil {
// 		log.Println(db.Error)
// 		return nil, db.Error
// 	}

// 	return &user, db.Error
// }

// // GetUsers 获取一批用户
// func GetUsers(pageNum int, pageSize int, maps interface{}) ([]*User, error) {
// 	var users []*User
// 	err := DB.Self.Where(maps).Offset(pageNum).Limit(pageSize).Find(&users).Error
// 	if err != nil && err != gorm.ErrRecordNotFound {
// 		return nil, err
// 	}
// 	return users, nil

// }

// // ExistIdentifierByUserID ...判断用户ID是否存在
// func ExistIdentifierByUserID(userID string) (bool, error) {
// 	var user User
// 	err := DB.Self.Select("id").Where("id = ?", userID).First(&user).Error
// 	if err != nil && err != gorm.ErrRecordNotFound {
// 		return false, err
// 	}

// 	if user.ID != "" {
// 		return true, nil
// 	}
// 	return false, nil

// }

// // AddUser 注册时向usess表添加注册信息
// func AddUser(data map[string]interface{}) error {

// 	user := User{
// 		ID:       data["id"].(string),
// 		Nickname: data["nickname"].(string),
// 		Realname: data["realname"].(string),
// 		Address:  data["address"].(string),
// 		Gender:   data["gender"].(int),
// 		Age:      data["age"].(int),
// 		Avatar:   data["avatar"].(string),
// 	}
// 	if err := DB.Self.Create(&user).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }

// // AddUserAndUserAuth ...
// func AddUserAndUserAuth(data map[string]interface{}) error {
// 	tx := DB.Self.Begin()
// 	user := User{
// 		ID:       data["id"].(string),
// 		Nickname: data["nickname"].(string),
// 		Realname: data["realname"].(string),
// 		Address:  data["address"].(string),
// 		Gender:   data["gender"].(int),
// 		Age:      data["age"].(int),
// 		Avatar:   data["avatar"].(string),
// 	}
// 	if err := DB.Self.Create(&user).Error; err != nil {
// 		tx.Rollback()
// 		return err
// 	}

// 	userAuth := UserAuth{
// 		Identifier:   data["identifier"].(string),
// 		Credential:   data["credential"].(string),
// 		IdentityType: data["identity_type"].(string),
// 		UserID:       data["user_id"].(string),
// 	}
// 	if err := DB.Self.Create(&userAuth).Error; err != nil {
// 		tx.Rollback()
// 		return err
// 	}
// 	tx.Commit()
// 	return nil

// }

// // Validate ...
// func (user *User) Validate() error {
// 	validate := validator.New()
// 	return validate.Struct(user)
// }
