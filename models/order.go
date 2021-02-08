package models

import (
	"time"

	"gorm.io/gorm"
)

// A Order struct to map the Entity Order
type Order struct {
	gorm.Model
	Code                 string     `json:"code" gorm:"unique" gorm:"type:varchar(255)"`
	PlatformCode         string     `json:"platform_code" gorm:"type:varchar(255);index"`
	OrderTypeName        string     `json:"order_type_name" gorm:"type:varchar(255)"`
	ShopName             string     `json:"shop_name" gorm:"type:varchar(255)"`
	ShopCode             string     `json:"shop_code" gorm:"type:varchar(255);index"`
	VIPName              string     `json:"vip_name" gorm:"column:vip_name"`
	VIPCode              string     `json:"vip_code" gorm:"column:vip_code;index"`
	VIPRealName          string     `json:"vipRealName" gorm:"column:vip_real_name"`
	BusinessMan          string     `json:"business_man" gorm:"type:varchar(255)"`
	Qty                  int        `json:"qty" gorm:"type:smallint"`
	Amount               float64    `json:"amount"`
	Payment              float64    `json:"payment"`
	WarehouseName        string     `json:"warehouse_name" gorm:"type:varchar(255)"`
	WarehouseCode        string     `json:"warehouse_code" gorm:"type:varchar(255)"`
	DeliveryState        int        `json:"delivery_state" gorm:"type:tinyint;index"`
	ExpressName          string     `json:"express_name" gorm:"type:varchar(255)"`
	ExpressCode          string     `json:"express_code" gorm:"type:varchar(255);index"`
	ReceiverArea         string     `json:"receiver_area" gorm:"type:varchar(255)"`
	PlatformTradingState string     `json:"platform_trading_state" gorm:"type:varchar(255)"`
	Deliveries           []Delivery `json:"delivery"`
	Details              []Detail   `json:"details"`
	Payments             []Payment  `json:"payments"`
	PayTime              time.Time  `json:"paytime" gorm:"type:datetime"`
	DealTime             time.Time  `json:"dealtime" gorm:"type:datetime"`
	CreateTime           time.Time  `json:"createtime" gorm:"type:datetime"`
	ModifyTime           time.Time  `json:"modifytime" gorm:"type:datetime"`
}

// Delivery struct to map the Entity of the Delivery.
type Delivery struct {
	gorm.Model
	Delivery      bool   `json:"delivery"`
	Code          string `json:"code" gorm:"index"`
	WarehouseName string `json:"warehouse_name"`
	WarehouseCode string `json:"warehouse_code" gorm:"index"`
	ExpressName   string `json:"express_name"`
	ExpressCode   string `json:"express_code" gorm:"index"`
	MailNo        string `json:"mail_no"`
	OrderID       uint   `json:"order_id"`
}

// Detail struct to map the Entity of the item details.
type Detail struct {
	gorm.Model
	OID              string  `json:"oid" gorm:"comment:子订单号;index"`
	Qty              float64 `json:"qty"`
	Price            float64 `json:"price" gorm:"comment:实际单价"`
	Amount           float64 `json:"amount" gorm:"comment:实际金额"`
	Refund           int     `json:"refund" gorm:"comment:退款状态,0:未退款,1:退款成功,2:退款中"`
	Note             string  `json:"note"`
	PlatformItemName string  `json:"platform_item_name" gorm:"comment:平台规格名称"`
	PlatformSkuName  string  `json:"platform_sku_name" gorm:"comment:平台规格代码"`
	ItemCode         string  `json:"item_code" gorm:"index"`
	ItemName         string  `json:"item_name"`
	ItemSimpleName   string  `json:"item_simple_name"`
	PostFee          float64 `json:"post_fee" gorm:"comment:物流费用"`
	DiscountFee      float64 `json:"discount_fee" gorm:"comment:让利金额"`
	AmountAfter      float64 `json:"amount_after" gorm:"comment:让利后金额"`
	OrderID          string  `json:"order_id"`
}

// Payment struct to map the Entity of the payment.
type Payment struct {
	gorm.Model
	Payment     float64   `json:"payment"`
	PayCode     string    `json:"payCode"`
	PayTypeName string    `json:"pay_type_name"`
	PayTime     time.Time `json:"payTime"`
	OrderID     string    `json:"order_id"`
}

// SaveAll stores all specified the orders in the database.
func (o *Order) SaveAll(orders *[]Order) (*[]Order, error) {
	if err := DB.Create(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}
