package repositories

import (
	"invoice-app/database"
	"invoice-app/database/models"
)

func GetInvoices() ([]models.Invoice, error) {
	var invoices []models.Invoice
	var err error
	err = database.Db.Preload("Client").Preload("Items").Find(&invoices).Error
	if err != nil {
		return nil, err
	}
	return invoices, err
}
