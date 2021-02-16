package models

// Amount represents each shop sales amount.
type Amount struct {
	ShopCode       string  `json:"shop_code"`
	ShopName       string  `json:"shop_name"`
	OrderNum       uint    `json:"order_num"`
	OrderAmount    float64 `json:"order_amount"`
	OrderAvgAmount float64 `json:"order_avg_amount"`
}
