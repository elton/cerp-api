package models

import (
	"time"

	"gorm.io/gorm"
)

// A Order struct to map the Entity Order
type Order struct {
	gorm.Model
	Code                 string    `json:"code" gorm:"unique" gorm:"type:varchar(255)"`
	PlatformCode         string    `json:"platform_code" gorm:"type:varchar(255)" gorm:"index"`
	OrderTypeName        string    `json:"order_type_name" gorm:"type:varchar(255)"`
	ShopName             string    `json:"shop_name" gorm:"type:varchar(255)"`
	ShopCode             string    `json:"shop_code" gorm:"type:varchar(255)" gorm:"index"`
	VIPName              string    `json:"vip_name" gorm:"column:vip_name"`
	VIPCode              string    `json:"vip_code" gorm:"column:vip_code" gorm:"index"`
	VIPRealName          string    `json:"vipRealName" gorm:"column:vip_real_name"`
	BusinessMan          string    `json:"business_man" gorm:"type:varchar(255)"`
	Qty                  int       `json:"qty" gorm:"type:smallint"`
	Amount               float64   `json:"amount"`
	Payment              float64   `json:"payment"`
	WarehouseName        string    `json:"warehouse_name" gorm:"type:varchar(255)"`
	WarehouseCode        string    `json:"warehouse_code" gorm:"type:varchar(255)"`
	DeliveryState        int       `json:"delivery_state" gorm:"type:tinyint" gorm:"index"`
	ExpressName          string    `json:"express_name" gorm:"type:varchar(255)"`
	ExpressCode          string    `json:"express_code" gorm:"index" gorm:"type:varchar(255)"`
	ReceiverArea         string    `json:"receiver_area" gorm:"type:varchar(255)"`
	PlatformTradingState string    `json:"platform_trading_state" gorm:"type:varchar(255)"`
	PayTime              time.Time `json:"paytime" gorm:"type:datetime"`
	DealTime             time.Time `json:"dealtime" gorm:"type:datetime"`
	CreateTime           time.Time `json:"createtime" gorm:"type:datetime"`
	ModifyTime           time.Time `json:"modifytime" gorm:"type:datetime"`
}

// SaveAll stores all specified the orders in the database.
func (o *Order) SaveAll(orders *[]Order) (*[]Order, error) {
	if err := DB.Create(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}
