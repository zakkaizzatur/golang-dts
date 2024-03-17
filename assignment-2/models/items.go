package models

import (
	"time"

	"gorm.io/gorm"
)

type Items struct {
	ID        uint           `gorm:"primaryKey;autoIncrement" json:"id,omitempty"`
	CreatedAt time.Time      `json:"created_at,omitempty"`
	UpdatedAt time.Time      `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	ItemCode    string `gorm:"not null;type:varchar(255)" json:"item_code"`
	Quantity    uint   `gorm:"not null;type:int" json:"quantity"`
	Description string `gorm:"not null;type:varchar(255)" json:"description"`
	OrderID     uint   `gorm:"not null;type:int" json:"-"`
}