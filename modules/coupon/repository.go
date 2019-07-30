package coupon

import (
	"micro-gin/model"
)

// Repository ...
type Repository interface {
	GetByID(id string) (*model.Coupon, error)
}
