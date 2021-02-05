package models

import "gorm.io/gorm"

// A Shop struct to map every shop information.
type Shop struct {
	gorm.Model
	ShopID   string `json:"shop_id" gorm:"type:varchar(255)"`
	Nick     string `json:"nick" gorm:"type:varchar(255)"`
	Code     string `json:"code" gorm:"type:varchar(255)"`
	Name     string `json:"name" gorm:"type:varchar(255)"`
	Note     string `json:"note"`
	TypeName string `json:"type_name" gorm:"type:varchar(255)"`
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
