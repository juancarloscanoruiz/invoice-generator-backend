package models

import (
	"time"

	"gorm.io/gorm"
)

type Item struct {
	ID        int `gorm:"primaryKey;autoIncrement"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string         `gorm:"not null"`
	Quantity  int            `gorm:"not null"`
	Price     float64        `gorm:"not null"`
	InvoiceID int            `gorm:"not null"`
	Invoice   Invoice        `gorm:"foreignKey:InvoiceID"`
}
