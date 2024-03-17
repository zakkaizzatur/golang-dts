package models

import (
	"time"

	"gorm.io/gorm"
)

type Orders struct {
	ID        uint           `gorm:"primaryKey;autoIncrement" json:"id,omitempty"`
	CreatedAt time.Time      `json:"created_at,omitempty"`
	UpdatedAt time.Time      `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	CustomerName string    `gorm:"not null;type:varchar(255)" json:"customer_name"`
	OrderedAt    time.Time `json:"ordered_at"`
	Items        []Items   `gorm:"foreignKey:OrderID;constraint:OnUpdate:CASCADE" json:"items"`
}