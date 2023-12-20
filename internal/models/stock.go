package models

// Model for logging all stock decreases for the sake of idempotence
type StockDecreaseLog struct {
	Id           int64 `json:"id" gorm:"primaryKey"`
	OrderId      int64 `json:"order_id"`
	ProductRefer int64 `json:"product_id"`
}
