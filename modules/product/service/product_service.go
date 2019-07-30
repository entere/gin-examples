package service

import (
	"micro-gin/model"
	"micro-gin/modules/product"
)

type productService struct {
	productRepo product.Repository
}

// NewProductService ... 需实现接口 product.Service
func NewProductService(pr product.Repository) product.Service {
	return &productService{
		productRepo: pr,
	}
}

// GetByID ... 实现product.Service接口方法
func (p *productService) GetByID(id string) (*model.Product, error) {
	product, err := p.productRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

// Get ...
// func (p *ProductService) GetByID() (*model.Product, error) {
// 	product, err := repositorie.UserRepository.Get(u.ID)
// 	// user, err := model.GetUser(u.ID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return user, nil
// }

// // Create ...
// func (p *productService) Create() error {
// 	user := map[string]interface{}{
// 		"id":       u.ID,
// 		"nickname": u.Nickname,
// 		"realname": u.Realname,
// 		"gender":   u.Gender,
// 		"age":      u.Age,
// 		"avatar":   u.Avatar,
// 		"address":  u.Address,
// 	}
// 	if err := repositorie.UserRepository.Create(user); err != nil {
// 		return err
// 	}
// 	return nil
// }
