package coupon

import "micro-gin/model"

// Service ...
type Service interface {
	GetByID(id string) (*model.Coupon, error)
}
