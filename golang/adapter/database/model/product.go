package model

import "gorm.io/gorm"

type Products struct {
	gorm.Model
	Name  string  `gorm:"type:varchar(255);not null" json:"name"`
	Price float32 `gorm:"not null" json:"price"`
	Stock uint    `gorm:"not null" json:"stock"`
}

func (Products) TableName() string {
	return "products"
}

func (m *Products) Create(db *gorm.DB) error {
	return db.Omit("ID").Create(&m).Error
}

func (m *Products) Update(db *gorm.DB) error {
	return db.Select("*").Where("id = ?", m.ID).Updates(&m).Error
}

func (m *Products) Delete(db *gorm.DB) error { //soft
	return db.Table(m.TableName()).Where("id = ?", m.ID).Delete(&m).Error
}

func (m *Products) Remove(db *gorm.DB) error { //permanent
	return db.Unscoped().Delete(&m).Error
}
