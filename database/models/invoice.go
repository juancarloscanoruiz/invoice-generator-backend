package models

import (
	"time"

	"gorm.io/gorm"
)

type Invoice struct {
	ID                 int            `gorm:"primaryKey;autoIncrement" json:"id"`
	CreatedAt          time.Time      `json:"createdAt"`
	UpdatedAt          time.Time      `json:"updatedAt"`
	DeletedAt          gorm.DeletedAt `gorm:"index" json:"deletedAt"`
	StreetAddress      string         `gorm:"not null" json:"streetAddress"`
	City               string         `gorm:"not null" json:"city"`
	PostCode           string         `gorm:"not null" json:"postCode"`
	Country            string         `gorm:"not null" json:"country"`
	Status             string         `gorm:"not null" json:"status"`
	IssueDate          time.Time      `gorm:"not null" json:"issueDate"`
	PaymentTerms       string         `gorm:"not null" json:"paymentTerms"`
	ProjectDescription string         `gorm:"not null" json:"projectDescription"`
	Total              float64        `gorm:"not null" json:"total"`
	ClientID           int            `gorm:"not null" json:"clientId"`
	Client             Client         `gorm:"foreignKey:ClientID" json:"client"`
	Items              []Item         `gorm:"foreignKey:InvoiceID" son:"items"`
}
