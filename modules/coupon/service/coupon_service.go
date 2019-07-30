package service

import (
	"micro-gin/model"
	"micro-gin/modules/coupon"
)

// couponService ...
type couponService struct {
	couponRepo coupon.Repository
}

// NewCouponService ... 需实现接口 coupon.Service
func NewCouponService(r coupon.Repository) coupon.Service {
	return &couponService{
		couponRepo: r,
	}
}

// GetByID ... 实现coupon.Service接口方法
func (c *couponService) GetByID(id string) (*model.Coupon, error) {
	coupon, err := c.couponRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return coupon, nil
}
