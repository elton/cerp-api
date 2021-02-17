package models

// Amount represents each shop sales amount.
type Amount struct {
	ID             int64   `json:"id" gorm:"primaryKey"`
	Period         string  `json:"period,omitempty" gorm:"size:256;uniqueIndex:idx_pershop"`
	ShopCode       string  `json:"shop_code" gorm:"size:256;uniqueIndex:idx_pershop"`
	ShopName       string  `json:"shop_name" gorm:"size:256"`
	OrderNum       uint    `json:"order_num"`
	OrderAmount    float64 `json:"order_amount"`
	OrderAvgAmount float64 `json:"order_avg_amount"`
}
