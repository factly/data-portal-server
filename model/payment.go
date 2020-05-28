package model

// Payment model
type Payment struct {
	Base
	Amount     int      `gorm:"column:amount" json:"amount" validate:"required"`
	Gateway    string   `gorm:"column:gateway" json:"gateway" validate:"required"`
	CurrencyID uint     `gorm:"column:currency_id" json:"currency_id" validate:"required"`
	Currency   Currency `gorm:"foreignkey:currency_id;association_foreignkey:id"  json:"currency"`
	Status     string   `gorm:"column:status" json:"status" validate:"required"`
}
