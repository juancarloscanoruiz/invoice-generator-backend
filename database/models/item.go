package models

import (
	"time"

	"gorm.io/gorm"
)

type Item struct {
	ID        uint `gorm:"primaryKey;autoIncrement"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string         `gorm:"not null"`
	Quantity  uint           `gorm:"not null"`
	Price     float64        `gorm:"not null"`
}
