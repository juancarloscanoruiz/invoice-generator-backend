package models

import (
	"time"

	"gorm.io/gorm"
)

type Item struct {
	ID        int            `gorm:"primaryKey;autoIncrement"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
	Name      string         `gorm:"not null" json:"name"`
	Quantity  int            `gorm:"not null" json:"quantity"`
	Price     float64        `gorm:"not null" json:"price"`
	InvoiceID int            `gorm:"not null" json:"invoiceId"`
	Invoice   Invoice        `gorm:"foreignKey:InvoiceID" json:"invoice"`
}
