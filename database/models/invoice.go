package models

import (
	"time"

	"gorm.io/gorm"
)

type Invoice struct {
	ID                 int `gorm:"primaryKey;autoIncrement"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DeletedAt          gorm.DeletedAt `gorm:"index" json:"deletedAt"`
	StreetAddress      string         `gorm:"not null"`
	City               string         `gorm:"not null"`
	PostCode           string         `gorm:"not null"`
	Country            string         `gorm:"not null"`
	Status             string         `gorm:"not null"`
	IssueDate          time.Time      `gorm:"not null"`
	PaymentTerms       string         `gorm:"not null"`
	ProjectDescription string         `gorm:"not null"`
	Total              float64        `gorm:"not null"`
	ClientID           int            `gorm:"not null"`
	Client             Client         `gorm:"foreignKey:ClientID"`
}
