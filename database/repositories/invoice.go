package repositories

import (
	"invoice-app/database"
	"invoice-app/database/models"

	"gorm.io/gorm/clause"
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

func GetInvoice(id int) (*models.Invoice, error) {
	var invoice models.Invoice
	var err error
	err = database.Db.Preload("Client").Preload("Items").Where("id = ?", id).First(&invoice).Error
	if err != nil {
		return nil, err
	}
	return &invoice, nil
}

func CreateInvoice(invoice models.Invoice) (*models.Invoice, error) {
	newInvoice := models.Invoice{
		StreetAddress:      invoice.StreetAddress,
		City:               invoice.City,
		PostCode:           invoice.PostCode,
		Country:            invoice.Country,
		Status:             invoice.Status,
		IssueDate:          invoice.IssueDate,
		PaymentTerms:       invoice.PaymentTerms,
		ProjectDescription: invoice.ProjectDescription,
		Client:             invoice.Client,
	}

	invoiceErr := database.Db.Create(&newInvoice).Error

	var newItems []models.Item
	for _, item := range invoice.Items {
		item.InvoiceID = newInvoice.ID
		newItems = append(newItems, item)
	}

	itemErr := database.Db.Create(&newItems).Error

	if invoiceErr != nil {
		return nil, invoiceErr
	}
	if itemErr != nil {
		return nil, itemErr
	}
	newInvoice.Items = newItems
	return &newInvoice, nil
}

func DeleteInvoice(invoiceId int) (*models.Invoice, error) {
	var invoice models.Invoice
	deleteErr := database.Db.Clauses(clause.Returning{}).Where("id = ?", invoiceId).Delete(&invoice).Error
	if deleteErr != nil {
		return nil, deleteErr
	}
	return &invoice, nil
}
