package models

import (
	"time"

	"gorm.io/gorm"
)

type Client struct {
	ID            int            `gorm:"primaryKey;autoIncrement" json:"id"`
	CreatedAt     time.Time      `json:"createdAt"`
	UpdatedAt     time.Time      `json:"updatedAt"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deletedAt"`
	Name          string         `gorm:"not null" json:"name"`
	Email         string         `gorm:"not null" json:"email"`
	StreetAddress string         `gorm:"not null" json:"streetAddress"`
	City          string         `gorm:"not null" json:"city"`
	PostCode      string         `gorm:"not null" json:"postCode"`
	Country       string         `gorm:"not null" json:"country"`
}
