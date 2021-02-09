package models

import (
	"time"

	"github.com/go-acme/lego/v3/log"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// A Shop struct to map every shop information.
type Shop struct {
	ID         int64  `gorm:"unique"`
	ShopID     int    `gorm:"primaryKey"`
	Nick       string `gorm:"size:256"`
	Code       string `gorm:"unique"`
	Name       string `gorm:"size:256"`
	Note       string `gorm:"size:256"`
	TypeName   string `gorm:"size:256;index"`
	CreateDate time.Time
	ModifyDate time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time      `gorm:"index"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
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
	log.Infof("Shop Last Updated: %v\n", lastUpdateAt)
	return lastUpdateAt, nil
}
