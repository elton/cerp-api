package models

import (
	"time"

	"github.com/go-acme/lego/v3/log"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// A Order struct to map the Entity Order
type Order struct {
	ID                   int64  `gorm:"unique"`
	Code                 string `gorm:"primaryKey;size:256"`
	PlatformCode         string `gorm:"size:256;index"`
	OrderTypeName        string `gorm:"size:256"`
	ShopName             string `gorm:"size:256"`
	ShopCode             string `gorm:"size:256;index"`
	VIPName              string `gorm:"size:256;column:vip_name"`
	VIPCode              string `gorm:"size:256;column:vip_code;index"`
	VIPRealName          string `gorm:"size:256;column:vip_real_name"`
	BusinessMan          string `gorm:"size:256"`
	Qty                  int8
	Amount               float64
	Payment              float64
	WarehouseName        string     `gorm:"size:256"`
	WarehouseCode        string     `gorm:"size:256"`
	DeliveryState        int8       `gorm:"index"`
	ExpressName          string     `gorm:"size:256"`
	ExpressCode          string     `gorm:"size:256;index"`
	ReceiverArea         string     `gorm:"size:256"`
	PlatformTradingState string     `gorm:"size:256"`
	Deliveries           []Delivery `gorm:"foreignKey:OrderCode;references:Code"`
	Details              []Detail   `gorm:"foreignKey:OrderCode;references:Code"`
	Payments             []Payment  `gorm:"foreignKey:OrderCode;references:Code"`
	PayTime              time.Time
	DealTime             time.Time
	CreateTime           time.Time
	ModifyTime           time.Time
	CreatedAt            time.Time
	UpdatedAt            time.Time      `gorm:"index"`
	DeletedAt            gorm.DeletedAt `gorm:"index"`
}

// Delivery struct to map the Entity of the Delivery.
type Delivery struct {
	ID            int64  `gorm:"primaryKey"`
	Delivery      bool   `gorm:"comment:发货状态"`
	Code          string `gorm:"size:256;comment:发货单据号"`
	WarehouseName string `gorm:"size:256"`
	WarehouseCode string `gorm:"size:256;index"`
	ExpressName   string `gorm:"size:256"`
	ExpressCode   string `gorm:"size:256"`
	MailNo        string `gorm:"size:256"`
	OrderCode     string `gorm:"size:256;index"`
	CreatedAt     time.Time
	UpdatedAt     time.Time      `gorm:"index"`
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

// Detail struct to map the Entity of the item details.
type Detail struct {
	ID               int64  `gorm:"primaryKey"`
	OID              string `gorm:"size:256;comment:子订单号"`
	Qty              float64
	Price            float64 `gorm:"comment:实际单价"`
	Amount           float64 `gorm:"comment:实际金额"`
	Refund           int     `gorm:"comment:退款状态,0:未退款,1:退款成功,2:退款中"`
	Note             string
	PlatformItemName string  `gorm:"size:256;comment:平台规格名称"`
	PlatformSkuName  string  `gorm:"size:256;comment:平台规格代码"`
	ItemCode         string  `gorm:"size:256;index;comment:商品代码"`
	ItemName         string  `gorm:"size:256;comment:商品名称"`
	ItemSimpleName   string  `gorm:"size:256;comment:商品简称"`
	PostFee          float64 `gorm:"comment:物流费用"`
	DiscountFee      float64 `gorm:"comment:让利金额"`
	AmountAfter      float64 `gorm:"comment:让利后金额"`
	OrderCode        string  `gorm:"size:256;index"`
	CreatedAt        time.Time
	UpdatedAt        time.Time      `gorm:"index"`
	DeletedAt        gorm.DeletedAt `gorm:"index"`
}

// Payment struct to map the Entity of the payment.
type Payment struct {
	ID          int64     `gorm:"primaryKey"`
	Payment     float64   `gorm:"comment:支付金额"`
	PayCode     string    `gorm:"size:256;comment:支付交易号"`
	PayTypeName string    `gorm:"size:256;comment:支付方式名称"`
	PayTime     time.Time `gorm:"comment:支付时间"`
	OrderCode   string    `gorm:"size:256;index"`
	CreatedAt   time.Time
	UpdatedAt   time.Time      `gorm:"index"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

// SaveAll stores all specified the orders in the database.
func (o *Order) SaveAll(orders *[]Order) (*[]Order, error) {
	// 在冲突时，更新除主键以外的所有列到新值。
	if err := DB.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&orders).Error; err != nil {
		return nil, err
	}

	return orders, nil
}

// GetLastUpdatedAt get the last updated timestamp of the order.
func (o *Order) GetLastUpdatedAt() (time.Time, error) {
	var lastUpdateAt time.Time
	var layout string = "2006-01-02 15:04:05"
	if err := DB.Raw("SELECT orders.updated_at FROM orders ORDER BY orders.updated_at DESC LIMIT 1").Scan(&lastUpdateAt).Error; err != nil {
		rtime, err := time.Parse(layout, "0000-00-00 00:00:00")
		return rtime, err
	}
	log.Infof("Order Last Updated: %v\n", lastUpdateAt)
	return lastUpdateAt, nil
}
