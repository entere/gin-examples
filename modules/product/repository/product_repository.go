package repository

import (
	"micro-gin/model"
	"micro-gin/modules/product"
)

type productRepository struct {
}

// NewProductRepository ...
func NewProductRepository() product.Repository {
	return &productRepository{}
}

func (*productRepository) GetByID(id string) (*model.Product, error) {
	var product model.Product
	db := model.DB.Self.Raw("select * from products where id =?", id).Scan(&product)
	if db.Error != nil {
		// log.Println(db.Error)
		return nil, db.Error
	}
	return &product, db.Error
}
