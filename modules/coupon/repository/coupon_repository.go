package repository

import (
	"micro-gin/model"
	"micro-gin/modules/coupon"
)

type couponRepository struct {
}

// NewCouponRepository ...
func NewCouponRepository() coupon.Repository {
	return &couponRepository{}
}

func (*couponRepository) GetByID(id string) (*model.Coupon, error) {
	var coupon model.Coupon
	db := model.DB.Self.Raw("select * from coupons where id =?", id).Scan(&coupon)
	if db.Error != nil {
		// log.Println(db.Error)
		return nil, db.Error
	}
	return &coupon, db.Error
}
