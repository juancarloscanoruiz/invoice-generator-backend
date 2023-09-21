package models

import (
	"time"

	"gorm.io/gorm"
)

type Client struct {
	ID            uint `gorm:"primaryKey;autoIncrement"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
	Name          string         `gorm:"not null"`
	Email         string         `gorm:"not null"`
	StreetAddress string         `gorm:"not null"`
	City          string         `gorm:"not null"`
	PostCode      string         `gorm:"not null"`
	Country       string         `gorm:"not null"`
}
