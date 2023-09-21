package models

import (
	"time"

	"gorm.io/gorm"
)

type Invoice struct {
	ID                 uint `gorm:"primaryKey;autoIncrement"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DeletedAt          gorm.DeletedAt `gorm:"index"`
	StreetAddress      string         `gorm:"not null"`
	City               string         `gorm:"not null"`
	PostCode           string         `gorm:"not null"`
	Country            string         `gorm:"not null"`
	Status             string         `gorm:"not null"`
	IssueDate          time.Time      `gorm:"not null"`
	PaymentTerms       string         `gorm:"not null"`
	ProjectDescription string         `gorm:"not null"`
	Total              float64        `gorm:"not null"`
	ClientID           uint           `gorm:"not null"`
	Client             Client         `gorm:"foreignKey:ID"`
	ItemID             uint           `gorm:"not null"`
	Item               []Item         `gorm:"foreignKey:ID"`
}
