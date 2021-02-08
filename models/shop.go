package models

import (
	"time"

	"gorm.io/gorm"
)

// A Shop struct to map every shop information.
type Shop struct {
	gorm.Model
	ShopID     string    `json:"shop_id" gorm:"type:varchar(255);index"`
	Nick       string    `json:"nick" gorm:"type:varchar(255)"`
	Code       string    `json:"code" gorm:"unique" gorm:"type:varchar(255);index"`
	Name       string    `json:"name" gorm:"type:varchar(255)"`
	Note       string    `json:"note"`
	TypeName   string    `json:"type_name" gorm:"type:varchar(255);index"`
	CreateDate time.Time `json:"create_date"`
	ModifyDate time.Time `json:"modify_date"`
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

	if err := DB.Create(&shops).Error; err != nil {
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
