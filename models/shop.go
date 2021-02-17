package models

import (
	"time"

	"github.com/bwmarrin/snowflake"
	"github.com/elton/cerp-api/utils/logger"
	"github.com/golang-module/carbon"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// A Shop struct to map every shop information.
type Shop struct {
	ID         int64          `json:"id" gorm:"unique"`
	ShopID     int            `json:"shop_id" gorm:"primaryKey"`
	Nick       string         `json:"nick" gorm:"size:256"`
	Code       string         `json:"code" gorm:"unique"`
	Name       string         `json:"name" gorm:"size:256"`
	Note       string         `json:"note" gorm:"size:256"`
	TypeName   string         `json:"type_name" gorm:"size:256;index"`
	CreateDate time.Time      `json:"create_date"`
	ModifyDate time.Time      `json:"modify_date"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at" gorm:"index"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

var node *snowflake.Node

func init() {
	// Create a new Node with a Node number of 1
	node, _ = snowflake.NewNode(1)
}

// Save create a new Shop
func (s *Shop) Save() (*Shop, error) {
	if err := DB.Create(&s).Error; err != nil {
		return nil, err
	}
	return s, nil
}

// SaveAll save all shop to database.
func (s *Shop) SaveAll(shops *[]Shop) (*[]Shop, error) {

	// 在冲突时，更新除主键以外的所有列到新值。
	if err := DB.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&shops).Error; err != nil {
		return nil, err
	}
	return shops, nil
}

// GetAllShops returns all shop from database.
func (s *Shop) GetAllShops() (*[]Shop, error) {
	shops := []Shop{}
	if err := DB.Find(&shops).Error; err != nil {
		return nil, err
	}
	return &shops, nil
}

// GetLastUpdatedAt get the datetime of last updated of shop.
func (s *Shop) GetLastUpdatedAt() (time.Time, error) {
	var lastUpdateAt time.Time
	var layout string = "2006-01-02 15:04:05"
	if err := DB.Raw("SELECT shops.updated_at FROM shops ORDER BY shops.updated_at DESC LIMIT 1").Scan(&lastUpdateAt).Error; err != nil {
		rtime, err := time.Parse(layout, "0000-00-00 00:00:00")
		return rtime, err
	}
	logger.Info.Printf("Shop Last Updated: %v\n", lastUpdateAt)
	return lastUpdateAt, nil
}

// GetAmountByShop returns sales amount for each shop.
func (s *Shop) GetAmountByShop(start, end string) ([]Amount, error) {
	var amounts []Amount

	if err := DB.Raw("SELECT orders.shop_code AS shop_code,	orders.shop_name AS shop_name, COUNT( orders.id ) AS order_num,	SUM(orders.payment) AS order_amount, SUM(orders.payment)/COUNT( orders.id ) AS order_avg_amount FROM orders WHERE orders.order_type_name = '销售订单' AND orders.paytime >= ? AND orders.paytime < ? GROUP BY	orders.shop_code,	orders.shop_name ORDER BY	orders.shop_code", start, end).Scan(&amounts).Error; err != nil {
		return []Amount{}, err
	}

	for _, amount := range amounts {
		amount.ID = node.Generate().Int64()
		amount.Period = carbon.Parse(start).ToFormatString("Y-m")

		// 在冲突时，更新除主键以外的所有列到新值。
		if err := DB.Clauses(clause.OnConflict{
			UpdateAll: true,
		}).Create(&amount).Error; err != nil {
			return nil, err
		}
	}
	return amounts, nil
}
