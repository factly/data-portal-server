package model

// Order model
type Order struct {
	Base
	UserID          uint      `gorm:"column:user_id" json:"user_id" validate:"required"`
	Status          string    `gorm:"column:status" json:"status" validate:"required"`
	PaymentID       uint      `gorm:"column:payment_id" json:"payment_id" sql:"DEFAULT:NULL"`
	Payment         *Payment  `gorm:"foreignkey:payment_id;association_foreignkey:id" json:"payment"`
	RazorpayOrderID string    `gorm:"column:razorpay_order_id" json:"razorpay_order_id"`
	Products        []Product `gorm:"many2many:order_item;" json:"products"`
}
